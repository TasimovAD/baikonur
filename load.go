package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//a, _ := exec.Command("curl", "https://raw.githubusercontent.com/adilSmile/piscine-go/master/sortparams/compare.go").Output()
	a, _ := exec.Command("curl", "https://raw.githubusercontent.com/TasimovAD/baikonur/master/hello.go").Output()
	fmt.Println(string(a))
}
