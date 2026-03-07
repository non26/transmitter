package push

import (
	"github.com/gin-gonic/gin"
)

type IPushHandler interface {
	HandlePush(c *gin.Context)
}

type pushHandler struct {
	pService IPushService
}

func NewPushHandler(pService IPushService) IPushHandler {
	return &pushHandler{pService: pService}
}

func (h *pushHandler) HandlePush(c *gin.Context) {
	ctx := c.Request.Context()
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		response := NewPushResponse().Error("Invalid JSON")
		c.JSON(400, response)
		return
	}
	err := h.pService.Push(ctx, payload)
	if err != nil {
		response := NewPushResponse().Error(err.Error())
		c.JSON(500, response)
		return
	}

	response := NewPushResponse().Success()
	c.JSON(200, response)
}
