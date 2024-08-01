package controllers

import (
	"lexium-utility/datahandler"
	"lexium-utility/repository"
	"lexium-utility/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScheduleTree(c *gin.Context) {
	var input datahandler.SelectedITR_Schema
	// Bind the JSON input to the struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	elements, err := repository.ReadCollection(input.ITR, input.Schema)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tree := utils.CreateTree(elements)

	c.JSON(http.StatusOK, gin.H{"tree": tree})

}

func GetElement(c *gin.Context) {
	var input datahandler.SelectedElement
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	element, err := repository.ReadElement(input.ITR, input.Schema, input.ElementID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"element": element})
}
