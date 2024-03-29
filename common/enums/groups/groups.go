package groups

import (
	"austin-v2/common/enums/channelType"
	"austin-v2/common/enums/messageType"
)

func GetAllGroupIds() []string {
	list := make([]string, 0)
	for _, ct := range channelType.TypeCodeEn {
		for _, mt := range messageType.TypeCodeEn {
			list = append(list, ct+"."+mt)
		}
	}
	return list
}
