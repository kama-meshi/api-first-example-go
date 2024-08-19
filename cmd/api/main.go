package main

import (
	"log"
	"net"

	"github.com/kama-meshi/api-first-example-go/internal/api"
	pkg "github.com/kama-meshi/api-first-example-go/pkg"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	// OASテンプレートの読み込み
	swagger, err := pkg.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading OAS template: %s", err)
	}
	// ホスト名での検証は行わない
	swagger.Servers = nil

	server := api.NewServer()

	e := echo.New()
	e.Use(echomiddleware.Logger())
	// OASテンプレートで指定したスキーマによる検証を行う
	e.Use(middleware.OapiRequestValidator(swagger))
	pkg.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", "8080")))
}
