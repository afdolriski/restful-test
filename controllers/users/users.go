package users

import (
	"net/http"

	"github.com/afdolriski/restful-test/database/models"
	"github.com/afdolriski/restful-test/lib/common"
	"github.com/afdolriski/restful-test/requests"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// User alias
type User = models.User

// CreateUserRequests for create validation
type CreateUserRequests = requests.CreateUserRequest

// UpdateUserRequests for update validation
type UpdateUserRequests = requests.UpdateUserRequest

// Index Controller
func Index(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []User

	db.Find(&users)
	length := len(users)
	serialized := make([]common.JSON, length, length)
	for i := 0; i < length; i++ {
		serialized[i] = users[i].Serialize()
	}

	c.JSON(http.StatusOK, serialized)
}

// Create Controller
func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var requestBody CreateUserRequests
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	user := User{
		Name:        requestBody.Name,
		Email:       requestBody.Email,
		PhoneNumber: requestBody.PhoneNumber,
	}

	db.Create(&user)
	c.JSON(http.StatusCreated, user.Serialize())
}

// Browse Controller
func Browse(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var user User

	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(http.StatusOK, user.Serialize())
}

// Update Controller
func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var requestBody UpdateUserRequests

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	user.Name = requestBody.Name
	user.Email = requestBody.Email
	user.PhoneNumber = requestBody.PhoneNumber

	db.Save(&user)
	c.JSON(http.StatusCreated, user.Serialize())
}

// Delete Controller
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	db.Delete(&user)
	c.Status(http.StatusNoContent)
}
