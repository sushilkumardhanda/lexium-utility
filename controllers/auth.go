package controllers

import (
	"lexium-utility/datahandler"
	"lexium-utility/redis_repository"
	"lexium-utility/repository"
	"lexium-utility/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

func Logout(c *gin.Context) {
	username, err := utils.ExtractTokenUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = redis_repository.Delete(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func Login(c *gin.Context) {

	var input datahandler.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := repository.GetUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username doesnot exist."})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}
	_, err = redis_repository.GetSessionId(u.Username)
	if err == redis.Nil {
		sid := utils.GenerateSessionID()
		err = redis_repository.SetSessionId(u.Username, sid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in redis query"})
			return
		}
		token, err := utils.GenerateToken(u.Username, sid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "message": "logged in"})

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in redis Query"})
	} else {
		sid := utils.GenerateSessionID()
		token, err := utils.GenerateToken(u.Username, sid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "message": "already logged in"})
	}
}
func LoginConfirm(c *gin.Context) {
	username, err := utils.ExtractTokenUsername(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sid, err := utils.ExtractTokenSessionID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = redis_repository.SetSessionId(username, sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func Verify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}
