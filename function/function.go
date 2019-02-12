package function

import (
	"fmt"
)

// 如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数
func Multipama(num int, nums ...int) {
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Printf("find it it is [%v]inclued in nums and it is %v \n", i, v)
			found = true
		}
	}
	if found {
		fmt.Printf("find (%v) \n", num)
	}
}

func Change(nums ...string) {
	nums[0] = "Go"

	fmt.Printf("num is %p and nums[0] is %p len is %d and cap is %d  \n", &(nums), &(nums[0]), len(nums), cap(nums))

	//这里相当于重建了一个
	nums = append(nums, "Haha")
	fmt.Println(nums)
	fmt.Printf("nums is %p  nums[0]is %p and len is %d  cap is %d\n", &(nums), &(nums[0]), len(nums), cap(nums))
	fmt.Println("next is main")
}
