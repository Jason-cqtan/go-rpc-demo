package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type HelloService interface {
	SayHello(name string) (string error)
}

type hello struct {
	host string
}

func (h hello) SayHello(name string) (string, error) {
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

func PrintParams(val interface{}) {
	// 反射
	t := reflect.TypeOf(val)
	num := t.NumMethod()
	for i := 0; i < num; i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
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
	h := hello{
		host: "http://localhost:8080/",
	}
	str, err := h.SayHello("golang")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)

	PrintParams(h)
}
