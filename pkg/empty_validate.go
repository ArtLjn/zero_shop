package pkg

import (
	"crypto/md5"
	"fmt"
	"reflect"
)

func isNilOrEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	}
	return false
}

func CheckFieldsNoEmpty(s interface{}) error {
	rv := reflect.ValueOf(s)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if isNilOrEmpty(field) {
			return fmt.Errorf("field %s is nil or empty", rv.Type().Field(i).Name)
		}
	}
	return nil
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
