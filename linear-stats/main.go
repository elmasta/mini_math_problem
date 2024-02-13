package main

import (
	"os"
	"io/ioutil"
	"strings"
	calc "calc/calc"
)

func main() {
	if len(os.Args) == 2 {
		datab, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			os.Stdout.Write([]byte("ERROR: open " + os.Args[1] + ": no such file or directory\n"))
			os.Exit(1)
		} else {
			data := strings.Split(string(datab), "\n")
			calc.LinearR(data)
		}
	}
}