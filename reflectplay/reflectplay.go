package reflectplay

import (
	"fmt"
	"reflect"
)

//reflect 实现了运行时反射。reflect 包会帮助识别 interface{} 变量的底层具体类型和具体值
//https://studygolang.com/articles/13178
//应该尽量避免使用反射

func Mainreflect() {
	i := 10
	fmt.Printf("i is %v, and Type is %T\n", i, i)
}

type Order struct {
	ordId      int
	customerId int
}

func Insertquery(o Order) string {
	result := fmt.Sprintf("Instert to DB and o.ordId is %d and customerId is %d", o.ordId, o.customerId)
	return result
}

func CreateQeury(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	v := reflect.ValueOf(q)
	name := t.Name()
	fmt.Println("t and v is below ---->")
	fmt.Printf("type is %v\n", t)
	fmt.Printf("kind is %v\n", k)
	fmt.Printf("value us %v\n", v)
	fmt.Printf("name is %v\n", name)
}

func TryReflect() {
	o := Order{
		12,
		23,
	}
	result := Insertquery(o)
	fmt.Println(result)
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T, value: %v\n", x, x)
	CreateQeury(o)
	createquery(o)
	createquery(123)
}

func createquery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		fmt.Printf("Num of the filed is %d\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			default:
				fmt.Println("unsupport type")
				return
			}
			fmt.Printf("Field:%d, type:%T,  value:%v\n", i, v.Field(i), v.Field(i))
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}

}
