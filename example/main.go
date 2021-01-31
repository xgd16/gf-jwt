package main

import (
	"github.com/gogf/gf-jwt/example/api"
	"github.com/gogf/gf-jwt/example/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// authHook is the HOOK function implements JWT logistics.
func middlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func main() {
	s := g.Server()
	s.BindHandler("/", api.Work.Works)
	s.BindHandler("POST:/login", api.Auth.LoginHandler)
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Middleware.CORS, middlewareAuth)
		g.ALL("/info", api.Work.Info)
		g.ALL("/refresh_token", api.Auth.RefreshHandler)
		g.ALL("/logout", api.Auth.LogoutHandler)
	})
	s.SetPort(8000)
	s.Run()
}
