package a

import (
	"errors"
	aa "github.com/pkg/errors"
	"os"
)

func A() error {

	return aa.Wrap(errors.New("haaaa"), "aaa123")
}

func B() error {
	err := A()
	if err != nil {

		return aa.Wrap(err, "A failed")
	}
	return nil
}

func C() error {
	//e := &errorString{s: "34344545"}
	//println(aa.As(E(), &e))
	//_ = fmt.Sprintf("%+v", E())
	return E()
}

func D() error {
	//e:= &errorString{
	//	"gfhfghjgjh",
	//}
	//return aa.New(e.Error())
	_, err := os.Open("78")
	return err
}

func E() error {
	return aa.Wrap(D(), "ee")
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

//go进行error处理时，第三方调用报错， 进行封装， 中间直接传递， 顶层+v打印
//Is 就是判断最底层的err
//As 就是判断错误类型
