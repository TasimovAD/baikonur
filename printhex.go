func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		base := "0123456789abcdef"
		res := ""
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			z01.PrintRune('0')
		}
		for num != 0 {
			res = string(base[num%16]) + res
			num /= 16
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}
