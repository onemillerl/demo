package middleware

// 全局认证中间件，用于从请求的会话中提取 user_id，并将其存储到 context.Context 中
import (
	"context"
	"fmt"

	"gomall_demo/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		token := session.Get("token")

		if userId == nil {
			c.Next(ctx)
			return
		}
		if token == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		ctx = context.WithValue(ctx, utils.TokenKey, token)

		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if utils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}
