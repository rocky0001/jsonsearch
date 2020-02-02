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
func executor(in string) {
	fmt.Println("Your input: " + in)
	args := formatStringList(in)
	switch args[0] {
	case "quit":
		os.Exit(0)
	case "select":
		currentstatus.PromptPrefix = args[1]
		LivePrefixState.LivePrefix = currentstatus.PromptPrefix+">> " 
		LivePrefixState.IsEnable = true
	case "list":
		//todo
		fmt.Println("list all searchable records")
	case "where":
		//todo
		fmt.Println("Search records")
	case "return":
		//todo
		fmt.Println("return ")
		currentstatus.PromptPrefix = "search"
		LivePrefixState.LivePrefix = currentstatus.PromptPrefix+">> " 
		LivePrefixState.IsEnable = true
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
	        	
	}
	
	return []prompt.Suggest{}

	

}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}
