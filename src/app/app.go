package app 

import (
	// "github.com/gorilla/mux"
	// "encoding/json"
	"fmt"
	"micky-svr/config"
	"micky-svr/middleware"
	"micky-svr/resource/page"
	"micky-svr/resource/user"
	"github.com/gin-gonic/gin"
)
type App struct {
	Router *gin.Engine
	// DB *config.PostGresConfig
}

type User struct {
	Username string
	Password string
}

func (a *App) Initialize() {
	// a.DB = config
	a.Router = gin.New()
	a.setRouters()
}

func AuthRequired() gin.HandlerFunc {
	return func( c *gin.Context ) {
		fmt.Println("middle ware")
	}
}

func AuthRequired1() gin.HandlerFunc {
	return func( c *gin.Context ) {
		fmt.Println("middle ware 1")
	}
}

func (app *App) setRouters() {
	app.Router.Use(gin.Logger())
	app.Router.Use(gin.Recovery())
	api := app.Router.Group("/api")
	api.GET("/hi", middleware.CheckLogin)
	obj := api.Group("/obj")
	obj.Use(middleware.AuthorizationMiddleware())
	{
		obj.GET("/login", func(c *gin.Context ){
			db, err := config.Connect()
			if err != nil {
				fmt.Println("Cant not connect db", err)
			}
			fmt.Println("connect", db)
			response := map[string]string{"status": "ok"}
			c.JSON(200, response)
		})
		obj.GET("/page",page.CreatePage)
	
		// api.POST("/user",user.CreateUser)
	}
	u := api.Group("/user")
	u.Use(AuthRequired1())
	user.UserRegister(u)
}
