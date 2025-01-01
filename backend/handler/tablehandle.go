package handler

import (
	model "apicha/foodshop/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tablehandler struct {
	db *gorm.DB
}

func NewTablehandler(db *gorm.DB) *Tablehandler {
	return &Tablehandler{db: db}
}

func (h *Tablehandler) GetTable(c *gin.Context) {
	var tables []model.Tablemodel
	if err := h.db.Find(&tables).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tables,
	})
}

func (h *Tablehandler) CreateTable(c *gin.Context) {
	var create model.Tablemodel
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.db.Create(&create).Error; err != nil { // Added .Error to check for creation errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, create) // Changed to StatusCreated for resource creation
}

func (h *Tablehandler) UpdateTable(c *gin.Context) {

	id := c.Param("id")

	var existingTable model.Tablemodel
	if err := h.db.First(&existingTable, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Table not found",
		})
		return
	}

	var updateData model.Tablemodel
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingTable.TableNumber = updateData.TableNumber

	if err := h.db.Save(&existingTable).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, existingTable)
}

func (h *Tablehandler) DeleteTable(c *gin.Context) {

	id := c.Param("id")

	if err := h.db.Delete(&model.Tablemodel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"delete": "successful",
	})
}
