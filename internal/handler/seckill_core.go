// internal/handler/seckill_core.go
package handler

import (
	"net/http"
	"seckill-system/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SeckillCoreHandler struct {
	seckillCoreService *service.SeckillCoreService
}

func NewSeckillCoreHandler() *SeckillCoreHandler {
	return &SeckillCoreHandler{
		seckillCoreService: service.NewSeckillCoreService(),
	}
}

// Seckill 秒杀接口
func (h *SeckillCoreHandler) Seckill(c *gin.Context) {
	// 获取参数
	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseUint(activityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的活动ID"})
		return
	}

	// TODO: 从JWT或Session中获取用户ID
	userID := uint(1) // 临时写死，实际项目中应该从认证中间件中获取

	// 执行秒杀
	if err := h.seckillCoreService.Seckill(userID, uint(activityID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "秒杀成功",
	})
}

// PreloadStock 预热库存接口
func (h *SeckillCoreHandler) PreloadStock(c *gin.Context) {
	activityIDStr := c.Param("id")
	activityID, err := strconv.ParseUint(activityIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的活动ID"})
		return
	}

	stockStr := c.Query("stock")
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的库存数量"})
		return
	}

	if err := h.seckillCoreService.PreloadStock(uint(activityID), stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "库存预热成功",
	})
}
