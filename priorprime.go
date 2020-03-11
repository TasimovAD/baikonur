func main() {
	fmt.Println(Prioprime(14))
}
func Prioprime(x int) int {
	res := 0
	for i := 2; i < x; i++ {
		if Prime(i) {
			fmt.Println(i)
			res += i
		}
	}
	return res
}
func Prime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
