func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 3 {
		str1 := os.Args[1]
		str2 := os.Args[2]
		i := 0
		for j := range str2 {
			if str1[i] == str2[j] {
				i++
				if i == len(str1) {
					z01.PrintRune('1')
					return
				}
			}
		}
		z01.PrintRune('0')
	}
}
