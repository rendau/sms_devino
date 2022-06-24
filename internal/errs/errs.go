package errs

import (
	"github.com/rendau/dop/dopErrs"
)

const (
	PhoneRequired   = dopErrs.Err("phone_required")
	BadPhone        = dopErrs.Err("bad_phone")
	MessageRequired = dopErrs.Err("message_required")
	FailToSend      = dopErrs.Err("service_not_available")
)
