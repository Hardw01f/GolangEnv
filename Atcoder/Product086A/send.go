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



func main(){
		lineInputs := InputString()
		splited := strings.Split(lineInputs," ")
		index_one, err := strconv.Atoi(splited[0])
		if err != nil{
				fmt.Println("error")
		}
		index_two,err := strconv.Atoi(splited[1])
		if err != nil{
				fmt.Println("error")
		}

		if ((index_one*index_two)%2 == 0){
				fmt.Println("Even")
		}else{
				fmt.Println("Odd")
		}

}
