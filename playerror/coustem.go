package playerror

import (
	"fmt"
)

//使用 Errorf 给错误添加更多信息
func ErrorNew(info string) error {
	//return errors.New(info)
	//在这里把错误包装了一下
	return fmt.Errorf("Error is [%s]", info)
}

///////////////////

//自定义错误在这里会被调用

func MainCoustmerError() {
	width, length := -1.2, -2.4
	area, err := rectArea(width, length)
	if err != nil {
		fmt.Println(err)
		if Err, ok := err.(*AreaError); ok {
			if Err.lengthNegative() {
				fmt.Printf("err is length is below zero and length is %v\n", Err.length)
			}
			if Err.widthNegative() {
				fmt.Printf("err is width below zero and width is %v\n", Err.width)
			}
			fmt.Println(Err.err)
		}
		return
	}
	fmt.Printf(" area is %v ", area)
}

func rectArea(width, length float64) (float64, error) {
	err := ""
	if width < 0 {
		err = "width is below zero"
	}
	if err == "" {
		if length < 0 {
			err = "length is below zero"
		}
	} else {
		err += ", length is below zero"
	}
	if err != "" {
		//使用下面的那个错误
		return 0, &AreaError{length: length, width: width, err: err}
	}
	return width * length, nil

}

//错误类型的命名约定是名称以 Error 结尾  自定义了一个错误类型
type AreaError struct {
	length float64
	width  float64
	err    string
}

func (e *AreaError) Error() string {
	return e.err
}

func (e *AreaError) lengthNegative() bool {
	return e.length < 0
}

func (e *AreaError) widthNegative() bool {
	return e.width < 0
}
