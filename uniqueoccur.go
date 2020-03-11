func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		str := os.Args[1]
		m := map[rune]int{}
		res := ""
		for _, r := range str {
			m[r]++
		}
		fmt.Println(m)
		for key, count := range m {
			if isUniq(m, key, count) {
				toPrint("true")
				res = "true"
			} else {
				toPrint("false")
				return
			}
		}
		toPrint(res)
	}
}
func toPrint(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
func isUniq(m map[rune]int, key rune, count int) bool {
	for s, num := range m {
		if s != key {
			if num == count {
				return false
			}
		}
	}
	return true
}
