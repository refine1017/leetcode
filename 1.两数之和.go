/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var num = nums[i]
		var need = target - num

		if _, found := m[need]; found {
			return []int{m[need], i}
		}

		m[num] = i
	}

	return []int{0, 0}
}

// @lc code=end

