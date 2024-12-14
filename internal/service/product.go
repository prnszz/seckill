// internal/service/product.go
package service

import (
	"errors"
	"seckill-system/internal/dao"
	"seckill-system/internal/model"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

// CreateProduct 创建商品
func (s *ProductService) CreateProduct(product *model.Product) error {
	return dao.DB.Create(product).Error
}

// GetProduct 获取商品详情
func (s *ProductService) GetProduct(id uint) (*model.Product, error) {
	var product model.Product
	if err := dao.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct 更新商品信息
func (s *ProductService) UpdateProduct(product *model.Product) error {
	if product.ID == 0 {
		return errors.New("product id is required")
	}

	// 只更新特定字段
	return dao.DB.Model(product).Updates(map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock,
		"status":      product.Status,
	}).Error
}

// DeleteProduct 删除商品
func (s *ProductService) DeleteProduct(id uint) error {
	return dao.DB.Delete(&model.Product{}, id).Error
}

// ListProducts 获取商品列表
func (s *ProductService) ListProducts(page, pageSize int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	// 获取总数
	if err := dao.DB.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := dao.DB.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
