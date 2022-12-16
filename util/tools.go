package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrWrapper(err error) (int, *Status) {
	return ErrorMap[err].HttpCode, ErrorMap[err].Status
}

func HandlerWrapper(f func(c *gin.Context) (error, any)) gin.HandlerFunc {
	return func(c *gin.Context) {
		err, data := f(c)
		if err == nil {
			c.JSON(http.StatusOK, data)
			return
		}
		for k, _ := range ErrorMap {
			if err == k {
				c.JSON(ErrWrapper(k))
				return
			}
		}
		c.JSON(ErrWrapper(COMMON_ERR))
		return
	}
}
