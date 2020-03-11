func main() {
	if len(os.Args) == 4 {
		var res int
		num1, f := strconv.Atoi(os.Args[1])
		num2, s := strconv.Atoi(os.Args[3])
		if f != nil || s != nil {
			fmt.Println(0)
			return
		}
		if os.Args[2] == "+" {
			res = num1 + num2
			if num1 > 0 && num2 > 0 && res < 0 {
				res = 0
			}
		}
		if os.Args[2] == "-" {
			res = num1 - num2
			if num1 < 0 && num2 < 0 && res < 0 {
				res = 0
			}
		}
		if os.Args[2] == "%" {
			if num2 == 0 {
				fmt.Println("no modulo by 0")
				return
			} else {
				res = num1 % num2
			}
		}
		if os.Args[2] == "/" {
			if num2 == 0 {
				fmt.Println("no division by 0")
				return
			} else {
				res = num1 / num2
			}
		}
		if os.Args[2] == "*" {
			res = num1 * num2
			if res/num1 != num2 {
				res = 0
			}
		}
		fmt.Println(res)
	}
}
