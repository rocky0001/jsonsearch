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
	//res := (searchconfig.JQ["users"].First().(map[string]interface{}))
	//res := (gojsonq.New().File("./tickets.json").First().(map[string]interface{}))
	//fmt.Printf("%#v\n", res["_id"])
	currentstatus.PromptPrefix = "search"
    currentstatus.Title = "Zendesk Search"
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(currentstatus.PromptPrefix+">> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle(currentstatus.Title),
	)
	p.Run()
	//jqUsers := gojsonq.New().File(appconfig.UsersJSONFile)
	// jqTickets := gojsonq.New().File(appconfig.TicketsJSONFile)
	// jqOrganizations := gojsonq.New().File(appconfig.OrganizationsJSONFile)
	//jq := gojsonq.New().File("./tickets.json")
	//fmt.Println(jq)
	//res, err:= gojsonq.New().File("./tickets.json").Select("tags").Where("_id", "=", "436bf9b0-1147-4c0a-8439-6f79833bff5b").GetR()
	//res := gojsonq.New().File("./tickets.json").Where("_id", "=", "436bf9b0-1147-4c0a-8439-6f79833bff5b").Get()//.([]interface{})
	//res := gojsonq.New().File("./tickets.json").First()
	// parameter := [3]string{"_id", "=", "436bf9b0-1147-4c0a-8439-6f79833bff5b"}
	
	// res := gojsonq.New().File("./tickets.json").Where(parameter[0],parameter[1],parameter[2]).Get()
	// res1 := gojsonq.New().File("./tickets.json").Get()
	var s CurrentJsonSearch
	s.CurrentJson = "Users"
	s.Key = "name"
	s.Operator = "="
	s.Value = "Loraine Pittman"
	s.Search()
	// res1 := jqUsers.First()
	// res2 := jqTickets.First()
	// res3 := jqOrganizations.First()
	// // fmt.Println(res)
	// // fmt.Println(res1)
	// fmt.Println(res1)
	// fmt.Println(res2)
	// fmt.Println(res3)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//name, _ := res.String()
	//fmt.Printf("%#v\n", res.(map[string]interface{})["_id"])
	// fmt.Printf("%#v\n", (res.(map[string]interface{})))
// 	mymap := (res.(map[string]interface{}))
// 	//keys := make([]int, len(mymap))
    
//    // i := 0
//     for k,_ := range mymap {
// 		// 
		
// 		fmt.Println(k)
// 	}
// 	res4 := jqUsers.Last()
//     fmt.Println(res4)
	//fmt.Println(keys)
    	// res := jq.Where("_id", "=", "436bf9b0-1147-4c0a-8439-6f79833bff5b").OrWhere("_id", "=", nil).Get()
	// fmt.Println(res)
	// output: [map[price:1350 id:1 name:MacBook Pro 13 inch retina] map[id:2 name:MacBook Pro 15 inch retina price:1700] map[id:<nil> name:HP core i3 SSD price:850]]
}
