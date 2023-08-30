package services

import "sync"

var once sync.Once
type BaseService struct {
}


//++++++++++++++++
//service构造工厂，单例模式
//+++++++++++++++

func NewBook() *BookService {
	var obj *BookService
	once.Do(func() {
		obj = &BookService{}
	})

	return obj
}