package actions

import (
	"errors"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	db "github.com/toddkao/ecomm2/models"
)

// Authentication Struct
type Authentication struct {
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
}

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// Auth Middleware to check if request contains correct headers
func Auth(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		header := c.Request().Header
		var err error
		if header.Get("appid") == "app-id goes here" && header.Get("appsecret") == "app-secret goes here" {
			err = next(c)
		} else {
			err = errors.New("invalid credentials")
		}
		// do some work after calling the next handler
		return err
	}
}

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "ecomm2",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Apply Auth middleware to all routes
		app.Use(Auth)

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))
		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		db.InitDB()
		// api endoint prefix
		api := app.Group("/api/app/")

		lg := &db.LocationGroup{}
		api.GET("locations_groups", lg.ShowAll)
		api.POST("locations_groups", lg.Insert)
		api.DELETE("locations_groups/{id}", lg.Delete)
	}

	return app
}
