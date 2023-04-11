package sliceHelper

import "reflect"

func InArr[T any](s T, list []T) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, s) {
			return true
		}
	}
	return false
}

func Unique[T any](list []T) []T {
	var amap = make(map[any]struct{})
	var newList []T
	for _, s := range list {
		if _, ok := amap[s]; ok {
			continue
		}
		amap[s] = struct{}{}
		newList = append(newList, s)
	}
	return newList
}
