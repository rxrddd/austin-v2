package transformHelper

func ArrayStringUniq(arr []string) []string {
	set := make(map[string]struct{}, 0)
	for _, s := range arr {
		set[s] = struct{}{}
	}
	var newStr = make([]string, 0)
	for k := range set {
		newStr = append(newStr, k)
	}
	return newStr
}

func GetIntKeysByMap(data map[int]string) []int {
	var list []int
	for key := range data {
		list = append(list, key)
	}
	return list
}
func GetStringValuesByMap(data map[int]string) []string {
	var list []string
	for _, value := range data {
		list = append(list, value)
	}
	return list
}
