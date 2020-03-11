func main() {
	defer z01.PrintRune('\n')
	if len(os.Args) == 2 {
		var res string
		str := os.Args[1]
		for i := len(str) - 1; i != 0; i-- {
			if str[i] == ' ' {
				res += str[i+1:] + " "
				str = str[:i]
			}
		}
		res += str
		for _, r := range res {
			z01.PrintRune(r)
		}
	}
}
