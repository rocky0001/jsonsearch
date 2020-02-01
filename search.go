package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

type JsonSearch struct {
	JQ *gojsonq.JSONQ
	CurrentJson string
	Key string
	Operator  string
	Value string
}

func (js *JsonSearch) Search() {
	res := js.JQ.Where(js.Key,js.Operator,js.Value).Get()
	fmt.Println(res)
}