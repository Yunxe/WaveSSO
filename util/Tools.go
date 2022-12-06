package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerWarpper(f func(c *gin.Context) (error, any)) gin.HandlerFunc {
	return func(c *gin.Context) {
		err, data := f(c)
		if err == nil {
			c.JSON(http.StatusOK, data)
			return
		}
		for k, v := range ErrorMap {
			if err == k {
				c.JSON(v.HttpCode, v.Status)
				return
			}
		}
		c.JSON(ErrorMap[COMMON_ERR].HttpCode, ErrorMap[COMMON_ERR].Status)
		return
	}
}
