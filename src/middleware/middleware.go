package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

func DummyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.GetHeader("Accept"))

		if c.GetHeader("Accept") != "application/vnd.studygroup+json;version=1" {
			respondWithError(400, "not accept", c)
		}
		c.Next()
	}
}
