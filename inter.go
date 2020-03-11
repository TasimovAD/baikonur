func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		var res, fin string

		for _, b := range os.Args[1] {
			for _, r := range os.Args[2] {
				if r == b {
					res += string(r)
				}
			}
		}
		for _, r := range res {
			if isUniq(r, fin) {
				fin = fin + string(r)
			}
		}
		for _, r := range fin {
			z01.PrintRune(r)
		}
	}
}

func isUniq(r rune, str string) bool {
	for _, b := range str {
		if r == b {
			return false
		}
	}
	return true
}

func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		str := map[rune]int{}
		for _, r := range os.Args[1] {
			for _, b := range os.Args[2] {
				if r == b && str != 1 {
					str[r]++
					fmt.Println(string(str))
				}
			}
		}
	}
}
