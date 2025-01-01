package handler

import (
	model "apicha/foodshop/Model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Menuhandler struct {
	db *gorm.DB
}

func Newmenuhandler(db *gorm.DB) *Menuhandler {
	return &Menuhandler{db: db}
}

func (h *Menuhandler) GetAll(c *gin.Context) {
	var tables []model.Menumodel
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

func (h *Menuhandler) Create(c *gin.Context) {

	var createmenu model.Menumodel

	// Bind form data to struct
	if err := c.ShouldBind(&createmenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind form data: " + err.Error(),
		})
		return
	}

	file, _ := c.FormFile("file")

	// Save the file
	if err := c.SaveUploadedFile(file, "./Picture/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file: " + err.Error(),
		})
		return
	}

	// Set the picture filename in the menu model
	createmenu.PictureMenu = file.Filename

	// Create the menu in database
	if err := h.db.Create(&createmenu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create menu: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Menu created successfully",
		"data":    createmenu,
	})
}

func (h *Menuhandler) Update(c *gin.Context) {
	var updatemenu model.Menumodel
	var existingTable model.Menumodel

	id := c.Param("id")

	if err := h.db.First(&existingTable, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find existing menu: " + err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&updatemenu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind request data: " + err.Error(),
		})
		return
	}

	file, err := c.FormFile("file")
	if err == nil {
		if existingTable.PictureMenu != "" {
			if err := os.Remove("./Picture/" + existingTable.PictureMenu); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to remove old picture: " + err.Error(),
				})
				return
			}
		}

		destination := "./Picture/" + file.Filename
		if err := c.SaveUploadedFile(file, destination); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save uploaded file: " + err.Error(),
			})
			return
		}

		updatemenu.PictureMenu = file.Filename
	}

	if err := h.db.Model(&existingTable).Updates(updatemenu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update menu: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, existingTable)
}

func (h *Menuhandler) Deletemenu(c *gin.Context) {
	var existingTable model.Menumodel
	id := c.Param("id")

	if err := h.db.First(&existingTable, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find existing menu: " + err.Error(),
		})
		return
	}

	if err := h.db.Delete(&existingTable).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if existingTable.PictureMenu != "" {
		if err := os.Remove("./Picture/" + existingTable.PictureMenu); err != nil && !os.IsNotExist(err) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete picture file: " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"delete": "successful",
	})
}
