package main

import (
	"fmt"
	"strconv"
    "github.com/thedevsaddam/gojsonq"
	
)

type JsonSearch struct {
	FromJson string
	JQ *gojsonq.JSONQ
	Key string
	Operator  string
	IsNew bool
	Value  interface{}
	Outputs []string
	RelatedSearch []SubSearch
}
type SubSearch struct {
	FromJson string
	JQ *gojsonq.JSONQ
	Key string
	RelatedKey string
	Note string
	ShortOutputs []string
}
func createSearch(json string,key string,op string,value interface{},isNew bool) JsonSearch {
	var js JsonSearch
	js.FromJson = json
	js.JQ = searchconfig.JQ[json]
	js.Key = key
	js.Value = value
	js.Operator = op
	js.Outputs = searchconfig.Outputs[json]
	js.IsNew = isNew
	switch js.FromJson {
	case "users":
		js.RelatedSearch =  append(js.RelatedSearch,SubSearch{
			"organizations",
			searchconfig.JQ["organizations"],
			"_id",
			"organization_id",
			"organization",
			searchconfig.ShortOutputs["organizations"],
		   },
		   SubSearch{
			"tickets",
			searchconfig.JQ["tickets"],
			"submitter_id",
			"_id",
			"submitted_tickets",
            searchconfig.ShortOutputs["tickets"],
		   },SubSearch{
			"tickets",
			searchconfig.JQ["tickets"],
			"assignee_id",
			"_id",
			"assigned_tickets",
            searchconfig.ShortOutputs["tickets"],
		   },
		)
	case "tickets":
		js.RelatedSearch =  append(js.RelatedSearch,
			SubSearch{
			"organizations",
			searchconfig.JQ["organizations"],
			"_id",
			"organization_id",
			"organization",
            searchconfig.ShortOutputs["organizations"],
		   },
		   SubSearch{
			"users",
			searchconfig.JQ["users"],
			"_id",
			"submitter_id",
			"submitter_users",
            searchconfig.ShortOutputs["users"],
		   },
		   SubSearch{
			"users",
			searchconfig.JQ["users"],
			"_id",
			"assignee_id",
			"assignee_users",
            searchconfig.ShortOutputs["users"],
		   },
		)
    case "organizations":
        js.RelatedSearch =  append(js.RelatedSearch,
			SubSearch{
			"users",
			searchconfig.JQ["users"],
			"organization_id",
			"_id",
			"users",
            searchconfig.ShortOutputs["users"],
		   },
		   SubSearch{
			"tickets",
            searchconfig.JQ["tickets"],
			"organization_id",
			"_id",
			"tickets",
            searchconfig.ShortOutputs["tickets"],
		   },
		)
		
	}
	return js
}

func  getRelatedRecords(js JsonSearch,value map[string]interface{}) []Results {
	if len(js.RelatedSearch) >0 {
		var results []Results
		var note string
		fmt.Println("Searching related records:")
		for o :=0;o<len(js.RelatedSearch);o++ {
			s := js.RelatedSearch[o].JQ.Reset().Where(js.RelatedSearch[o].Key,"=",value[js.RelatedSearch[o].RelatedKey]).Get()
			fmt.Println("Searching related records from:",js.RelatedSearch[o].FromJson,"where",js.RelatedSearch[o].Key,"=",value[js.RelatedSearch[o].RelatedKey])
	        if subresults,ok := (s.([]interface{})); ok {
	        	if len(subresults) > 0 {
	        		for p,res := range subresults {
	        			
						note = js.RelatedSearch[o].Note + "_"+strconv.Itoa(p+1)
						outputs := js.RelatedSearch[o].ShortOutputs
						//add dividing  line
					    results = append(results,Results{
					    	".....",
					    	"..........",
					    	"...........",
					    })
						for q :=0; q < len(outputs); q++ {
							//t.AddLine(json,output[h],(res[output[h]]))
							results = append(results,Results{
								note,
								outputs[q],
								res.(map[string]interface{})[outputs[q]],
							})
						}
	        			
					}
					
	        	}
	        }

		}
		return results
	}

	return nil
	
}

func (js *JsonSearch) Search() interface{} {
		var res interface{}
		if js.IsNew {
			res = js.JQ.Reset().Where(js.Key,js.Operator,js.Value).Get()
		} else {
	        res = js.JQ.Where(js.Key,js.Operator,js.Value).Get()
		}
		return res
}
type Results struct {
	Note string
	Key string
	Value interface{}
}
func getSearchResults(js JsonSearch) []Results {

	var results  []Results
	res := js.Search()
	if arr,ok := (res.([]interface{})); ok {
		if len(arr) > 0 {
			//t := tabby.New()
			for i :=0; i < len(arr); i++ {
				var note string
				note = js.FromJson+"_"+strconv.Itoa(i+1)
				outputs := searchconfig.Outputs[js.FromJson]
				for q :=0; q < len(outputs); q++ {
					results = append(results,Results{
						note,
						outputs[q],
						arr[i].(map[string]interface{})[outputs[q]],
					})
				}
				results = append(results,getRelatedRecords(js,arr[i].(map[string]interface{}))...)
				results = append(results,Results{
					"-----",
					"----------",
					"---------------",
				})
			
			}
			return results
		} else {
			fmt.Println("Warning: No record was found.")
			return nil
	       }
    }
	return nil
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