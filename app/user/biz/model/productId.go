package model

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// 商品表（Product）
type ProductId struct {
	gorm.Model
	Productid int32 `gorm:"type:int;not null;uniqueIndex"`
	UserID    uint  `gorm:"index"` // 外键，指向 User 的 ID
	IsDeleted bool  `gorm:"default:false"`
}

type ProductIdQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// TableName 指定表名
func (ProductId) TableName() string {
	return "product_id"
}

func NewProductIdQuery(ctx context.Context, db *gorm.DB) *ProductIdQuery {
	return &ProductIdQuery{
		ctx: ctx,
		db:  db,
	}
}

// 获取产品 ID
func (c *ProductIdQuery) GetProductByIds(ids []int32) ([]ProductId, error) {
	var products []ProductId
	err := c.db.WithContext(c.ctx).Where("productid IN ?", ids).Find(&products).Error
	return products, err
}

// 根据产品 ID 删除产品
func (c *ProductIdQuery) DeleteProductById(id int32) error {
	if id <= 0 {
		return errors.New("invalid product ID")
	}

	// 执行更新操作
	result := c.db.WithContext(c.ctx).
		Model(&ProductId{}).
		Where("productid = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"updated_at": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no product found or already deleted")
	}
	return nil
}
