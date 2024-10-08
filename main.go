package main

import (
	"gee"
	"net/http"
)

// run后在另一个console查看
func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello suisbuds\n")
	})
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"suisbuds"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
