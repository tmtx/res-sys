package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/pkg/bus"
)

type Router struct {
	CommandBus bus.MessageBus
}

func (r *Router) GetPrefix() string {
	return "/guests"
}

func (r *Router) Routes() []server.RouteHandler {
	return []server.RouteHandler{
		{Path: "/create", Callback: r.Create, Method: http.MethodGet},
	}
}

func (r *Router) Create(c echo.Context) error {
	params := bus.MessageParams{
		"name":  "Test",
		"email": "kf@karlis.dev",
	}

	validatorMessages, err := r.CommandBus.DispatchSync(
		bus.NewCommand(app.CreateGuest, params),
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if len(validatorMessages) > 0 {
		return c.JSON(http.StatusOK, validatorMessages)
	}

	return server.SuccessResponse(c)
}
