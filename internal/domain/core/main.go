package core

import (
	"time"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/sms_devino/internal/cns"
	"github.com/rendau/sms_devino/internal/domain/entities"
	"github.com/rendau/sms_devino/internal/domain/util"
	"github.com/rendau/sms_devino/internal/errs"
)

func (c *St) Send(obj *entities.SendReqSt) (*entities.SendRepSt, error) {
	if err := c.ValidateSendReqObj(obj); err != nil {
		return nil, err
	}

	reqObj := &entities.DevinoSendReqSt{
		Messages: []entities.DevinoSendReqMessageSt{
			{
				From:     c.senderName,
				To:       obj.To,
				Text:     obj.Text,
				Validity: cns.SmsValidity,
				Priority: cns.SmsPriority,
			},
		},
	}

	repObj := &entities.DevinoSendRepSt{}

	resp, err := c.httpc.Send(&httpc.OptionsSt{
		Method:        "POST",
		Uri:           "sms/messages",
		RetryCount:    1,
		RetryInterval: 2 * time.Second,

		ReqObj: reqObj,
		RepObj: repObj,
	})
	if err != nil {
		return nil, err
	}

	if len(repObj.Result) != 1 {
		resp.LogError("Fail to send SMS, wrong msg-count in reply", nil)
		return nil, dopErrs.ServiceNA
	}

	resultItem := repObj.Result[0]

	if resultItem.Code != cns.DevinoMessageStatusCodeOk {
		resp.LogError("Fail to send SMS", nil, "reply_code", resultItem.Code)
		if resultItem.Code == cns.DevinoMessageStatusCodeRejected {
			return nil, errs.FailToSend
		}
		return nil, dopErrs.ServiceNA
	}

	result := &entities.SendRepSt{
		ID: resultItem.MessageId,
	}

	return result, nil
}

func (c *St) ValidateSendReqObj(obj *entities.SendReqSt) error {
	obj.To = util.NormalizePhone(obj.To)
	if !util.ValidatePhone(obj.To) {
		return errs.BadPhone
	}

	if obj.Text == "" {
		return errs.MessageRequired
	}

	return nil
}
