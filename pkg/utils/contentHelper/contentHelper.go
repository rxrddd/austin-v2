package contentHelper

import "encoding/json"

func GetContentModel(contentModel interface{}, v interface{}) {
	marshal, _ := json.Marshal(contentModel)
	_ = json.Unmarshal(marshal, &v)
}
