package auth

import (
	"context"
	"log"
	"mock_project/internal/jwt"
	"mock_project/pb"

	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type AccContext struct {
	AccEmail string
	Role  string
}

func Auth(s pb.AccountServiceRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		token := authHeader[7:]
		log.Println("TOKEN",token)
		accountEmail, err := jwt.ParseToken(token)
		if err != nil {
			c.Next()
			return
		}

		r := c.Request
		ctx := r.Context()
		pbAccount, err := s.GetAccountByEmail(ctx, &pb.GetAccountByEmailRequest{Email: accountEmail})
		if err != nil {
			c.Next()
			return
		}

		acc := &AccContext{AccEmail: accountEmail, Role: pbAccount.Role.String()}

		ctx = context.WithValue(ctx, userCtxKey, acc)
		c.Request = r.WithContext(ctx)

		c.Next()
		// return
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *AccContext {
	raw, _ := ctx.Value(userCtxKey).(*AccContext)
	return raw
}


