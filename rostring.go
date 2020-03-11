func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		res := swap(splitWS(os.Args[1]))
		res1 := ""
		for i, r := range res {
			if i != len(res)-1 {
				res1 = res1 + r + " "
			} else {
				res1 = res1 + r
			}
		}
		for _, b := range res1 {
			z01.PrintRune(b)
		}
	}
}

func swap(arg []string) []string {
	counter := 0
	for i := range arg {
		counter++
		if counter != len(arg) {
			arg[i], arg[counter] = arg[counter], arg[i]
		}
	}
	return arg
}
func splitWS(str string) []string {
	var s string
	count := 0
	n := 0
	str = str + " "
	for i := range str {
		if str[i] == ' ' {
			s = str[n:i]
			n = i + 1
			if s != "" {
				count++
			}
		}
	}
	n = 0
	arr := make([]string, count)
	j := 0
	for i := range str {
		if str[i] == ' ' {
			s = str[n:i]
			n = i + 1
			if s != "" {
				arr[j] = s
				j++
			}
		}
	}
	return arr
}
