// internal/service/seckill.go
package service

import (
	"errors"
	"seckill-system/internal/dao"
	"seckill-system/internal/model"
	"time"
)

type SeckillActivityService struct{}

func NewSeckillActivityService() *SeckillActivityService {
	return &SeckillActivityService{}
}

// CreateActivity 创建秒杀活动
func (s *SeckillActivityService) CreateActivity(activity *model.SeckillActivity) error {
	// 验证商品是否存在
	var product model.Product
	if err := dao.DB.First(&product, activity.ProductID).Error; err != nil {
		return errors.New("商品不存在")
	}

	// 验证时间是否合法
	if activity.StartTime.After(activity.EndTime) {
		return errors.New("结束时间不能早于开始时间")
	}

	// 验证库存是否充足
	if activity.SeckillStock > product.Stock {
		return errors.New("秒杀库存不能大于商品库存")
	}

	// 检查是否存在时间冲突的活动
	var count int64
	err := dao.DB.Model(&model.SeckillActivity{}).
		Where("product_id = ? AND status != 2 AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))",
			activity.ProductID, activity.StartTime, activity.StartTime, activity.EndTime, activity.EndTime).
		Count(&count).Error

	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该商品在此时间段已存在秒杀活动")
	}

	return dao.DB.Create(activity).Error
}

// GetActivity 获取秒杀活动详情
func (s *SeckillActivityService) GetActivity(id uint) (*model.SeckillActivity, error) {
	var activity model.SeckillActivity
	if err := dao.DB.First(&activity, id).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// UpdateActivity 更新秒杀活动
func (s *SeckillActivityService) UpdateActivity(activity *model.SeckillActivity) error {
	if activity.ID == 0 {
		return errors.New("活动ID不能为空")
	}

	// 查询原有活动
	var oldActivity model.SeckillActivity
	if err := dao.DB.First(&oldActivity, activity.ID).Error; err != nil {
		return errors.New("活动不存在")
	}

	// 如果活动已经开始，则不允许修改
	if oldActivity.Status == 1 {
		return errors.New("活动已开始，不能修改")
	}

	return dao.DB.Save(activity).Error
}

// ListActivities 获取活动列表
func (s *SeckillActivityService) ListActivities(page, pageSize int, status int8) ([]model.SeckillActivity, int64, error) {
	var activities []model.SeckillActivity
	var total int64

	query := dao.DB.Model(&model.SeckillActivity{})
	if status != -1 {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&activities).Error; err != nil {
		return nil, 0, err
	}

	return activities, total, nil
}

// UpdateActivityStatus 更新活动状态
func (s *SeckillActivityService) UpdateActivityStatus() error {
	now := time.Now()

	// 更新已开始的活动
	err := dao.DB.Model(&model.SeckillActivity{}).
		Where("status = 0 AND start_time <= ?", now).
		Update("status", 1).Error
	if err != nil {
		return err
	}

	// 更新已结束的活动
	return dao.DB.Model(&model.SeckillActivity{}).
		Where("status = 1 AND end_time <= ?", now).
		Update("status", 2).Error
}
