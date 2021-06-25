package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
}
