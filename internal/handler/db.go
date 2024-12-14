// internal/handler/db.go
package handler

import (
	"net/http"
	"seckill-system/internal/dao"

	"github.com/gin-gonic/gin"
)

func TestDB(c *gin.Context) {
	sqlDB, err := dao.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取数据库实例失败",
			"error":   err.Error(),
		})
		return
	}

	// 测试数据库连接
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "数据库连接失败",
			"error":   err.Error(),
		})
		return
	}

	// 获取数据库统计信息
	stats := sqlDB.Stats()

	c.JSON(http.StatusOK, gin.H{
		"message": "数据库连接成功",
		"stats": gin.H{
			"max_open_connections": stats.MaxOpenConnections,
			"open_connections":     stats.OpenConnections,
			"in_use":               stats.InUse,
			"idle":                 stats.Idle,
		},
	})
}
