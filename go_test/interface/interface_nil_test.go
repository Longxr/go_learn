package interface_test

import (
	"fmt"
	"reflect"
	"testing"
)

type IPeople interface {
	hello()
}
type People struct {
}

func (p *People) hello() {
	fmt.Println("People hello()")
}

//错误：将nil的people给空接口后接口就不为nil,因为interface中的value为nil但type不为nil

func errFunc() *People {

	return nil
}

//正确处理返回nil给接口的方法,返回时go就确定了接口是不是nil
func rightFunc() IPeople {

	return nil
}
func TestInterfaceNil(t *testing.T) {

	var i IPeople
	i = errFunc()
	if i == nil { //想通过接口是否为nil来判断故障，却始终判断接口非空
		fmt.Println("errFunc对了哦，外部接收到也是nil")
		fmt.Println(reflect.TypeOf(i))
	} else {
		fmt.Println("errFunc错了咦，外部接收到不是nil哦")
		fmt.Println(reflect.TypeOf(i))
	}

	i = rightFunc()
	if i == nil {
		fmt.Println("rightFunc对了哦，外部接收到也是nil")
		fmt.Println(reflect.TypeOf(i))
	} else {
		fmt.Println("rightFunc错了咦，外部接收到不是nil哦")
		fmt.Println(reflect.TypeOf(i))
	}
}