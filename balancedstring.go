func main() {
	if len(os.Args) == 2 {
		var count, res int
		for _, r := range os.Args[1] {
			if r == 'C' {
				count++
			} else {
				count--
			}
			if count == 0 {
				res++
			}
		}
		fmt.Println(res)
	} else {
		fmt.Println()
		return
	}
}
