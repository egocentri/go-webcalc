  package last

import (
    "github.com/gin-gonic/gin"
    "github.com/egocentri/go-webcalc/package_web/handler.go"
func main() {
	a := gin.Default()
	a.POST("/api/v1/calculate", handler.calculateHandler)
	a.Run(":8080")
}
