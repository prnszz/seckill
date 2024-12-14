// test/db_test.go
package test

import (
	"seckill-system/internal/dao"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	// 加载测试配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dao.InitMySQL()
}

func TestDatabaseConnection(t *testing.T) {
	sqlDB, err := dao.DB.DB()
	if err != nil {
		t.Fatalf("获取数据库实例失败: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("数据库连接测试失败: %v", err)
	}

	// 测试数据库配置
	stats := sqlDB.Stats()
	t.Logf("数据库连接池状态: 最大连接数=%d, 使用中=%d, 空闲=%d",
		stats.MaxOpenConnections,
		stats.InUse,
		stats.Idle,
	)
}
