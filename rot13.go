func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if (r >= 65 && r <= 77) || (r >= 97 && r <= 109) {
				z01.PrintRune(r + 13)
			} else if (r >= 78 && r <= 92) || (r >= 110 && r <= 122) {
				z01.PrintRune(r - 13)
			} else {
				z01.PrintRune(r)
			}
		}
	}
}
