package main

import (
	"fmt"
	"strings"
	"regexp"
    "os"
	prompt "github.com/c-bata/go-prompt"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}
var inputExpression = regexp.MustCompile(`(?P<command>select|list|where)\s{1}`)
var jsonFileNames []string = []string{"users","tickets", "organizations"}
var validOperators []string = []string{"=","!=","<",">","<=",">=","in","notIn","startsWith","endsWith","contains"}
func getRegexGroups(text string) map[string]string {
	if !inputExpression.Match([]byte(text)) {
		return nil
	}

	match := inputExpression.FindStringSubmatch(text)
	result := make(map[string]string)
	for i, name := range inputExpression.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}
func executorList(args []string) {
	if len(args) == 1 {
		fmt.Println("List all the searchable fields:")
		fmt.Println("users:",searchconfig.Fields["users"])
		fmt.Println("tickets:",searchconfig.Fields["tickets"])
		fmt.Println("organizations:",searchconfig.Fields["organizations"])
	} else if len(args) > 2 || !isInputValid(args[1],jsonFileNames){
		fmt.Println("Incorrect arguments. Syntax: list [users|tickets|organizaitons]")
	} else {
		fmt.Println("List all the searchable fields:")
		fmt.Println(args[1],":",searchconfig.Fields[args[1]])
	}

}
func executorWhere(args []string) {
	if len(args) < 4 || !isInputValid(args[1],searchconfig.Fields[currentstatus.PromptPrefix]) || !isInputValid(args[2],validOperators) {
		fmt.Println("Incorrect arguments. Syntax: where <field>  <operator> <value> ")
	} else {
		 var s CurrentJsonSearch
	     s.CurrentJson = currentstatus.PromptPrefix
	     s.Key = args[1]
	     s.Operator = args[2]
		 s.Value = args[3]
		 s.Search()
	}

}
func executor(in string) {
	fmt.Println("Your input: " + in)
	args := formatStringList(in)
	switch currentstatus.PromptPrefix {
	case "search":
	    switch args[0] {
	    case "quit":
	    	os.Exit(0)
	    case "select":
	    	currentstatus.PromptPrefix = args[1]
	    	LivePrefixState.LivePrefix = currentstatus.PromptPrefix+">> " 
	    	LivePrefixState.IsEnable = true
	    case "list":
			executorList(args)
		default:
			fmt.Println("choose right options")
		}
	case "users","tickets","organizations":
		switch args[0] {
	    case "where":
	    	executorWhere(args)
	    case "return":
	    	fmt.Println("return ")
	    	currentstatus.PromptPrefix = "search"
	    	LivePrefixState.LivePrefix = currentstatus.PromptPrefix+">> " 
			LivePrefixState.IsEnable = true
		default:
			fmt.Println("where _ = ?")
		}
	}
	// if in == "quit" {
	// 	LivePrefixState.IsEnable = false
	// 	LivePrefixState.LivePrefix = in
		
	// }
	// LivePrefixState.LivePrefix = currentstatus.PromptPrefix+">> " 
	// LivePrefixState.IsEnable = true
}
func rootListCompleter() []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "select", Description: "Select the JSON file to search"},
		{Text: "list", Description: "List all the searchable fields"},
		{Text: "quit", Description: "Exit the application"},
	}
	return s
}
func operatorListCompleter() []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "=", Description: "Select the JSON file to search"},
		{Text: "!=", Description: "List all the searchable fields"},
		{Text: ">", Description: "Exit the application"},
		{Text: "<", Description: "Exit the application"},
		{Text: ">", Description: "Exit the application"},
		{Text: ">=", Description: "Exit the application"},
		{Text: "<=", Description: "Exit the application"},
		{Text: "in", Description: "Exit the application"},
		{Text: "notIn", Description: "Exit the application"},
		{Text: "contains", Description: "Exit the application"},
		{Text: "startsWith", Description: "Exit the application"},
		{Text: "endsWith", Description: "Exit the application"},
		
	}
	return s
}
func subListCompleter() []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "where", Description: "Search records from current JSON"},
		{Text: "return", Description: "Return to root menu "},
	}
	return s
}
func jsonListCompleter() []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: ""},
		{Text: "tickets", Description: ""},
		{Text: "organizations", Description: ""},
	}
	return s
}
func fieldListCompleter(name string) []prompt.Suggest {
	l := searchconfig.Fields[name]
	s := make([]prompt.Suggest,len(l))
	for i:=0;i<len(l);i++{
		s[i] = prompt.Suggest{
			Text: l[i], Description: "",
		}
	}
	return s
}
func formatStringList(s string) []string {
	s = strings.TrimSpace(s)
	s = regexp.MustCompile(`[\s\p{Zs}]{2,}`).ReplaceAllString(s, " ")
	l := strings.Split(s, " ")
	fmt.Println(l)
	return l
}
func completer(in prompt.Document) []prompt.Suggest {
	p := in.TextBeforeCursor()
	var list  []prompt.Suggest
	args := formatStringList(p)
	fmt.Println("first:",args[0])
	fmt.Println("beforecursor:",in.GetWordBeforeCursor())
	fmt.Println(len(strings.Split(in.Text, " ")))
	switch currentstatus.PromptPrefix {
	case "search":
		if len(strings.Split(in.Text, " ")) < 2 {
			list = rootListCompleter() 
			return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
		} else if len(strings.Split(in.Text, " ")) < 3 {
			group := getRegexGroups(in.Text)
	        if group != nil {
	        	command := group["command"]
        
	        		if (command == "select" || command == "list")  { 
	        		   
	        			   list = jsonListCompleter() 
	        		   }
	        		   return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
	        }  	
		}
		case "users","tickets","organizations":
			if len(strings.Split(in.Text, " ")) < 2 {
				list = subListCompleter() 
				return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
			} else {
				
				group := getRegexGroups(in.Text)
				if group != nil {
					command := group["command"]
			            fmt.Println(command)
						if (command == "where" )  { 
							if len(strings.Split(in.Text, " ")) < 3 {
							  list = fieldListCompleter(currentstatus.PromptPrefix)
						    } else if len(strings.Split(in.Text, " ")) < 4 {
                                list = operatorListCompleter()
							}
						   return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
				        }
				}
			}      	
	}
	return []prompt.Suggest{}
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}
