package service

import (
	"context"
	"log"
	"os"
	"time"

	auth "gomall_demo/rpc_gen/kitex_gen/auth"

	"github.com/golang-jwt/jwt"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (*auth.DeliveryResp, error) {
	// 构建 JWT token
	log.Printf("准备开始调用 Token  generated userid: %d", req.UserId) // 添加日志
	log.Printf("生成Token的信息: %+v", req)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": req.UserId,                            // 用户 ID
		"role":   req.Role,                              // 用户角色（如 admin, shop_owner, customer）
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 24小时过期
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")

	// 对 JWT 进行签名，生成 token 字符串
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err // 如果签名失败，返回错误
	}
	log.Printf("Token successfully generated: %s", tokenString) // 添加日志

	// 返回生成的 token
	return &auth.DeliveryResp{
		Token: tokenString,
	}, nil
}
