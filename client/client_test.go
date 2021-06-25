package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSetFuncField(t *testing.T) {
	//assert.Nil(t, hello{}, "dsfs")

	s := []interface{}{0, 1, 2, 3, 5}

	newAdd, _ := Add(s, 4, 5)
	fmt.Println(newAdd)
	newDelete, _ := Delete(s, 5)
	fmt.Println(newDelete)

	u := &User{}
	v := reflect.ValueOf(u).Elem()
	v.FieldByName("Name").SetString("tome")
	v.FieldByName("Age").SetInt(18)
	v.FieldByName("Say").Set(reflect.MakeFunc(v.FieldByName("Say").Type(), func(args []reflect.Value) (results []reflect.Value) {
		fmt.Println("hello world")
		return []reflect.Value{reflect.ValueOf([]interface{}{})}
	}))
	res := u.Say()
	fmt.Println(res)
	fmt.Println(u)
}

// 指定索引添加元素
func Add(values []interface{}, value interface{}, index int) ([]interface{}, error) {
	if index < 0 || index >= len(values) {
		return nil, errors.New("index 非法")
	}
	var res []interface{}
	for i := 0; i < index; i++ {
		res = append(res, values[i])
	}
	res = append(res, value)

	res = append(res, values[index:]...)
	return res, nil
}

// 指定索引删除元素
func Delete(values []interface{}, index int) ([]interface{}, error) {
	if index < 0 || index >= len(values) {
		return nil, errors.New("index 非法")
	}
	var res []interface{}
	res = append(values[:index], values[index+1:]...)
	return res, nil
}

type User struct {
	Name string
	Age  int
	Say  func() []interface{}
}
