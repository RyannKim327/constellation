package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// BotProtection rejects requests whose User-Agent matches a known scraper agent
// (scrapy, python-requests, httpx) with 403 Access denied. This is a simple
// heuristic and should be hardened before being relied on in production.
func BotProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := strings.ToLower(c.GetHeader("User-Agent"))

		blocked := []string{
			"scrapy",
			"python-requests",
			"httpx",
		}

		if strings.Contains(strings.Join(blocked, ","), ua) {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "Access denied",
			})
			return
		}
		c.Next()
	}
}
