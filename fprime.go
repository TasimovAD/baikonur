func main() {
    arg := os.Args[1]
    divider := 2
    num, _ := strconv.Atoi(arg)
    for divider <= num {
        if num%divider == 0 {
            fmt.Print(divider)
            if divider == num {
                fmt.Println("")
                return
            }
            fmt.Print("*")
            num /= divider
            divider = 1
        }
        divider++
    }
}
