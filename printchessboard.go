func main() {
	err := "Error"
	if len(os.Args) == 3 {
		x, _ := strconv.Atoi(os.Args[1])
		y, _ := strconv.Atoi(os.Args[2])
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i%2 == 0 {
					if j%2 == 0 {
						z01.PrintRune('#')
					} else {
						z01.PrintRune(' ')
					}
				} else {
					if j%2 == 0 {
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('#')
					}
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		for _, r := range err {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
