package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	jwt "github.com/xgd16/gf-jwt/v2"
	"testing"
	"time"
)

func TestTokenBlackTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		auth := jwt.New(&jwt.GfJWTMiddleware{
			Realm:         "client",
			Key:           []byte("uziA9mep5qZ2bMeR4OpKyKV7TFKb60x2L56LJb1p1aFR5i5C24bOu8vzCtRx4SiA"),
			Timeout:       time.Minute * 5,
			MaxRefresh:    time.Minute * 5,
			IdentityKey:   "id",
			TokenLookup:   "header: Authorization, query: token, cookie: jwt",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
		})

		g.DumpWithType(auth.TokenToBlackList(gctx.New(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRGYW5zIjo1LCJhZGRGYW5zRm9yRGFpbHkiOjQsImFnZW50SWQiOjMwLCJhZ2VudFN0YXR1cyI6MSwiZW1haWwiOm51bGwsImV4cCI6MTY5ODkxMTkzNzAxMywiZXhwaXJlVGltZSI6IjIwMjMtMTAtMjggMTY6MjY6NDIiLCJpZCI6MSwibGFzdExvZ2luVGltZSI6IjIwMjMtMTAtMjYgMTU6NTU6MTAiLCJsb2dpblR5cGUiOjEsIm1vYmlsZSI6bnVsbCwibmlja05hbWUiOiLotbDmtYHnqIsiLCJvbmxpbmVTdGF0dXMiOjEsIm9yaWdfaWF0IjoxNjk4MzA3MTM3MDEzLCJwcmVzZW5jZSI6MCwicmVzZXRUaW1lRm9yRGFpbHkiOiIwNTowMDowMCIsInJlc291cmNlR3JvdXBJZCI6bnVsbCwic3RhdHVzIjoxLCJ0cmFuc2xhdGVTZXR0aW5ncyI6InpoLUNIUyIsInVzZXJOYW1lIjoib28xMjM0NTYifQ.8ckj4Er7ez4U-M0HR1PT_NAmVwUWySUDXE20m2PXjmk"))
	})
}
