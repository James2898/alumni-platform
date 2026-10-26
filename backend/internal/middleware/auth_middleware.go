package middleware

import (
	"net/http"
	"strings"

	"github.com/James2898/alumni-platform/backend/internal/domain"
	"github.com/James2898/alumni-platform/backend/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserKey   = "current_user"
	ContextUserIDKey = "current_user_id"
	ContextRoleKey   = "current_user_role"
)

func AuthMiddleware(authSvc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var tokenStr string

		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		} else if cookie, err := c.Cookie("access_token"); err == nil {
			tokenStr = cookie
		}

		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			c.Abort()
			return
		}

		claims, err := authSvc.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextRoleKey, claims.Role)
		c.Next()
	}
}

func RequireRole(roles ...domain.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get(ContextRoleKey)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		userRole, ok := val.(domain.UserRole)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "invalid role context"})
			c.Abort()
			return
		}

		authorized := false
		for _, role := range roles {
			if userRole == role {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "you do not have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
