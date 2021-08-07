package collection

import "reflect"

// StrMap описывает хранение произвольных ключей с их строковыми значениями.
type StrMap map[string]string

// NewStrMap создает новый объект StrMap.
func NewStrMap() *StrMap {
	return (*StrMap)(&map[string]string{})
}

// Add добавляет несуществующий ключ.
func (m *StrMap) Add(key, val string) {
	if !m.Exists(key) {
		(*m)[key] = val
	}
}

// Delete удаляет ключ.
func (m *StrMap) Delete(key string) {
	delete(*m, key)
}

// Exists проверяет наличие ключа.
func (m *StrMap) Exists(key string) bool {
	_, ok := (*m)[key]
	return ok
}

// Value возвращает значение ключа или nil, если он не найден.
func (m *StrMap) Value(key string) string {
	return (*m)[key]
}

// IsEmpty проверяет коллекцию на пустоту.
func (m *StrMap) IsEmpty() bool {
	return len(*m) == 0
}

// Clean сбрасывает всю коллекцию в nil, если поля map не содержит ключей.
func (m *StrMap) Clean() {
	if len(*m) == 0 {
		m = nil
	}
}

// MapDiff сравнивает 2 объекта map и возвращает изменения по отношению первого объекта.
// Возвращает изменения для аналогичных ключей, удаленные и новые элементы.
// В таком же порядке и рекомендуется проводить изменения на целевом объекте.
// Если исходные типы объектов не одинаковы, возвращаются пустые структуры данных по изменениям.
func MapDiff(map1 interface{}, map2 interface{}) (
	upd map[interface{}]interface{}, del map[interface{}]interface{},
	ins map[interface{}]interface{}) {
	ins = map[interface{}]interface{}{}
	upd = map[interface{}]interface{}{}
	del = map[interface{}]interface{}{}
	if reflect.TypeOf(map1) != reflect.TypeOf(map2) {
		return upd, del, ins
	}
	var found bool
	m1 := reflect.ValueOf(map1)
	m2 := reflect.ValueOf(map2)
	for _, key1 := range m1.MapKeys() {
		k1 := key1.Interface()
		found = false
		for _, key2 := range m2.MapKeys() {
			k2 := key2.Interface()
			if k1 == k2 {
				found = true
				if !reflect.DeepEqual(m1.MapIndex(key1).Interface(), m2.MapIndex(key2).Interface()) {
					upd[k1] = m2.MapIndex(key2).Interface()
				}
			}
		}
		if !found {
			del[k1] = m1.MapIndex(key1).Interface()
		}
	}
	for _, key2 := range m2.MapKeys() {
		k2 := key2.Interface()
		found = false
		for _, key1 := range m1.MapKeys() {
			k1 := key1.Interface()
			if k1 == k2 {
				found = true
			}
		}
		if !found {
			ins[k2] = m2.MapIndex(key2).Interface()
		}
	}
	return upd, del, ins
}
