package sender

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewSms,
	NewSms2,
	SenderList,
)

func SenderList(
	sms *Sms,
	sms2 *Sms2,
) []Handler {
	return []Handler{
		sms,
		sms2,
	}
}
