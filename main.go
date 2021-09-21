package main

import (
	"github.com/gin-gonic/gin"
	"main.go/cantroller"
	"main.go/models"
	"main.go/services"
	
	//"main.go/middleware"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	models.ConnectDataBase()


	r.POST("/login",services.Login)
	r.GET("/getuser", cantroller.ShowUserData)
	r.PATCH("/update/:id", cantroller.Updatedata)
	
	r.POST("/post", cantroller.Createdata)
	r.DELETE("/delete/:id", cantroller.DeleteData)
	
	r.Run()
}
