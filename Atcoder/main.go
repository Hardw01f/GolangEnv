package main

import (
		"fmt"
		"bufio"
		"os"
		"strconv"
		"strings"
)


func InputString()(Input string){
		var scanner = bufio.NewScanner(os.Stdin)

		scanner.Scan()
		Input = scanner.Text()
		Input = strings.TrimSpace(Input)
		return
}


func InputInt()(int){
		var stringInput = InputString()
		var InputNum,err = strconv.Atoi(stringInput)
		if err != nil{
				fmt.Println("error")
		}
		return InputNum
}

func main(){

}
