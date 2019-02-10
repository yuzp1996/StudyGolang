package playerror

import (
	"fmt"
)

func ErrorNew(info string)error{
	//return errors.New(info)
	return fmt.Errorf("Error is [%s]", info)
}



type AreaError struct{
	length float64
	width float64
	err string
}

func (e *AreaError)Error() string{
	return e.err
}

func (e *AreaError)lengthNegative()bool{
    return e.length < 0
}

func(e *AreaError)widthNegative()bool{
	return e.width < 0
}




func rectArea(width, length float64)(float64, error){
	err := ""
	if width < 0 {
		err = "width is below zero"
	}
	if err == ""{
		if length < 0{
			err = "length is below zero"
		}
	}else{
		err += ", length is below zero"
	}
	if err != ""{
		return 0, &AreaError{length:length,width:width,err:err}
	}
	return width*length, nil

}

func MainCoustmerError(){
	width,length := -1.2, -2.4
	area, err := rectArea(width,length)
	if err != nil{
		if Err,ok := err.(*AreaError); ok{
			if Err.lengthNegative(){
				fmt.Printf("err is length is below zero and length is %v\n",Err.length)
			}
			if Err.widthNegative(){
				fmt.Printf("err is width below zero and width is %v\n", Err.width)
			}
			fmt.Println(Err.err)
		}
		return
	}
	fmt.Printf(" area is %v ", area)
}











