package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int{
	res:=0
	if base[0]=='-'{
		return 0
	} else {
		for _,i:=range s{
			for t,j:=range base{
				if i==j{
					res=res*len(base)+t
				}
			}
		}
		return res
	}
}
