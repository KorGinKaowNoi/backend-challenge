package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	rootMax()
	setup()
}

func setup() {
	r := gin.Default()
	r.GET("/beef/summary", func(c *gin.Context) {
		response, err := getMeats()
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
		}

		c.JSON(http.StatusAccepted, response)
	})
	r.Run()
}
