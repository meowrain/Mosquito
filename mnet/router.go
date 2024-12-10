package mnet

import "mosquito/miface"

/*
BaseRouter  实现router时，先嵌入BaseRouter基类，根据需求对这个基类方法重写
之所以BaseRouter的方法都为空，是因为有的Router不希望有PreHandle和PostHandle这两个业务
*/
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request miface.IRequest) {

}

func (br *BaseRouter) Handle(request miface.IRequest) {
}

func (br *BaseRouter) PostHandle(request miface.IRequest) {

}
