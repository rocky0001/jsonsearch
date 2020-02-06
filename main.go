package main

import (
	//"fmt"
	//"github.com/thedevsaddam/gojsonq"
	prompt "github.com/c-bata/go-prompt"
)

  
type CurrentStatus struct {
	PromptPrefix string
	Title string
}
var currentstatus CurrentStatus


func main() {
	currentstatus.PromptPrefix = "search"
    currentstatus.Title = "JSON Search"
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(currentstatus.PromptPrefix+">> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle(currentstatus.Title),
	)
	p.Run()
	
	
}
