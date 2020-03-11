func main() {
	defer z01.PrintRune('\n')
	result := "false"
	var way int
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r == 'U' || r == 'R' {
				way++
			}
			if r == 'D' || r == 'L' {
				way--
			}
		}
		if way == 0 {
			result = "true"
		}
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
}
