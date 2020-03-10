func ReverseBits2(oct byte) byte {
    var buf [8] byte
    var result byte = 0
    for i := 0; i < 8; i++{
        buf[i] = oct % 2
        oct /= 2
    }
    for i := 0; i < 8; i++{
        result += buf[i] * byte(pow2(7-i))
    }
    fmt.Println(buf)
    return result
}

func pow2(x int) int{
    var result = 1
    for i := 0; i < x; i++ {
        result *= 2
    }
    return result
}
