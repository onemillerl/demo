package service

import (
	"context"
	"errors"
	"log"
	"os"

	auth "gomall_demo/rpc_gen/kitex_gen/auth"

	"github.com/golang-jwt/jwt"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
// Run 解析并验证 Token，同时返回用户 ID 和角色
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (*auth.VerifyResp, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	log.Printf("前端传过来的Token: %s", req.Token)

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Println("Token 验证失败:", err)
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok1 := claims["userid"].(float64) // JWT 默认解析数字为 float64
		role, ok2 := claims["role"].(string)
		if !ok1 || !ok2 {
			return nil, errors.New("invalid token payload")
		}
		log.Printf("Token 认证成功的claims", claims)

		log.Printf("Token 认证成功, UserID: %d, Role: %s", int64(userID), role)

		return &auth.VerifyResp{
			Res:    true,
			UserId: int32(userID),
			Role:   role,
		}, nil
	}

	return nil, errors.New("invalid token")
}
