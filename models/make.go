package models

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// DBField 数据库字段过滤
func DBFiled(reflectType reflect.Type, buffer *bytes.Buffer) {
	if reflectType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < reflectType.NumField(); i++ {
		jsonTag := reflectType.Field(i).Tag.Get("json")
		dbTag := reflectType.Field(i).Tag.Get("db")

		if jsonTag == "" && dbTag == "" {
			DBFiled(reflectType.Field(i).Type, buffer)
			continue
		}

		dbProfile := strings.Split(dbTag, ";")
		dbFiled := fmt.Sprintf("%s %s", jsonTag, strings.Join(dbProfile, " "))
		buffer.WriteString(dbFiled)
		buffer.WriteString(",")
	}
}

// GenerateFiled 生成更新字段
func GenerateFiled(f interface{}) string {
	rType := reflect.TypeOf(f)
	rValue := reflect.ValueOf(f)
	rList := []string{}
	for i := 1; i < rValue.NumField(); i++ {
		filed := rType.Field(i)
		key := filed.Tag.Get("json")
		value := rValue.Field(i).Interface()
		rList = append(rList, fmt.Sprintf(`%s="%v"`, key, value))
	}
	return strings.Join(rList, ",")
}
