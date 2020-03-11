func main() {
	test1 := IsAnagram("a", "a")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("abbc", "    a b bcс   ")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

func isAnagram(str1, str2 string) bool {
	str := str1 + str2
	sum := byte(0)
	for _, v := range str {
		if v == ' ' {
			continue
		}
		sum ^= byte(v)
	}
	//fmt.Println(string(sum))
	return sum == 0
}
