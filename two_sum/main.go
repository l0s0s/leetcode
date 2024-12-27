package main

func twoSumEz(nums []int, target int) []int {
	for i, num := range nums {
		for j, num2 := range nums[i+1:] {
			if num+num2 == target {
				return []int{i, j + i + 1}
			}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	ells := map[int]int{}

	for i, num := range nums {
		if j, ok := ells[target-num]; ok {
			return []int{j, i}
		}

		ells[num] = i
	}

	return []int{}
}
