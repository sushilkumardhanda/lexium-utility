package controllers

import (
	"lexium-utility/datahandler"
	"lexium-utility/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScheamList(c *gin.Context) {
	var input datahandler.SelectedITR
	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	schemaList, err := repository.GetScheamList(input.ITR)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"schemaList": schemaList})
	return
}
