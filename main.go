package main

import "github.com/axliupore/axapi/axapi-backend/initialize"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title 					   Ax-API 接口文档
// @version 1.0
// @description 			   api接口文档
// @securityDefinitions.apiKey ApiKeyAuth
// @in 						   header
// @name 					   a-token
// @BasePath 				   /api
func main() {
	initialize.InitConfig()
}
