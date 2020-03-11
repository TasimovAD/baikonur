func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 2))
}
func Nauuo(plus, minus, rand int) string {
	var res string
	if plus == minus && rand == 0 {
		res = "0"
	} else if plus > (minus + rand) {
		res = "+"
	} else if (plus + rand) < minus {
		res = "-"
	} else {
		res = "?"
	}
	return res
}
