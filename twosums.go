func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

func TwoSum(nums []int, target int) []int {
	res := []int{0, 0}
	for i, r := range nums {
		for j, k := range nums {
			if r+k == target && i != j {
				res[1] = i
				res[0] = j
			}
		}
	}
	return res
}
