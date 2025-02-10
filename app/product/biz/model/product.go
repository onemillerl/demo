package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	IsDeleted   bool       `json:"is_deleted" gorm:"default:false"` // 软删除标志
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where("is_deleted = ? AND (name LIKE ? OR description LIKE ?)", false,
		"%"+q+"%", "%"+q+"%").Find(&products).Error
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p *ProductQuery) CreateProduct(name, description, picture string, price float32, categoryNames []string) (*Product, error) {
	categoryQuery := NewCategoryQuery(p.ctx, p.db)
	categories, err := categoryQuery.GetCategoriesByNames(categoryNames) // 这里会返回一个[]Category类型
	if err != nil {
		return nil, fmt.Errorf("failed to find categories: %v", err)
	}

	// Create the new product
	newProduct := &Product{
		Name:        name,
		Description: description,
		Picture:     picture,
		Price:       price,
		Categories:  categories, // categories是[]Category类型，直接赋值
	}

	// Save the product to the database
	if err := p.db.WithContext(p.ctx).Create(newProduct).Error; err != nil {
		return nil, fmt.Errorf("failed to create product: %v", err)
	}

	return newProduct, nil
}

func (p *ProductQuery) SoftDeleteProduct(productId int) error {
	// 将 IsDeleted 字段设置为 true，表示商品已被删除
	if err := p.db.WithContext(p.ctx).Model(&Product{}).Where("id = ?", productId).Update("is_deleted", true).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductQuery) UpdateProduct(productId int, updatedProduct *Product) error {
	err := p.db.WithContext(p.ctx).Model(&Product{}).Where("id = ?", productId).Updates(updatedProduct).Error
	return err
}
