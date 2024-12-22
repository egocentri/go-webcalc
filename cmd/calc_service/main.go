package main

import (
    "github.com/gin-gonic/gin"
    "github.com/egocentri/go-webcalc/package_web"
)

func main() {
    a := gin.Default()
    a.POST("/api/v1/calculate", package_web.CalculateHandler) 
    a.Run(":8080")
}
