package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	serviceName := r.Header.Get("sparrow-service")
	methodName := r.Header.Get("sparrow-service-method")

	filterIvk := &filterInvoker{
		Invoker: &httpInvoker{},
		filters: []Filter{logFilter},
	}
	output, _ := filterIvk.Invoke(&Invocation{
		MethodName:  methodName,
		ServiceName: serviceName,
		Input:       data,
	})
	fmt.Fprintf(w, "%s", string(output))
}

// 启动服务器
func main() {
	AddService(&userService{})
	AddService(&helloService{})
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
