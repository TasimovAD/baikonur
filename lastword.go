func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		var start, end int
		var slice string
		str := os.Args[1]
		for i := len(str) - 1; i != 0; i-- {
			if end == 0 && str[i] != ' ' {
				end = i + 1
			}
			if end > 0 && str[i] == ' ' {
				start = i + 1
			}
		}
		slice = str[start:end]
		for _, r := range slice {
			z01.PrintRune(r)
		}
	}
}
