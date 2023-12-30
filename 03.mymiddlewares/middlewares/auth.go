package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//func Authenticate(c *gin.Context) {
//
//	if c.Request.Header.Get("Token") != "auth" {
//
//		c.AbortWithStatusJSON(401, gin.H{
//			"message": "token is not present",
//		})
//
//		return
//
//	}
//
//	c.Next()
//}

func Authenticate() gin.HandlerFunc {

	fmt.Println("Calling authenticate middleware...")
	return func(c *gin.Context) {
		if c.Request.Header.Get("Token") != "auth" {

			c.AbortWithStatusJSON(401, gin.H{
				"message": "token is not present",
			})

			return

		}

		c.Next()
	}

}
