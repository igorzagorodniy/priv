package service

import "reflect"

// 2个slice 的交集,O(n^2)
func SimpleIntersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}

	return set
}

// 2个slice 的并集,O(n^2)
func SimpleUnion(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		set = append(set, el)
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if !contains(a, el) {
			set = append(set, el)
		}
	}

	return set
}

// 2个slice 是否有交集
func HasIntersect(a interface{}, b interface{}) bool {
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			return true
		}
	}

	return false
}

func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}
