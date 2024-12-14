// internal/handler/seckill.go
package handler

import (
	"net/http"
	"seckill-system/internal/model"
	"seckill-system/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SeckillActivityHandler struct {
	seckillService *service.SeckillActivityService
}

func NewSeckillActivityHandler() *SeckillActivityHandler {
	return &SeckillActivityHandler{
		seckillService: service.NewSeckillActivityService(),
	}
}

// Create 创建秒杀活动
func (h *SeckillActivityHandler) Create(c *gin.Context) {
	var activity model.SeckillActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.seckillService.CreateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "秒杀活动创建成功",
		"data":    activity,
	})
}

// Get 获取秒杀活动详情
func (h *SeckillActivityHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的活动ID"})
		return
	}

	activity, err := h.seckillService.GetActivity(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": activity,
	})
}

// Update 更新秒杀活动
func (h *SeckillActivityHandler) Update(c *gin.Context) {
	var activity model.SeckillActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.seckillService.UpdateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "秒杀活动更新成功",
		"data":    activity,
	})
}

// List 获取秒杀活动列表
func (h *SeckillActivityHandler) List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	statusStr := c.DefaultQuery("status", "-1")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		status = -1
	}

	activities, total, err := h.seckillService.ListActivities(page, pageSize, int8(status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":      activities,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}
