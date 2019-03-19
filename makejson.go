package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	Object := make(map[string]string)
	input := bufio.NewReader(os.Stdin)
	flg := true
	for flg {
		fmt.Println("Enter the Name:")
		temp, _, _ := input.ReadLine()
		nm := string(temp)
		fmt.Println("Enter the Address:")
		temp, _, _ = input.ReadLine()
		add := string(temp)
		Object[nm] = add
		jObj, _ := json.Marshal(Object)
		fmt.Println("Json Object:\n", jObj)
		fmt.Println("In String Format:\n", string(jObj))
		fmt.Print("Press 1 to repeat:")
		var x int
		fmt.Scan(&x)
		if x != 1 {
			flg = false
		}
	}
}
