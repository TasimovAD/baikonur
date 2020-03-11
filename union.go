func main() {
	defer z01.PrintRune('\n')
	var res string
	if len(os.Args) == 3 {
		join := os.Args[1] + os.Args[2]
		for _, r := range join {
			if isUniq(res, r) {
				res += string(r)
			}
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}
func isUniq(str string, k rune) bool {
	for _, r := range str {
		if r == k {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		z01.PrintRune('\n')
		return
	}
	arg1 := arg[0]
	arg2 := arg[1]
	m := map[rune]int{}
	for _, value := range arg1 + arg2 {
		if m[value] < 1 {
			m[value]++
			z01.PrintRune(rune(value))
		}
	}
	z01.PrintRune('\n')
}
