package cns

const (
	SmsSenderName = "Mechta.kz"
	SmsValidity   = 1800
	SmsPriority   = 3
)

const (
	DevinoMessageStatusCodeOk       = "OK"
	DevinoMessageStatusCodeRejected = "REJECTED"
)

const (
	DevinoErrBillingError    = "billing.error"     // Требуется оплата
	DevinoErrForbidden       = "forbidden"         // Отправка запрещена
	DevinoErrUnknown         = "unknown"           // Неизвестная ошибка
	DevinoErrInvalid         = "invalid"           // Неправильно указан номер телефона (messages[i].to | messages[i].validity | messages[i].callbackUrl)
	DevinoErrLengthTooLong   = "length.too.long"   // Превышена максимальная длина номера телефона (messages[i].to | messages[i].text)
	DevinoErrMustBeNotNull   = "must.be.not.null"  // Массив messages не может быть пустым
	DevinoErrNotAvailable    = "not.available"     // Неправильно указан отправитель
	DevinoErrTooManyMessages = "too.many.messages" // Превышен максимальный размер массива messages
)
