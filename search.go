package main

import (
	"fmt"
	//"github.com/thedevsaddam/gojsonq"
)

type CurrentJsonSearch struct {
	CurrentJson string
	Key string
	Operator  string
	Value  interface{}
}

func (js *CurrentJsonSearch) Search() {
	res := searchconfig.JQ[js.CurrentJson].Where(js.Key,js.Operator,js.Value).Get()
	fmt.Println(res)
}

func isInputValid(input string, ops []string) bool {
	for _, o := range ops {
		if o == input {
			return true
		}
	}
    fmt.Println("Error: ",input,"is invalid. The valid options: ",ops)
	return false
}