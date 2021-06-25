package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type HelloService interface {
	SayHello(name string) (string error)
}

type hello struct {
	host      string
	FuncField func(in *Input) (*Output, error)
}


type Input struct {
	Name string
}

type Output struct {
	Msg string
}


func (h *hello) SayHello(name string) (string, error) {
	client := http.Client{}
	r, err := client.Get(h.host + name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(s), nil
}

func SetFuncField(val interface{}) {
	// 反射
	// TypeOf 获得对象的类型信息，例如该类型（结构体）有啥字段，字段是啥类型
	// ValueOf 获得对象运行时表示，例如有啥字段，字段的值是啥
	//t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	ele := v.Elem() // 指针指向的结构体
	t := ele.Type() // 指针指向结构体的类型信息
	// NumMethod只能返回公共方法
	num := t.NumField()
	for i := 0; i < num; i++ {
		field := t.Field(i)
		fieldValue := ele.Field(i) // 用指针指向的结构体来访问
		if fieldValue.CanSet() {

			fn := func(args []reflect.Value) (result []reflect.Value) {
				in := args[0].Interface()
				out := reflect.New(field.Type.Out(0).Elem()).Interface()
				inData, err := json.Marshal(in)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				client := http.Client{}
				r, err := client.Post("http://localhost:8080/", "application/json", bytes.NewReader(inData))
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				err = json.Unmarshal(data, out)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				return []reflect.Value{reflect.ValueOf(out), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}

			fieldValue.Set(reflect.MakeFunc(field.Type, fn))

		}
	}
}

// 假如有个雷同的方法，需要修改上面方法某些参数
// 最直接是复制一遍新写（不推荐）
// 另一种方法是利用反射
// 反射一种运行时获得一些运行程序本身信息的机制
//
//首先，获得方法原本的信息
//
//其次，将方法的内容，改为http调用的内容

func main() {
	h := &hello{
		host: "http://localhost:8080/",
	}
	str, err := h.SayHello("golang")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)

	SetFuncField(h)
	msg, _ := h.FuncField(&Input{
		Name: "reflect",
	})
	fmt.Println(msg)

}
