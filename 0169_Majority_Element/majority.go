package  main

import "fmt"
import "sort"

// 使用排序
func majorityElement2(nums []int) int {
    length := len(nums)
    sort.Ints(nums)
    return nums[length/2]
}

func majorityElement(nums []int) int {
	count := 1
	majority := nums[0]
	for i := 1; i < len(nums); i++ {
		if count==0 {
			majority = nums[i]
		}
		if majority == nums[i] {
			count++
		}else{
			count--
		}
	}

	return majority
}

func main()  {
	var nums = []int{2,2,1,1,1,1,1,2,2}
	var res = majorityElement2(nums)
	fmt.Println(res)
}


