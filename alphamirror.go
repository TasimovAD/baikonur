func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r >= 65 && r <= 90 {
				z01.PrintRune(90 - (r - 65))
			} else if r >= 97 && r <= 122 {
				z01.PrintRune(122 - (r - 97))
			} else {
				z01.PrintRune(r)
			}
		}
	}
}
