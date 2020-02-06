package main

import (
	"fmt"
	"strconv"
	"github.com/cheynewallace/tabby"
	
)

type CurrentJsonSearch struct {
	CurrentJson string
	Key string
	Operator  string
	Value  interface{}
}

func  searchRelatedRecords(file string,json string,key string, value interface{}, t *tabby.Tabby) {

	s := searchconfig.JQ[json].Reset().Where(key,"=",value).Get()
	if results,ok := (s.([]interface{})); ok {
		if len(results) > 0 {
			for i,res := range results {
				var rec string
				rec = file + "_"+strconv.Itoa(i+1)
				addTableLine(t,rec,searchconfig.ShortOutputs[json],res.(map[string]interface{}))
				
			}
		}
	}
	
}

func (js *CurrentJsonSearch) Search(new bool) {
	var res interface{}
	if new {
		res = searchconfig.JQ[js.CurrentJson].Reset().Where(js.Key,js.Operator,js.Value).Get()
	} else {
        res = searchconfig.JQ[js.CurrentJson].Where(js.Key,js.Operator,js.Value).Get()
	}
	
	
	if arr,ok := (res.([]interface{})); ok {
		if len(arr) > 0 {
			t := tabby.New()
			for i :=0; i < len(arr); i++ {
				var rec string
				rec = js.CurrentJson+"_"+strconv.Itoa(i+1)
				t.AddHeader(rec, "---Key---","---Value---")
				addTableLine(t,rec,searchconfig.Outputs[js.CurrentJson],arr[i].(map[string]interface{}))
				//search related entities from other json.
				switch js.CurrentJson {
				case "users":
					//search organization_id
					searchRelatedRecords("organization","organizations","_id",(arr[i].(map[string]interface{}))["organization_id"],t)
					//search tickets with submitter_id
					searchRelatedRecords("submitted_tickets","tickets","submitter_id",(arr[i].(map[string]interface{}))["_id"],t)
					//search tickets with assignee_id
					searchRelatedRecords("assigned_tickets","tickets","assignee_id",(arr[i].(map[string]interface{}))["_id"],t)
				case "tickets":
                    //search organization_id
					searchRelatedRecords("organization","organizations","_id",(arr[i].(map[string]interface{}))["organization_id"],t)
					//search users with submitter_id
					searchRelatedRecords("submitter_user","users","_id",(arr[i].(map[string]interface{}))["submitter_id"],t)
					//search tickets with assignee_id
					searchRelatedRecords("assignee_user","users","_id",(arr[i].(map[string]interface{}))["assignee_id"],t)
				case "organizations":
					//search organization_id from users
					searchRelatedRecords("users","users","organization_id",(arr[i].(map[string]interface{}))["_id"],t)
					//search organization_id from tickets
					searchRelatedRecords("tickets","tickets","organization_id",(arr[i].(map[string]interface{}))["_id"],t)
				}	
			}
			t.Print()
		} else {
            fmt.Println("Warning: No record was found.")
	       }
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