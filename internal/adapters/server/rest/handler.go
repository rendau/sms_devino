package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"

	"github.com/rendau/sms_devino/internal/domain/entities"
)

func (a *St) hSend(c *gin.Context) {
	reqBody, err := c.GetRawData()
	if dopHttps.Error(c, err) {
		return
	}

	reqObj := &entities.SendReqSt{}
	err = json.Unmarshal(reqBody, reqObj)
	if dopHttps.Error(c, err) {
		a.lg.Infow("Fail to unmarshal request", "err", err, "reqBody", string(reqBody))
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
