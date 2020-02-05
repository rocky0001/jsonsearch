package main

import (
	"fmt"
	"github.com/cheynewallace/tabby"
	//"github.com/thedevsaddam/gojsonq"
)

type CurrentJsonSearch struct {
	CurrentJson string
	Key string
	Operator  string
	Value  interface{}
}

func (js *CurrentJsonSearch) Search() {
	
	res := searchconfig.JQ[js.CurrentJson].Reset().Where(js.Key,js.Operator,js.Value).Get()
	// res = searchconfig.JQ[js.CurrentJson].Reset().From("tags").Get()
	// fmt.Println(res)
	//var arr map[string]interface{}
	
	if arr,ok := (res.([]interface{})); ok {
		if len(arr) > 0 {
			t := tabby.New()
			for i :=0; i < len(arr); i++ {
				//fmt.Println("name.",i," value: ", (arr[i].(map[string]interface{}))["name"])
				t.AddHeader("---JSON---", "---Key---","---Value---")
				// for h :=0; h < len(searchconfig.Outputs[js.CurrentJson]); h++ {
				// 	t.AddLine(js.CurrentJson,searchconfig.Outputs[js.CurrentJson][h],  (arr[i].(map[string]interface{}))[searchconfig.Outputs[js.CurrentJson][h]])
				// }
				addTableLine(t,js.CurrentJson,searchconfig.Outputs[js.CurrentJson],arr[i].(map[string]interface{}))
				//search related entities from other json.
				switch js.CurrentJson {
				case "users":
					//search organization_id
					orgsearch :=searchconfig.JQ["organizations"].Reset().Where("_id","=",(arr[i].(map[string]interface{}))["organization_id"]).Get()
					if orgs,orgok := (orgsearch.([]interface{})); orgok {
						if len(orgs) > 0 {
                            for _,org := range orgs {
								addTableLine(t,"organizations",searchconfig.ShortOutputs["organizations"],org.(map[string]interface{}))
								
							}
						}
				// case "tickets":

				// case "organizations":
					}	
				}
			}
            t.Print()
		}
	} else {
        fmt.Println("Warning: No record was found.")
	}

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