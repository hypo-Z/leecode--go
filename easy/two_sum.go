package main

import "fmt"

/*给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/
func twoSum(nums []int, target int) []int {
	num2index := make(map[int]int, len(nums))
	for i, num := range nums {
		pair := target - num //核心是用目标值减去循环遍历出来的值，判断是否是数组内的值
		if j, ok := num2index[pair]; ok && i != j {
			return []int{j, i} // 注意返回值顺序，向后遍历 nums，所以 i 在 j 后
		}
		num2index[num] = i
	}
	return nil
}

func main() {
	var in = []int{3, 3, 12, 234, 34, 12, 3, 23, 21, 5, 68, 654, 334, 234, 0, 45, 76, 67, 54, 89, 77, 6, 4}
	out := twoSum(in, 234)
	fmt.Println(in[out[0]], in[out[1]])
}
