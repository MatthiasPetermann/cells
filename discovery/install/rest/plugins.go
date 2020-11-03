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

// Package rest is used once at install-time when running install via browser
package rest

import (
	"context"

	"github.com/jcuga/golongpoll"
	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/plugins"
	"github.com/pydio/cells/common/service"
)

func init() {
	plugins.RegisterInstall(func(ctx context.Context) {
		service.NewService(
			service.Name(common.SERVICE_REST_NAMESPACE_+common.SERVICE_INSTALL),
			service.Context(ctx),
			service.Tag(common.SERVICE_TAG_DISCOVERY),
			service.Description("RESTful Installation server"),
			service.WithWeb(func() service.WebHandler {
				eventManager, _ := golongpoll.StartLongpoll(golongpoll.Options{})
				return &Handler{
					eventManager,
				}
			}),
		)
	})
}
