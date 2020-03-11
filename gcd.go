func main() {
	if len(os.Args) == 3 {
		res := 0
		num1, _ := strconv.Atoi(os.Args[1])
		num2, _ := strconv.Atoi(os.Args[2])
		for i := num1; i > 0; i-- {
			if num1%i == 0 && num2%i == 0 {
				res = i
				break
			}
		}
		fmt.Println(res)
	} else {
		fmt.Println()
		return
	}
}
