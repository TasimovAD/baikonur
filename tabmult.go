func main() {
	var res string
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		for i := 1; i <= 9; i++ {
			res = res + iToA(i) + " x " + iToA(num) + " = " + iToA(num*i) + "\n"
		}
		for _, r := range res {
			z01.PrintRune(r)
		}
	} else{
		z01.PrintRune('\n')
	}

}
