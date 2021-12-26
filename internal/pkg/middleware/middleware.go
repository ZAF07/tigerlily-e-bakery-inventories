package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
			fmt.Println("MIDDLEWARE")
        c.Next()
    }
}