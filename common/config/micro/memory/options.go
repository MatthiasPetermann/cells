package memory

import (
	"context"

	"github.com/pydio/go-os/config"

	"github.com/pydio/cells/common/config/source"
)

type changeSetKey struct{}

func withData(d []byte, f string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, changeSetKey{}, &config.ChangeSet{
			Data: d,
			// Format: f,
		})
	}
}

// WithChangeSet allows a changeset to be set
func WithChangeSet(cs *source.ChangeSet) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, changeSetKey{}, cs)
	}
}

// WithJSON allows the source data to be set to json
func WithJSON(d []byte) source.Option {
	return withData(d, "json")
}

// WithYAML allows the source data to be set to yaml
func WithYAML(d []byte) source.Option {
	return withData(d, "yaml")
}
