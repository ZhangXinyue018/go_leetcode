package commons

import "reflect"

func CheckArrayEqual(array1 interface{}, array2 interface{}) bool {
	slide1 := ReflectSlide(array1)
	slide2 := ReflectSlide(array2)
	if len(slide1) != len(slide2) {
		return false
	}
	for i := 0; i < len(slide1); i++ {
		if slide1[i] != slide2[i] {
			return false
		}
	}
	return true
}

func ReflectSlide(slide interface{}) []interface{} {
	s := reflect.ValueOf(slide)
	if s.Kind() != reflect.Slice {
		panic("array")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}
