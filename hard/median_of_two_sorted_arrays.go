package main

import (
	"fmt"
	"strconv"
)

/*给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000*/

func medianOfTwoSortedArrays(a, b []int) float64 {
	var c []int
	var sum = 0
	c = append(a, b...)
	for _, v := range c {
		sum = sum + v
	}
	return float64(sum) / float64(len(c))
}

func main() {
	out,_:=strconv.ParseFloat(fmt.Sprintf("%.4f", medianOfTwoSortedArrays([]int{1, 2, 3, 4, 5}, []int{4})), 64)
	fmt.Println(out)
}
