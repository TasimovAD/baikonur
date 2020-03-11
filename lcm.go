func main() {
	fmt.Println(Lcm(7, 34))
}
func Lcm(first, second int) int {
	res := 0
	for i := second; i <= first*second; i++ {
		if i%first == 0 && i%second == 0 {
			res = i
			break
		}
	}
	return res
}
