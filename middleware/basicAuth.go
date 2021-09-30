package middleware

import (
	"github.com/gin-gonic/gin"
	//"main.go/models"
)
// BasicAuth takes as argument a map[string]string where
// the key is the user name and the value is the password.
func BasicAuth() gin.HandlerFunc {
	//var user models.Userinfo
	return gin.BasicAuth(gin.Accounts{
		"user": "Password",
	})
}
