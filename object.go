package drop

import "reflect"

type Object struct {
	data map[string]interface{}
}

func (o Object) Contain(key string, expect interface{}) bool {
	actual, ok := o.data[key]
	if !ok {
		return false
	}
	if !reflect.DeepEqual(actual, expect) {
		return false
	}
	return true
}
