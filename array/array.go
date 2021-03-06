package array

import "fmt"

func Testarray(desc string) {
	var Arrayone [3]int
	fmt.Printf("ArrayOne start is %v \n", Arrayone)

	Arrayone[0] = 1
	Arrayone[1] = 2
	fmt.Printf("Array one is %v \n", Arrayone)

	Arraytwo := [...]int{11, 22, 33}
	fmt.Printf("ArrayTwo is %v \n", Arraytwo)

	a := [3]string{"China", "UK", "US"}
	b := a
	b[0] = "NewChina"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

}

func Changearray(num [5]string) {
	//for and range is good at traversaling array
	for i := 0; i < len(num); i++ {
		fmt.Printf("num[%d] is %s \n", i, num[i])
	}
	for i, v := range num {
		fmt.Printf("num[%d] is %s \n", i, v)

	}
	num[0] = "NewChina"
	fmt.Printf("num is %v \n", num)
}

func Printarrays(num [2][3]int) {
	//  for or range can traversal array
	for i1, v1 := range num {
		for i2, v2 := range v1 {
			fmt.Printf("num[%d][%d] is %v\n", i1, i2, v2)
		}
	}

	// array can be created by [...] instead of [5]
	Array := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array is %v\n", Array)

	Slice := Array[0:3]
	for i, _ := range Slice {
		Slice[i]++
	}
	fmt.Printf("After slice++ Array is %v\n", Array)

	// slice'cap is the array len
	fmt.Printf("len(Slice) is %d, and cap(Slice) is %d \n", len(Slice), cap(Slice))

	slicecap := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(slicecap) is %d, and cap(slicecap) is %d \n", len(slicecap), cap(slicecap))
	slicecap = append(slicecap, 3)
	//the cap of the slice will be the double
	fmt.Printf("after append a element,len(slicecap) is %d, and cap(slicecap) is %d \n", len(slicecap), cap(slicecap))

	var slicenil []int
	if slicenil == nil {
		fmt.Println("slicenil is nil")
	}
	slicenil = append(slicenil, 1, 2)
	fmt.Println(slicenil)
	//... append one slice to another slice
	slicenil = append(slicenil, slicecap...)
	fmt.Println(slicenil)

}

func SlicewillChange(slice []int) {
	for i := range slice {
		slice[i]++
	}
}
