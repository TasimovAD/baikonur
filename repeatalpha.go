func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		str := os.Args[1]
		for _, r := range str {
			if r >= 65 && r <= 90 {
				for i := 0; i <= int(r)-65; i++ {
					z01.PrintRune(r)
				}
			}
			if r >= 97 && r <= 122 {
				for i := 0; i <= int(r)-97; i++ {
					z01.PrintRune(r)
				}
			} else {
				z01.PrintRune(r)
			}
		}
	}
}
