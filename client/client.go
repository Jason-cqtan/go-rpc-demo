package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	ycp, err := NewYamlConfigProvider("")
	if err != nil {
		panic("配置文件未找到")
	}

	err = InitApplication(WithCfgProvider(ycp))
	if err != nil {
		panic("初始化应用失败")
	}

}

func SetFuncField(val Service) {
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

				serviceName := val.ServiceName()
				cfg, err := App.CfgProvider.GetServiceConfig(serviceName)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				req, err := http.NewRequest("POST", cfg.Endpoint, bytes.NewReader(inData))
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("sparrow-service", serviceName)
				req.Header.Set("sparrow-service-method", field.Name)

				resp, err := client.Do(req)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				data, err := ioutil.ReadAll(resp.Body)
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

type Service interface {
	ServiceName() string
}
