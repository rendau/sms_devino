package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/sms_devino/internal/domain/entities"
)

func (a *St) hSend(c *gin.Context) {
	reqObj := &entities.SendReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	if reqObj.Sync {
		result, err := a.cr.Send(reqObj)
		if dopHttps.Error(c, err) {
			return
		}

		c.JSON(http.StatusOK, result)
	} else {
		go func() { _, _ = a.cr.Send(reqObj) }()
	}
}
