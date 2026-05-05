package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/AkimioJR/MediaWarp/utils"
	"github.com/gin-gonic/gin"
)

func ImageCache(ttl time.Duration, reg *regexp.Regexp) gin.HandlerFunc {
	cachePool, err := utils.NewCache(ttl)
	if err != nil {
		panic(fmt.Sprintf("create image cache pool failed: %v", err))
	}
	cacheFunc := getCacheBaseFunc(cachePool, "图片", reg.String())

	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodGet && reg.MatchString(ctx.Request.URL.Path) {
			cacheFunc(ctx)
		}
		ctx.Next()
	}
}
