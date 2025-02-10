package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gomall_demo/app/user/biz/dal/mysql"

	"gomall_demo/app/user/biz/model"
	"gomall_demo/app/user/infra/rpc"

	"gomall_demo/rpc_gen/kitex_gen/auth"
	user "gomall_demo/rpc_gen/kitex_gen/user"

	"golang.org/x/crypto/bcrypt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var secretKey = []byte("your_secret_key") // 用于签名 JWT 的密钥

var dbtest *gorm.DB

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info

func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	"root",
	// 	"root", // 修正拼写错误
	// 	"192.168.227.128",
	// 	"user", // 修正拼写错误
	// )
	// // dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	// dbtest, err = gorm.Open(mysql.Open(dsn),
	// 	&gorm.Config{
	// 		PrepareStmt:            true,
	// 		SkipDefaultTransaction: true,
	// 	},
	// )
	log.Printf("登录Email: %s", req.Email)
	// log.Printf("登录mysql: %+v", dbtest) // mysql.DB
	// y := &{Config:0xc001cbe630 Error:<nil> RowsAffected:0 Statement:0xc000606380 clone:1}
	row, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	log.Printf("GetByEmail row: %s,%d,%s", row.Email, row.ID, row.Role)

	if err != nil {
		return nil, err
	}
	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	// 检查用户状态
	if row.IsDeleted == true {
		return nil, fmt.Errorf("user account has been deleted")
	}

	// 2. 调用 auth 服务生成 token
	log.Printf("准备调用 auth 服务的 RPC 接口，UserId: %d, Role: %s", row.ID, row.Role)
	authresptoken, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: int32(row.ID),
		Role:   row.Role,
	})
	log.Printf("Auth RPC 调用结果: %v, 错误: %v", authresptoken, err)
	// if err != nil {
	// 	return nil, err
	// }
	// // 生成 JWT Token
	// token, err := generateJWT(int(row.ID), role)
	// if err != nil {
	// 	klog.Error("Error generating JWT token", "error", err)
	// 	return nil, err
	// }
	// klog.Infof("JWT generated successfully: userID=%d, role=%s, token=%s", row.ID, row.Role, authresptoken.Token)
	// log.Printf("Token : %s", authresptoken.Token)

	resp = &user.LoginResp{
		UserId: int32(row.ID),
		// Token:  "authresptoken.Token", // 返回生成的 Token
		Token: authresptoken.Token, // 返回生成的 Token

	}

	return resp, nil
}
