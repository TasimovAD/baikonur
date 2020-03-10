package main

func main() {

}

func AtoiBase(input string, base string) int {
    if !inputValid(input, base) {
        return 0
    }
    var buffer int = 0
    for i := 0; i < len(input); i++ {
        buffer += pos(input[len(input)-i-1], base) * pow(len(base), i)
    }
    return buffer
}

func inputValid(input string, base string) bool {
    if len(base) < 2 {
        return false
    }
    if !isUnique(base) {
        return false
    }
    if contains(base, '+') || contains(base, '-'){
        return false
    }
    return true
}

func isUnique(s string) bool {
    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s); j++ {
            if s[i] == s[j] {
                return false
            }
        }
    }
    return true
}

func contains(s string, char uint8) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == char {
            return true
        }
    }
    return false
}

func pos(u uint8, s string) int {
    for i := 0; i < len(s); i++ {
        if s[i] == u {
            return i
        }
    }
    return -1
}

func pow(x int, base int) int {
    result := 1
    for i := 0; i < base; i++ {
        result *= x
    }
    return result
}
