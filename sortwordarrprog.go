func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
}
func SortWordArr(array []string) {
	count := 0
	for range array {
		count++
	}
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
