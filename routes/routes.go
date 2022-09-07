package routes

import (
	"adm-num/ent"
	"adm-num/graphql/resolver"

	"adm-num/utils"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(app *echo.Echo, client *ent.Client, config utils.Config) {

	// do not forget the middleware

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.AllowOrigins},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	//app.Use(middleware.Logger())

	app.GET("/", PlaygroundHandler())
	app.POST("/query", GraphqlHandler(client))
	app.Any("/subscriptions", GraphqlWsHandler(client))

	app.GET("/ws", PlaygroundWsHandler())
}

func PlaygroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func PlaygroundWsHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL WS", "/subscription")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func GraphqlHandler(client *ent.Client) echo.HandlerFunc {
	h := handler.NewDefaultServer(resolver.NewSchema(client))
	h.Use(entgql.Transactioner{TxOpener: client})

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func GraphqlWsHandler(client *ent.Client) echo.HandlerFunc {
	h := handler.New(resolver.NewSchema(client))
	h.AddTransport(transport.POST{})
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	h.Use(extension.Introspection{})
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
