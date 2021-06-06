package collection

import (
	"reflect"
	"strings"
)

// SplitWithTrim split a string with a delimiter and trim every element of the list.
func SplitWithTrim(str, delim string) []string {
	rawFlds := strings.Split(str, delim)
	var ret []string
	for _, el := range rawFlds {
		if v := strings.TrimSpace(el); v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}

// ContainsStr looks for item into string slice and returns true if found.
func ContainsStr(searchItem string, slice []string) bool {
	for _, el := range slice {
		if searchItem == el {
			return true
		}
	}
	return false
}

// RemoveStr удаляет первое вхождение элемента в slice.
func RemoveStr(searchItem string, slice []string) []string {
	for i := len(slice) - 1; i >= 0; i-- {
		if searchItem == slice[i] {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Contains looks for item into some slice and returns true if found.
func Contains(searchItem interface{}, slice interface{}) bool {
	lst := reflect.ValueOf(slice)
	for i := 0; i < lst.Len(); i++ {
		if reflect.DeepEqual(searchItem, lst.Index(i).Interface()) {
			return true
		}
	}
	return false
}

// Index ищет схожий образец в коллекции и возвращает его индекс, если находит, или -1.
func Index(searchItem interface{}, slice interface{}) int {
	lst := reflect.ValueOf(slice)
	for i := 0; i < lst.Len(); i++ {
		if reflect.DeepEqual(searchItem, lst.Index(i).Interface()) {
			return i
		}
	}
	return -1
}

// SliceDiffElem описывает изменение одного элемента слайса для конкретной группы изменений..
type SliceDiffElem struct {
	elem interface{}
	row  int
}

// SliceDiff сравнивает 2 слайса и возвращает изменения по отношению первого объекта.
// Возвращает слайсы изменений элементов в одинаковых индексах, удаленные и добавленные элементы.
// В таком же порядке и рекомендуется проводить изменения на целевом объекте.
// В выходных слайсах удаленных и добавленных элементов порядок не соблюдается.
// Если исходные типы объектов не одинаковы, возвращаются пустые структуры данных по изменениям.
func SliceDiff(sl1 interface{}, sl2 interface{}) (
	upd []SliceDiffElem, del []SliceDiffElem, ins []SliceDiffElem) {
	if reflect.TypeOf(sl1) != reflect.TypeOf(sl2) {
		return upd, del, ins
	}
	var found bool
	lst1 := reflect.ValueOf(sl1)
	lst2 := reflect.ValueOf(sl2)
	sl1Len := lst1.Len()
	sl2Len := lst2.Len()
	if sl1Len == sl2Len {
		for i := sl1Len - 1; i >= 0; i-- {
			if !reflect.DeepEqual(lst1.Index(i).Interface(), lst2.Index(i).Interface()) {
				upd = append(upd, SliceDiffElem{lst2.Index(i).Interface(), i})
			}
		}
	} else {
		for i := sl1Len - 1; i >= 0; i-- {
			found = false
			for j := sl2Len - 1; j >= 0; j-- {
				if reflect.DeepEqual(lst1.Index(i).Interface(), lst2.Index(j).Interface()) {
					found = true
				}
			}
			if !found {
				del = append(del, SliceDiffElem{lst1.Index(i).Interface(), i})
			}
		}
		for i := sl2Len - 1; i >= 0; i-- {
			found = false
			for j := sl1Len - 1; j >= 0; j-- {
				if reflect.DeepEqual(lst2.Index(i).Interface(), lst1.Index(j).Interface()) {
					found = true
				}
			}
			if !found {
				ins = append(ins, SliceDiffElem{lst2.Index(i).Interface(), i})
			}
		}
	}
	return
}
