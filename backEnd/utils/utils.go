package utils

import (
	"fmt"
	"reflect"
	"errors"
)

func setField(obj interface{}, name string, value interface{}) error{
	structData := reflect.ValueOf(obj).Elem()

	fieldValue := structData.FieldByName(name)
	if !fieldValue.IsValid(){
		return fmt.Errorf("utils.setField() No such field: %s in obj", name)
	}
	if !fieldValue.CanSet(){
		return fmt.Errorf("cannot set %s field value", name)
	}
	fieldType := fieldValue.Type()
	val := reflect.ValueOf(value)
	valTypeStr := val.Type().String()
	fieldTypeStr := fieldType.String()
	if valTypeStr == "float64" && fieldTypeStr == "int"{
		val = val.Convert(fieldType)
	}else if fieldType != val.Type(){
		return errors.New("Provided value type " + valTypeStr + " didn't match obj field type " + fieldTypeStr)
	}
	fieldValue.Set(val)
	return nil
}

func SetStructByJSON(obj interface{}, mapData map[string] interface{}) error{
	for key, value := range mapData{
		fmt.Println("----------------")
		fmt.Println(key)
		fmt.Println(value)
		if err := setField(obj, key, value); err != nil{
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}