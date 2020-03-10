package main

func main(){
}

func Itoa(n int) string {
    isNegative:=false
    if n==0 {
        return "0"
    }

    if n<0 { 
        isNegative=true
    }
    var a int
    var res string
    for n!=0 {

        if !isNegative {
            a=n%10+48
            n=n/10
            res=string(rune(a))+res
        } else {
            a=(n%10)*(-1)+48
            n=n/10
            res=string(rune(a))+res
        }
        
    }

    if isNegative {
        res="-"+res
    }
    return res
}
