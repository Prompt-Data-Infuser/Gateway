package handler

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/promptech1/infuser-gateway/router/middleware"
)

// Register: Rest API 대응을 위한 route 정보 정의
func (h *Handler) Register(apiGroup *echo.Group) {
	apiGroup.Use(middleware.KeyExtractor(h.ctx))

	apiGroup.GET("/:nameSpace/:version/:operation", h.ExecuteAPI)
	apiGroup.GET("/:nameSpace/:version/:operation/meta", h.ExecuteAPI)
}
