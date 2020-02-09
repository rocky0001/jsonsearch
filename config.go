package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"
	"os"
	"sort"
	"github.com/thedevsaddam/gojsonq"

)



type AppConfig struct {
	UsersJSONFile         string `yaml:"usersJsonFile"`
	TicketsJSONFile       string `yaml:"ticketsJsonFile"`
	OrganizationsJSONFile string `yaml:"organizationsJsonFile"`
	Outputfields          struct {
		Users         []string `yaml:"users"`
		Tickets       []string `yaml:"tickets"`
		Organizations []string `yaml:"organizations"`
	} `yaml:"outputfields"`

	Relatedoutputfields   struct {
		Users         []string `yaml:"users"`
		Tickets       []string `yaml:"tickets"`
		Organizations []string `yaml:"organizations"`
	} `yaml:"relatedoutputfields"`
}

type SearchConfig struct {
	
	JQ        map[string]*gojsonq.JSONQ
	Fields    map[string][]string
	Outputs   map[string][]string
	ShortOutputs   map[string][]string
} 
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n",args...)
	os.Exit(1)
}

func getSearchFields(s string) []string {
    search :=searchconfig.JQ[s].First()
	err := searchconfig.JQ[s].Error()
    if err != nil {
		exitErrorf("Get Searchable fields error:",err)
	}
	res := search.(map[string]interface{})
	fields := make([]string, len(res))
	i := 0
	for k,_ := range res {
		fields[i] =k
		i++
	}	
	sort.Strings(fields)
	return fields
}
var appconfig AppConfig 
var searchconfig SearchConfig
var fx gojsonq.QueryFunc

func init() {
	configFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
		exitErrorf("Read Config Yaml File Error: ", err)
	}
	err = yaml.Unmarshal(configFile,&appconfig)
	if err != nil {
		exitErrorf("Parsing Config File Error: ", err)
	}
	fx = func(x, y interface{}) (bool, error) {
		arr := reflect.ValueOf(x)
		if arr.Kind() != reflect.Slice {
			return false, fmt.Errorf("%v invalid data type,\"has\" only support query list data,such as: tags,domain_names. your input:",x)
		}
		for _,t := range x.([]interface{}) {
		   if t ==y {
			   return true,nil
		   }	
		 }
		 return false,nil
	   }
	searchconfig.JQ = make(map[string]*gojsonq.JSONQ)
	searchconfig.Outputs = make(map[string][]string)
	searchconfig.ShortOutputs = make(map[string][]string)
	searchconfig.Fields = make(map[string][]string)
	searchconfig.JQ["users"] = gojsonq.New().File(appconfig.UsersJSONFile)
	searchconfig.JQ["tickets"] = gojsonq.New().File(appconfig.TicketsJSONFile)
	searchconfig.JQ["organizations"] = gojsonq.New().File(appconfig.OrganizationsJSONFile)
	searchconfig.Outputs["users"] = appconfig.Outputfields.Users
	searchconfig.Outputs["tickets"] = appconfig.Outputfields.Tickets
	searchconfig.Outputs["organizations"] = appconfig.Outputfields.Organizations
	searchconfig.ShortOutputs["users"] = appconfig.Relatedoutputfields.Users
	searchconfig.ShortOutputs["tickets"] = appconfig.Relatedoutputfields.Tickets
	searchconfig.ShortOutputs["organizations"] = appconfig.Relatedoutputfields.Organizations
	searchconfig.Fields["users"] = getSearchFields("users")
	searchconfig.Fields["tickets"] = getSearchFields("tickets")
	searchconfig.Fields["organizations"] = getSearchFields("organizations")
	searchconfig.JQ["users"].Macro("has",fx)
	searchconfig.JQ["tickets"].Macro("has",fx)
	searchconfig.JQ["organizations"].Macro("has",fx)
}