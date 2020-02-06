package main

import (
	"fmt"
	//"github.com/thedevsaddam/gojsonq"
	prompt "github.com/c-bata/go-prompt"
)

  
type CurrentStatus struct {
	PromptPrefix string
	Title string
}
var currentstatus CurrentStatus
const help1 = `
 welcome to Zendesk Search
 please input your command to continue, Type 'quit' to exit the application 
        <command> <options> 
	  select users|tickets|organizations
	  list  [users|tickets|organizations]
	  
`
const help2 = `
 please type your command to continue, Type 'return' to return root menu
	commands:
		>> where <key> <operator> <value>  //start a new search from current json file.
		>> filter <key> <operator> <value> //filter from last search results.
		>> return  //retrun to the root menu
	operators:
		** "=", Description: "For equality matching"
		** "!=", Description: "For not equality matching,
		** ">", Description: "Check if value of given key in data is Greater than val,
		** "<", Description: "Check if the value of given key in data is Less than val,
		** ">=", Description: "Check if the value of given key in data is Greater than or Equal of val,
		** "<=", Description: "Check if the value of given key in data is Less than or Equal of val,
		** "contains", Description: "Exit the application,
		** "startsWith", Description: "Check if the value of given key in data starts with (has a prefix of) the given val.
		** "endsWith", Description: "Check if the value of given key in data ends with (has a suffix of) the given val.,
		***** "has", Description: "Check if the list of given key in data has val, only for tags,domain_names.
	values: 
		** ture/false(bool)/nil(empty string "") are keywords, if they are strings, double quotes are required,for example: "true" (string), true(bool)
		** double quotes also requied for multiple strings with spaces for example: "this is the vaule"
		** digital numbers will be int if no double qutoes
	
	  
`
func main() {
	currentstatus.PromptPrefix = "search"
	currentstatus.Title = "JSON Search"
	fmt.Println(help1)
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(currentstatus.PromptPrefix+">> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle(currentstatus.Title),
	)
	p.Run()
	
	
}
