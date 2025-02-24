package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// 用户表（User）
type User struct {
	gorm.Model
	Email          string      `gorm:"uniqueIndex;type:varchar(255);not null"`
	PasswordHashed string      `gorm:"type:varchar(255);not null"`
	Nickname       string      `gorm:"type:varchar(100);not null"` // 昵称
	Role           string      `gorm:"type:varchar(50);not null"`  // 角色，如 buyer/seller/admin
	Products       []ProductId `gorm:"foreignKey:UserID"`          // 关联 ProductId
	IsDeleted      bool        `gorm:"default:false"`              // 软删除标记字段
}

type UserQuery struct {
	ctx context.Context
	db  *gorm.DB
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

func NewUserQuery(ctx context.Context, db *gorm.DB) *UserQuery {
	return &UserQuery{
		ctx: ctx,
		db:  db,
	}
}

// 创建用户
func (p *UserQuery) CreateUser(useremail, passwordHashed, nickname, role string, productIds []int32) (*User, error) {
	newUser := &User{
		Email:          useremail,
		PasswordHashed: passwordHashed,
		Nickname:       nickname,
		Role:           role,
	}

	if err := p.db.WithContext(p.ctx).Create(newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to create newUser: %v", err)
	}

	// 如果有产品 ID，则关联产品
	if len(productIds) > 0 {
		productIdQuery := NewProductIdQuery(p.ctx, p.db)
		products, err := productIdQuery.GetProductByIds(productIds)
		if err != nil {
			return nil, fmt.Errorf("failed to find products: %v", err)
		}
		// 这里更新 User 关联的 Products
		if err := p.db.WithContext(p.ctx).Model(newUser).Association("Products").Append(products); err != nil {
			return nil, fmt.Errorf("failed to associate products with user: %v", err)
		}
	}
	return newUser, nil
}

// 通过 Email 获取用户
func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Preload("Products").Where("email = ?", email).First(&user).Error
	return
}

// 通过 UserID 获取用户
func GetByUserId(db *gorm.DB, ctx context.Context, userID int32) (user *User, err error) {
	err = db.WithContext(ctx).Preload("Products").Where("id = ?", userID).First(&user).Error
	return
}

// 软删除用户
func DeleteByUserId(db *gorm.DB, ctx context.Context, userID int32) error {
	return db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("is_deleted", true).Error
}

// 更新用户信息

func (p *UserQuery) UpdateUser(db *gorm.DB, ctx context.Context, userID int32, Email, PasswordHashed, Nickname, Role string, Productid int32) (*User, error) {
	// 查找用户
	var user User
	err := db.WithContext(ctx).Where("id = ? AND is_deleted = false", userID).First(&user).Error
	if err != nil {
		return nil, err // 找不到用户或其他错误
	}

	// 更新字段
	if Email != "" {
		user.Email = Email // 如果 Email 不为空，更新 Email
	}
	if PasswordHashed != "" {
		user.PasswordHashed = PasswordHashed // 更新密码（已加密）
	}
	if Nickname != "" {
		user.Nickname = Nickname // 更新昵称
	}
	if Role != "" {
		user.Role = Role // 更新角色
	}
	// 保存更新后的用户信息
	err = db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return nil, err // 保存失败
	}

	// 返回更新后的用户信息
	return &user, nil
}

// 添加商品到用户
func (p *UserQuery) AddProductToUser(db *gorm.DB, ctx context.Context, userID uint, productID int32) (*User, error) {
	// 首先查询用户是否存在
	var user User
	err := db.WithContext(ctx).Where("id = ? AND is_deleted = false", userID).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user with id %d not found: %w", userID, err) // 用户不存在
	}
	// 创建商品记录
	product := ProductId{
		Productid: productID,
		UserID:    userID,
	}
	// 插入商品记录
	err = db.WithContext(ctx).Create(&product).Error
	if err != nil {
		return nil, fmt.Errorf("failed to add product to user: %w", err) // 插入商品失败
	}

	// 返回更新后的用户
	return &user, nil
}

// 获取某个用户的所有商品
func GetUserProducts(db *gorm.DB, ctx context.Context, userID uint) ([]ProductId, error) {
	var products []ProductId
	err := db.WithContext(ctx).Where("user_id = ?", userID).Find(&products).Error
	return products, err
}

// GetUserInfo 通过 ID 获取用户信息
// func GetUserInfo(db *gorm.DB, ctx context.Context, userID int32) (user *User, err error) {
// 	err = db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
// 	return user, err
// }

func (p *UserQuery) GetUserInfo(db *gorm.DB, ctx context.Context, userid int32) (*User, []int32, error) {
	var user User

	// 在查询条件中添加 is_deleted 检查，确保用户未被删除
	err := p.db.WithContext(p.ctx).
		Preload("Products", "is_deleted = ?", false).      // 只预加载未删除的产品
		Where("id = ? AND is_deleted = ?", userid, false). // 确保用户未被删除
		First(&user).Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find user: %v", err)
	}

	// 提取 productIds
	productIds := make([]int32, len(user.Products))
	for i, product := range user.Products {
		productIds[i] = product.Productid
	}

	return &user, productIds, nil
}
