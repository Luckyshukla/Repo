package main

import (
	"github.com/gin-gonic/gin"
	"main.go/cantroller"
	"main.go/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()
	r.GET("/getuser", cantroller.ShowUserData)
	r.PATCH("/update/:id", cantroller.Updatedata)
	//r.GET("/get",   cantroller.findDetails)
	r.POST("/post", cantroller.Createdata)
	r.DELETE("/delete/:id", cantroller.DeleteData)
	r.Run()
}
