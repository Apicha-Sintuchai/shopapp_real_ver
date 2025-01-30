package handler

import (
	model "apicha/foodshop/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Orderhandler struct {
	db *gorm.DB
}

type Day struct {
	Day string `json:"day"`
}

func Neworderhandle(db *gorm.DB) *Orderhandler {
	return &Orderhandler{db: db}
}

func (h *Orderhandler) GetAll(c *gin.Context) {
	var findall []model.Customerordermodel
	if err := h.db.Preload("Orders").Where("donestatus = ?", "").Find(&findall).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, findall)

}

func (h *Orderhandler) Create(c *gin.Context) {
	var newOrder model.Customerordermodel
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var totalprice int
	for _, v := range newOrder.Orders {
		i, _ := strconv.Atoi(v.Price)

		totalprice += i
	}
	newOrder.Status = "ยังไม่ชำระเงิน"
	newOrder.Totalprice = strconv.Itoa(totalprice)
	if err := h.db.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

// จ่ายเงิน
func (h *Orderhandler) Update(c *gin.Context) {
	id := c.Param("id")
	var updateOrder model.Customerordermodel
	if err := c.ShouldBindJSON(&updateOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.db.Model(&model.Customerordermodel{}).Where("id = ?", id).Updates(updateOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateOrder)
}

func (h *Orderhandler) Donestatus(c *gin.Context) {
	id := c.Param("id")
	var changestatus model.Customerordermodel

	if err := h.db.First(&changestatus, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	changestatus.Donestatus = "สำเร็จ"

	if err := h.db.Save(&changestatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, changestatus)
}

func (h *Orderhandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.Delete(&model.Customerordermodel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}

func (h *Orderhandler) GetMoney(c *gin.Context) {
	var findall []model.Customerordermodel
	if err := h.db.Preload("Orders").Where("donestatus = ? AND status = ? ", "สำเร็จ", "ชำระเงินแล้ว").Find(&findall).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, findall)

}

func (h *Orderhandler) Dopayment(c *gin.Context) {
	var findall []model.Customerordermodel
	if err := h.db.Preload("Orders").Where("status = ? ","ยังไม่ชำระเงิน").Find(&findall).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, findall)
}

func (h *Orderhandler) Updatepayment(c *gin.Context) {
	Id := c.Param("id")
	var update model.Customerordermodel

	if err := h.db.First(&update, Id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	update.Status = "ชำระเงินแล้ว"

	if err := h.db.Save(&update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, update)
}
