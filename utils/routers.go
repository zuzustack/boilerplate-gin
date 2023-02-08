package utils

import (
	"zuzustack/learn/boilerplate/controllers"

	"github.com/gin-gonic/gin"
)


func GetRouters(r *gin.Engine){

	// Where to set your router
	routers(r,"/ping", controllers.GetUser).Get()


}


type router struct{
	path string
	r *gin.Engine
	controller gin.HandlerFunc
}

func routers(r *gin.Engine, path string, controller gin.HandlerFunc) router {
	route := router{
		path: path,
		r: r,
		controller: controller,
	}
	return route
}


func (route router) Get(){
	route.r.GET(route.path, route.controller)
}

func (route router) Post(){
	route.r.POST(route.path, route.controller)
}

func (route router) Put(){
	route.r.PUT(route.path, route.controller)
}

func (route router) Delete(){
	route.r.DELETE(route.path, route.controller)
}