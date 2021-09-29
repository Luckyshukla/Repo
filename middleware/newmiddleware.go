package middleware
// import(
// 	"github.com/gin-gonic/gin"
// 	jwt "github.com/kyfk/gin-jwt"
// 	"net/http"
// )
// func NewAuth() (jwt.Auth, error) {
// 	return jwt.New(jwt.Auth{
// 		SecretKey: []byte("must change here"),
// 		Authenticator: func(c *gin.Context) (jwt.MapClaims, error) {
// 			// var req struct {
// 			// 	Username string `json:"username"`
// 			// 	Password string `json:"password"`
// 			// }
// 			// if err := c.ShouldBind(&req); err != nil {
// 			// 	return nil, jwt.ErrorAuthenticationFailed
// 			// }

// 			// u := naiveDatastore[req.Username] // change here fetching from read datastore
// 			// if u.Password != req.Password {
// 			// 	return nil, jwt.ErrorAuthenticationFailed
// 			// }

// 			////

// 			var input req

// 			if err := c.ShouldBindJSON(&input); err != nil {
// 				return nil, jwt.ErrorAuthenticationFailed
// 			}
	
// 			var user models.Userinfo
			
// 			if err := models.DB.Where("username = ? AND password = ?",input.Username,input.Password).First(&user).Error; err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Please provide valid login detail"})
// 			return nil, jwt.ErrorAuthenticationFailed
// 			//

// 			return jwt.MapClaims{
// 				"username": user.Username,
// 				"password": user.password,
// 			}, nil
// 		},
// 		UserFetcher: func(c *gin.Context, claims jwt.MapClaims) (interface{}, error) {
// 			username, ok := claims["username"].(string)
// 			if !ok {
// 				return nil, nil
// 			}
// 			u, ok := naiveDatastore[username]
// 			if !ok {
// 				return nil, nil
// 			}
// 			return u, nil
// 		},
// 	})
// }
