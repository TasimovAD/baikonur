func main() {
	fmt.Println(FirstRune("firstruna"))
}
func FirstRune(str string) (res rune) {
	count := 0
	for range str {
		count++
	}
	return rune(str[count-1])
}
