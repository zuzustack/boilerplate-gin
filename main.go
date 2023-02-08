package main

import (
	"zuzustack/learn/boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	utils.GetRouters(r)

	r.Run(":5555") // listen and serve on 0.0.0.0:8080
}
