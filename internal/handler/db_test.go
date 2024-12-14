// internal/handler/db_test.go
package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"seckill-system/internal/handler"

	"github.com/gin-gonic/gin"
)

func TestDBConnection(t *testing.T) {
	// 设置测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	r := gin.New()
	r.GET("/test/db", handler.TestDB)

	// 创建测试请求
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test/db", nil)

	// 执行请求
	r.ServeHTTP(w, req)

	// 检查响应状态码
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 %d, 得到 %d", http.StatusOK, w.Code)
	}
}
