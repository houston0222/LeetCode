func twoSum(nums []int, target int) []int {
    hashMap := map[int]int{}

    for i, v := range nums {
        residual := target - v
        if j, found := hashMap[residual]; found {
			return []int{j, i}
		}


        hashMap[v] = i
    }

    return nil // should never reach this
    
}