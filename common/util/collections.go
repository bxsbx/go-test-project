package util

// list转map
func ListObjToMap[T, V any, K comparable](slice []T, f func(obj T) (K, V)) map[K]V {
	objMap := make(map[K]V, len(slice))
	for _, v := range slice {
		key, val := f(v)
		objMap[key] = val
	}
	return objMap
}

// list转换其他类型的list
func ListObjToListObj[T any, U any](slice []T, f func(obj T) U) []U {
	list := make([]U, len(slice))
	for i, v := range slice {
		list[i] = f(v)
	}
	return list
}

// list元素去重
func RemoveRepeatFromList[T comparable](slice []T) []T {
	itemMap := make(map[T]bool, len(slice))
	for _, v := range slice {
		itemMap[v] = true
	}
	list := []T{}
	for k, _ := range itemMap {
		list = append(list, k)
	}
	return list
}

// listObj元素去重
func RemoveRepeatFromListObj[T any, U comparable](slice []T, f func(obj T) U) []T {
	itemMap := make(map[U]T, len(slice))
	for _, v := range slice {
		itemMap[f(v)] = v
	}
	list := []T{}
	for _, v := range itemMap {
		list = append(list, v)
	}
	return list
}

// 拷贝数组，
func CopyList[T any](src []T) []T {
	list := make([]T, len(src))
	for i, v := range src {
		list[i] = v
	}
	return list
}

// 拷贝map
func CopyMap[T comparable, U any](src map[T]U) map[T]U {
	dtsMap := make(map[T]U, len(src))
	for k, v := range src {
		dtsMap[k] = v
	}
	return dtsMap
}

// 拷贝Chan，暂时不考虑使用
func CopyChan[T any](src chan T) chan T {
	dtsChan := make(chan T, len(src))
	for v := range src {
		dtsChan <- v
	}
	return dtsChan
}

// 从map中获取keys
func GetKeysFromMap[T comparable](aMap map[T]interface{}) []T {
	list := make([]T, len(aMap))
	i := 0
	for k := range aMap {
		list[i] = k
		i++
	}
	return list
}

// 从map中获取keys
func GetValuesFromMap[K comparable, V any](aMap map[K]V) []V {
	list := make([]V, len(aMap))
	i := 0
	for _, v := range aMap {
		list[i] = v
		i++
	}
	return list
}
