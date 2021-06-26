package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Invoker interface {
	Invoke(inv *Invocation) ([]byte, error)
}

type Invocation struct {
	MethodName string
	ServiceName string
	Input []byte
}

type httpInvoker struct {
	
}

func (h *httpInvoker) Invoke(inv *Invocation) ([]byte, error) {
	serviceName := inv.ServiceName
	methodName := inv.MethodName
	data := inv.Input
	service, err := GetService(serviceName)
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(service)

	method := val.MethodByName(methodName)
	inType := method.Type().In(0)
	in := reflect.New(inType.Elem())
	err = json.Unmarshal(data, in.Interface())
	if err != nil {
		return nil, err
	}
	res := method.Call([]reflect.Value{in})

	output, err := json.Marshal(res[0].Interface())
	if err != nil {
		return nil, err
	}
	return output, nil
}

// 也是一个装饰器模式，我们会在这里组织filter
type filterInvoker struct {
	Invoker
	filters []Filter
}

func (f *filterInvoker)Invoke(inv *Invocation) ([]byte, error) {
	for _, flt := range f.filters {
		flt(inv)
	}
	return f.Invoker.Invoke(inv)
}

// 通过扩展 filter 来完成别的工作
type Filter func(inv *Invocation)

func logFilter(inv *Invocation) {
	fmt.Printf("log filter ===== service name: %s, method name: %s \n", inv.ServiceName, inv.MethodName)
}