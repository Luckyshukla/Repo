package main

import (
	//"fmt"
	//"path/filepath"
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"main.go/cantroller"
	"main.go/models"
	"main.go/services"
	"main.go/middleware"
	"net/http"
	"time"
	
)

func main() {

	router := gin.Default()
	gin.DisableConsoleColor()

// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)


//Custom HTTP configuration
	s := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}


	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())

	models.ConnectDataBase()






	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.Static("/", "./Upload")
	// router.POST("/upload", func (c *gin.Context) {
	// 		name := c.PostForm("name")
	// 		email := c.PostForm("email")
		
				
	// 		file, err := c.FormFile("file")
	// 		if err != nil {
	// 			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 		return
	// 	}
			 	
	// 	filename := filepath.Base(file.Filename)
	// 	if err := c.SaveUploadedFile(file, filename); err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 		return
	// 	}
		
	// 		 c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	// 	 })

	// Serves literal characters
	router.GET("/purejson", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	router.POST("/login",services.Login)
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	

	authorized := router.Group("/", middleware.BasicAuth())


	authorized.GET("/getuser", cantroller.ShowUserData)
	authorized.PUT("/update/:id", cantroller.Updatedata)
	
	authorized.POST("/post", cantroller.Createdata)
	authorized.DELETE("/delete/:id", cantroller.DeleteData)
	
	s.ListenAndServe()
}
