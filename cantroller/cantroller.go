package cantroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

// Read Data from data base

func ShowUserData(c *gin.Context) {
	var info[]  models.Userinfo
	models.DB.Find(&info)

	c.JSON(200, gin.H{"data": info})
}

//Create User data

type CreateUserInput struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"Last_name"`
	Phone      string `json:"phone" `
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func Createdata(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//Create User
	user := models.Userinfo{First_name: input.First_name, Last_name: input.Last_name, Phone: input.Phone, Username: input.Username, Password: input.Password}
	models.DB.Create(&user)
	c.JSON(200, gin.H{"data": user})

}

//Update data

type UpdateUserdata struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"Last_name"`
	Phone      string `json:"phone" `
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func Updatedata(c *gin.Context) {
	//Get modal if exist
	var user models.Userinfo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	// Validate input
	var input UpdateUserdata
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": input})
		return
	}
	models.DB.Table("userinfos").Where("id = ?", c.Param("id")).Updates(&input)
	c.JSON(200, gin.H{"data": input})
}

// func Updatedata(c *gin.Context) {
// 	var person UpdateUserdata
// 	id := c.Params.ByName("Username")
// 	if err := models.DB.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	}
// 	c.BindJSON(&person)
// 	models.DB.Save(&person)
// 	//db.Save(&person)
// 	c.JSON(200, person)
// }

//Delete data
func DeleteData(c *gin.Context) {
	var user models.Userinfo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Not found !"})
		return
	}
	models.DB.Delete(&user)

	c.JSON(200, gin.H{"data": true})
}