package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSetFuncField(t *testing.T) {

	path := `../test/client.yaml`
	ycp, _ := NewYamlConfigProvider(path)

	_ = InitApplication(WithCfgProvider(ycp))

	helloService := &hello{
	}

	SetFuncField(helloService)

	res, err := helloService.SayHello(&Input{
		Name: "golang",
	})

	assert.Nil(t, err)
	assert.Equal(t, "Hello, golang", res.Msg)
}

type hello struct {
	SayHello func(in *Input) (*Output, error)
}

func (h hello) ServiceName() string {
	return "hello"
}

type Input struct {
	Name string
}

type Output struct {
	Msg string
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

type HashSet interface {
	Set(key string)
	Size() int
	Exist(key string) bool
}

type hashset struct {
	m map[string]interface{}
}

func (h *hashset) Set(key string) {
	h.m[key] = ""

}

func (h *hashset) Size() int {
	// len array slice map
	// cap array slice channel
	// for range array slice map
	return len(h.m)
}

func (h *hashset) Exist(key string) bool {
	_, ok := h.m[key]
	return ok
}

// 装饰器，hashset接口
type safeset struct {
	HashSet
	mutex sync.RWMutex
}

func (s *safeset) Size() int {
	s.mutex.RLocker()
	defer s.mutex.RUnlock()
	return s.HashSet.Size()
}

func (s *safeset) Set(key string) {
	s.mutex.Unlock()
	defer s.mutex.Unlock()
	s.HashSet.Set(key)
}

func (s *safeset) Exist(key string) bool {
	s.mutex.RUnlock()
	defer s.mutex.RUnlock()
	return s.HashSet.Exist(key)
}
