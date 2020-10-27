/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

// Package service acts as a factory for all Pydio services.
//
// Pydio services are wrapped around micro services with additional information and ability to declare themselves to the
// registry. Services can be of three main different type :
// - Generic Service : providing a Runner function, they can be used to package any kind of server library as a pydio service
// - Micro Service : GRPC-based services implementing specific protobuf-services
// - Web Service : Services adding more logic and exposing Rest APIs defined by the OpenAPI definitions generated from protobufs.
//
// Package provides additional aspects that can be added to any service and declared by "WithXXX" functions.
package service

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gyuho/goraph"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	microregistry "github.com/micro/go-micro/registry"
	"github.com/micro/go-web"
	"github.com/micro/misc/lib/addr"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/boltdb"
	"github.com/pydio/cells/common/config"
	"github.com/pydio/cells/common/dao"
	"github.com/pydio/cells/common/log"
	defaults "github.com/pydio/cells/common/micro"
	"github.com/pydio/cells/common/registry"
	servicecontext "github.com/pydio/cells/common/service/context"
	"github.com/pydio/cells/common/sql"
	errorUtils "github.com/pydio/cells/common/utils/error"
	net2 "github.com/pydio/cells/common/utils/net"
)

type Service interface {
	registry.Service

	Init(...ServiceOption)
	Options() ServiceOptions
	Done() chan (struct{})
}

func buildForkStartParams(name string) []string {
	params := []string{
		"start",
		"--fork",
		"--registry", viper.GetString("registry"),
		"--registry_address", viper.GetString("registry_address"),
		"--registry_cluster_address", viper.GetString("registry_cluster_address"),
		"--registry_cluster_routes", viper.GetString("registry_cluster_routes"),
		"--broker", viper.GetString("broker"),
		"--broker_address", viper.GetString("broker_address"),
	}
	if viper.GetBool("enable_metrics") {
		params = append(params, "--enable_metrics")
	}
	if viper.GetBool("enable_pprof") {
		params = append(params, "--enable_pprof")
	}
	params = append(params, name)
	return params
}

// Service for the pydio app
type service struct {
	// Computed by external functions during listing operations
	nodes    []*microregistry.Node
	excluded bool

	opts ServiceOptions
	node goraph.Node

	done chan (struct{})
}

// Checker is a function that checks if the service is correctly Running
type Checker interface {
	Check() error
}

type CheckerFunc func() error

// Check implements the Chercker interface
func (f CheckerFunc) Check() error {
	return f()
}

type Runner interface {
	Run() error
}

type RunnerFunc func() error

func (f RunnerFunc) Run() error {
	return f()
}

type Addressable interface {
	Addresses() []net.Addr
}

type Starter interface {
	Start() error
}

type Stopper interface {
	Stop() error
}

type StopperFunc func() error

func (f StopperFunc) Stop() error {
	defer func() {
		recover()
	}()
	return f()
}

type StopFunctionKey struct{}

// HandlerProvider returns a handler function from a micro service
type HandlerProvider func(micro.Service) interface{}

// NewService provides everything needed to run a service, no matter the type
func NewService(opts ...ServiceOption) Service {

	s := &service{
		opts: newOptions(append(mandatoryOptions, opts...)...),
		done: make(chan struct{}),
	}
	//opts: newOptions(append(mandatoryOptions(), opts...)...),

	name := s.Options().Name

	// Checking that the service is not bound to a certain IP
	peerAddress := config.Get("services", name, "PeerAddress").String()

	if peerAddress != "" && !net2.PeerAddressIsLocal(peerAddress) {
		log.Debug("Ignoring this service as peerAddress is not local", zap.String("name", name), zap.String("ip", peerAddress))
		return nil
	}

	ctx := s.Options().Context
	if ctx == nil {
		ctx = context.Background()
	}

	// Setting context
	// ctx, cancel := context.WithCancel(s.Options().Context)
	ctx = servicecontext.WithServiceName(ctx, name)

	if s.IsGRPC() {
		ctx = servicecontext.WithServiceColor(ctx, servicecontext.ServiceColorGrpc)
	} else if s.IsREST() {
		ctx = servicecontext.WithServiceColor(ctx, servicecontext.ServiceColorRest)

		// TODO : adding web services automatic dependencies to auth, this should be done in each service instead
		if s.Options().Name != common.SERVICE_REST_NAMESPACE_+common.SERVICE_INSTALL {
			s.Init(WithWebAuth())
		}
	} else {
		ctx = servicecontext.WithServiceColor(ctx, servicecontext.ServiceColorOther)
	}

	// Setting config
	s.Init(
		Context(ctx),
		Version(common.Version().String()),
	)

	// Finally, register on the main app registry
	s.Options().Registry.Register(s)

	return s
}

var mandatoryOptions = []ServiceOption{
	// Adding the config to the context
	AfterInit(func(s Service) error {
		ctx := s.Options().Context

		ctx = servicecontext.WithConfig(ctx, config.Get("services", s.Name()))

		s.Init(Context(ctx))

		return nil
	}),

	// Setting config watchers
	AfterInit(func(s Service) error {
		watchers := s.Options().Watchers
		if len(watchers) == 0 {
			return nil
		}
		registerWatchers(s, "services/"+s.Name(), watchers)
		return nil
	}),

	// Adding a check before starting the service to ensure only one is started if unique
	BeforeStart(func(s Service) error {

		// TODO - REDO THAT
		// if s.MustBeUnique() && defaults.RuntimeIsCluster() {
		// 	ctx := s.Options().Context
		// 	serviceName := s.Name()
		// 	cluster := registry.GetCluster(ctx, serviceName, &registry.NullFSM{})
		// 	nodeId := s.Options().Micro.Server().Options().Id
		// 	if err := cluster.Join(nodeId); err != nil {
		// 		return err
		// 	}
		// 	s.Init(Cluster(cluster))
		// 	<-cluster.LeadershipAcquired()
		// }

		return nil
	}),

	BeforeStop(func(s Service) error {
		if s.MustBeUnique() && defaults.RuntimeIsCluster() && s.Options().Cluster != nil {
			return s.Options().Cluster.Leave()
		}
		return nil
	}),

	// Adding a check before starting the service to ensure all dependencies are running
	BeforeStart(func(s Service) error {
		ctx := s.Options().Context

		log.Logger(ctx).Debug("BeforeStart - Check dependencies")

		for _, d := range s.Options().Dependencies {

			log.Logger(ctx).Debug("BeforeStart - Check dependency", zap.String("service", d.Name))

			err := Retry(func() error {
				runningServices, err := registry.ListRunningServices()
				if err != nil {
					return err
				}

				for _, r := range runningServices {
					if d.Name == r.Name() {
						return nil
					}
				}

				return fmt.Errorf("dependency %s not found", d.Name)
			}, 2*time.Second, 20*time.Minute) // This is long for distributed setup

			if err != nil {
				return err
			}
		}

		log.Logger(ctx).Debug("BeforeStart - Valid dependencies")

		return nil
	}),

	// Adding the dao to the context
	BeforeStart(func(s Service) error {

		ctx := s.Options().Context

		// Only if we have a DAO
		if s.Options().DAO == nil {
			return nil
		}

		var d dao.DAO
		driver, dsn := config.GetDatabase(s.Name())

		var prefix string
		switch v := s.Options().Prefix.(type) {
		case func(Service) string:
			prefix = v(s)
		case string:
			prefix = v
		default:
			prefix = ""
		}

		switch driver {
		case "mysql":
			if c := sql.NewDAO(driver, dsn, prefix); c != nil {
				d = s.Options().DAO(c)
			}
		case "sqlite3":
			if c := sql.NewDAO(driver, dsn, prefix); c != nil {
				d = s.Options().DAO(c)
			}
		case "boltdb":
			if c := boltdb.NewDAO(driver, dsn, prefix); c != nil {
				d = s.Options().DAO(c)
			}
		default:
			return fmt.Errorf("unsupported driver type: %s", driver)
		}

		if d == nil {
			return fmt.Errorf("Storage is not available")
		}

		ctx = servicecontext.WithDAO(ctx, d)

		s.Init(Context(ctx))

		return nil

	}),
}

func (s *service) Init(opts ...ServiceOption) {
	// process options
	for _, o := range opts {
		o(&s.opts)
	}
}

func (s *service) Options() ServiceOptions {
	return s.opts
}

func (s *service) BeforeInit() error {
	for _, f := range s.Options().BeforeInit {
		if err := f(s); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) AfterInit() error {
	for _, f := range s.Options().AfterInit {
		if err := f(s); err != nil {
			return err
		}
	}

	return nil
}

// Start a service and its dependencies
func (s *service) Start(ctx context.Context) {
	for _, f := range s.Options().BeforeStart {
		if err := f(s); err != nil {
			log.Logger(ctx).Error("Could not prepare start ", zap.Error(err))
			return
		}
	}

	if s.Options().MicroInit != nil {
		debug.SetPanicOnFault(true)

		if err := s.Options().MicroInit(s); err != nil {
			log.Logger(ctx).Error("Could not micro init ", zap.Error(err))
			return
		}

		go func() {
			for {
				err := s.Options().Micro.Run()
				if err == nil {
					break
				}

				if errorUtils.IsServiceStartNeedsRetry(err) {
					log.Logger(ctx).Info("Service failed to start - restarting in 10s", zap.Error(err))
					<-time.After(10 * time.Second)
					continue
				}

				log.Logger(ctx).Error("Could not run ", zap.Error(err))
				break
			}
		}()
	}

	// if s.Options().Web != nil {
	// 	go func() {
	// 		if err := s.Options().WebInit(s); err != nil {
	// 			log.Logger(ctx).Error("Could not web init ", zap.Error(err))
	// 			cancel()
	// 			return
	// 		}
	// 		if err := s.Options().Web.Run(); err != nil {
	// 			log.Logger(ctx).Error("Could not run ", zap.Error(err))
	// 			if stopper, ok := s.Options().Micro.(Stopper); ok {
	// 				stopper.Stop()
	// 			}
	// 			cancel()
	// 		}
	// 	}()
	// }

	for _, f := range s.Options().AfterStart {
		if err := f(s); err != nil {
			log.Logger(ctx).Error("Could not finalize start ", zap.Error(err))
		}
	}
}

// ForkStart uses a fork process to start the service
func (s *service) ForkStart(ctx context.Context, retries ...int) {

	name := s.Options().Name
	// ctx := s.Options().Context
	// cancel := s.Options().Cancel

	// Do not do anything
	cmd := exec.CommandContext(ctx, os.Args[0], buildForkStartParams(name)...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Logger(ctx).Error("Could not initiate fork ", zap.Error(err))
		// cancel()
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Logger(ctx).Error("Could not initiate fork", zap.Error(err))
		// cancel()
	}
	scannerOut := bufio.NewScanner(stdout)
	go func() {
		for scannerOut.Scan() {
			log.StdOut.WriteString(strings.TrimRight(scannerOut.Text(), "\n") + "\n")
		}
	}()
	scannerErr := bufio.NewScanner(stderr)
	go func() {
		for scannerErr.Scan() {
			log.StdOut.WriteString(strings.TrimRight(scannerErr.Text(), "\n") + "\n")
		}
	}()

	log.Logger(ctx).Debug("Starting SubProcess: " + name)
	if err := cmd.Start(); err != nil {
		log.Logger(ctx).Error("Could not start process", zap.Error(err))
		// cancel()
	}
	log.Logger(ctx).Debug("Started SubProcess: " + name)

	cmd.Wait()

	r := 0
	if len(retries) > 0 {
		r = retries[0]
	}
	if r >= 4 {
		log.Logger(ctx).Error("SubProcess finished: but reached max retries")
		// cancel()
		return
	} else {
		<-time.After(2 * time.Second)
		log.Logger(ctx).Error("SubProcess finished with error: trying to restart now")
		s.ForkStart(ctx, r+1)
	}
}

// Start a service and its dependencies
func (s *service) Stop() {

	ctx := s.Options().Context
	// cancel := s.Options().Cancel

	for _, f := range s.Options().BeforeStop {
		if err := f(s); err != nil {
			log.Logger(ctx).Error("Could not prepare stop ", zap.Error(err))
		}
	}

	// Cancelling context should stop the service altogether
	if stopper, ok := s.Options().Micro.(Stopper); ok {
		stopper.Stop()
	}

	for _, f := range s.Options().AfterStop {
		if err := f(s); err != nil {
			log.Logger(ctx).Error("Could not finalize stop ", zap.Error(err))
		}
	}
}

// IsRunning provides a quick way to check that a service is running.
func (s *service) IsRunning() bool {
	ctx := s.getContext()

	if err := s.Check(ctx); err != nil {
		return false
	}
	return true
}

// Check the status of the service (globally - not specific to an endpoint)
func (s *service) Check(ctx context.Context) error {

	running, err := registry.ListRunningServices()
	if err != nil {
		return err
	}

	for _, r := range running {
		if s.Name() == r.Name() {
			return nil
		}
	}

	return fmt.Errorf("Not found")
}

func (s *service) AddDependency(name string) {
	if name == s.Name() {
		return
	}
	s.Init(Dependency(name, []string{""}))
}

func (s *service) GetDependencies() []registry.Service {

	var r []registry.Service

	for _, d := range s.Options().Dependencies {
		for _, rr := range s.Options().Registry.GetServicesByName(d.Name) {
			r = append(r, rr)
		}
	}

	return r
}

func (s *service) Name() string {
	return s.Options().Name
}

func (s *service) Tags() []string {
	return s.Options().Tags
}

func (s *service) Version() string {
	return s.Options().Version
}

func (s *service) Description() string {
	return s.Options().Description
}

func (s *service) Regexp() *regexp.Regexp {
	return s.Options().Regexp
}

func (s *service) Address() string {
	defaultAddress := "0.0.0.0"
	address, err := addr.Extract(defaultAddress)
	if err != nil {
		return defaultAddress
	}
	port := s.Options().Port

	if port != "" {
		address = net.JoinHostPort(address, port)
	}

	return address
}

func (s *service) SetExcluded(ex bool) {
	s.excluded = ex
}

func (s *service) IsExcluded() bool {
	return s.excluded
}

func (s *service) SetRunningNodes(nodes []*microregistry.Node) {
	s.nodes = nodes
}

func (s *service) RunningNodes() []*microregistry.Node {

	nMap := make(map[string]*microregistry.Node)
	for _, p := range registry.GetPeers() {
		for _, ms := range p.GetServices(s.Name()) {
			for _, n := range ms.Nodes {
				if _, ok := nMap[n.Id]; !ok {
					nMap[n.Id] = n
				}
			}
		}
	}
	var nodes []*microregistry.Node
	for _, n := range nMap {
		nodes = append(nodes, n)
	}
	return nodes
}

func (s *service) DAO() interface{} {
	return s.Options().DAO
}

func (s *service) IsGeneric() bool {
	return !strings.HasPrefix(s.Name(), common.SERVICE_GRPC_NAMESPACE_) &&
		!strings.HasPrefix(s.Name(), common.SERVICE_WEB_NAMESPACE_) &&
		!strings.HasPrefix(s.Name(), common.SERVICE_REST_NAMESPACE_)
}

func (s *service) IsGRPC() bool {
	return strings.HasPrefix(s.Name(), common.SERVICE_GRPC_NAMESPACE_)
}

func (s *service) IsREST() bool {
	return strings.HasPrefix(s.Name(), common.SERVICE_WEB_NAMESPACE_) ||
		strings.HasPrefix(s.Name(), common.SERVICE_REST_NAMESPACE_)
}

// RequiresFork reads config fork=true to decide whether this service starts in a forked process or not.
func (s *service) AutoStart() bool {
	ctx := s.Options().Context
	return s.Options().AutoStart || servicecontext.GetConfig(ctx).Val("autostart").Bool()
}

// RequiresFork reads config fork=true to decide whether this service starts in a forked process or not.
func (s *service) RequiresFork() bool {
	ctx := s.Options().Context
	return s.Options().Fork || servicecontext.GetConfig(ctx).Val("fork").Bool()
}

// RequiresFork reads config fork=true to decide whether this service starts in a forked process or not.
func (s *service) MustBeUnique() bool {
	ctx := s.Options().Context
	return s.Options().Unique || servicecontext.GetConfig(ctx).Val("unique").Bool()
}

// func (s *service) Client() (string, client.Client) {
// 	return s.Options().Micro.Server().Options().Name, s.Options().Micro.Client()
// }

func (s *service) MatchesRegexp(o string) bool {
	if reg := s.Options().Regexp; reg != nil && reg.MatchString(o) {
		if matches := reg.FindStringSubmatch(o); len(matches) == 2 {
			s.Init(
				Name(matches[0]),
				Source(matches[1]),
			)

			return true
		}
	}

	return false
}

func (s *service) Done() chan struct{} {
	return s.done
}

func (s *service) getContext() context.Context {
	// if m, ok := (s.micro).(micro.Service); ok {
	// 	return m.Options().Context
	// } else if w, ok := (s.micro).(web.Service); ok {
	// 	return w.Options().Context
	// }

	return nil
}

// RestHandlerBuilder builds a RestHandler
type RestHandlerBuilder func(service web.Service, defaultClient client.Client) interface{}
