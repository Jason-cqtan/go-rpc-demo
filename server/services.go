package main

import (
	"errors"
	"fmt"
	"sync"
)

type Service interface {
	ServiceName() string
}

var services sync.Map

func AddService(service Service) {
	services.Store(service.ServiceName(), service)
}

var ErrServiceNotFound = errors.New("service not found")

func GetService(name string) (Service, error) {
	if service, ok := services.Load(name); ok {
		return service.(Service), nil
	}
	return nil, ErrServiceNotFound
}

type HelloService interface {
	Service
	SayHello(input *Input) (*Output, error)
}

type UserService interface {
	Service
	GetUser(req *GetUserReq) (*GetUserResp, error)
}

type userService struct {
	
}

type helloService struct {

}

func (h *helloService) SayHello(input *Input) (*Output, error) {
	fmt.Printf("Say Hello From " + input.Name)
	return &Output{
		Msg: "Hello, " + input.Name,
	}, nil
}

func (h *helloService) ServiceName() string {
	return "hello"
}

func (u *userService) GetUser(req *GetUserReq) (*GetUserResp, error) {
	return &GetUserResp{
		Id: req.Id,
		Name: fmt.Sprintf("mock_name_%d", req.Id),
	}, nil
}

func (u *userService) ServiceName() string {
	return "user"
}

type GetUserReq struct {
	Id int64
}

type GetUserResp struct {
	Id int64
	Name string
}

type Input struct {
	Name string
}

type Output struct {
	Msg string
}



