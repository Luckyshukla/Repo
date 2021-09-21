package services
import(
	"net/http"
	"github.com/gin-gonic/gin"
	"main.go/models"
	"os"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"strings"
	//jwt "github.com/kyfk/gin-jwt"

)
type CheckLoginCredential struct {
	Username string `json "username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input CheckLoginCredential

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	var user models.Userinfo
	
	if err := models.DB.Where("username = ? AND password = ?",input.Username,input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please provide valid login detail"})
		return 
		//return nil, jwt.ErrorAuthenticationFailed
	}
	token, err := CreateToken(user.ID)
  		if err != nil {
     	c.JSON(http.StatusUnprocessableEntity, err.Error())
     	return
  	}
	  
	  c.JSON(http.StatusOK, gin.H{"access token":token})
}





func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") 
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
  }




func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}


func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("token")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
	   return strArr[1]
	}
	return ""
}

  
	func TokenValid(r *http.Request) error {
		token, err := VerifyToken(r)
		if err != nil {
		   return err
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		   return err
		}
		return nil
	}
	func Logout(c *gin.Context) {

		c.JSON(http.StatusOK, "Successfully logged out")
	}