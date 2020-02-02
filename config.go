package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"github.com/thedevsaddam/gojsonq"

)



type AppConfig struct {
	UsersJSONFile         string `yaml:"usersJsonFile"`
	TicketsJSONFile       string `yaml:"ticketsJsonFile"`
	OrganizationsJSONFile string `yaml:"organizationsJsonFile"`
	Outputfileds          struct {
		Users         []string `yaml:"users"`
		Tickets       []string `yaml:"tickets"`
		Organizations []string `yaml:"organizations"`
	} `yaml:"outputfileds"`
}

type SearchConfig struct {
	
	JQ        map[string]*gojsonq.JSONQ
	Fields    map[string][]string
	Outputs   map[string][]string
} 
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n",args...)
	os.Exit(1)
}

func getSearchFields(s string) []string {
	res := (searchconfig.JQ[s].First().(map[string]interface{}))
	fields := make([]string, len(res))
	i := 0
	for k,_ := range res {
		fields[i] =k
		i++
	}	
	return fields
}
var appconfig AppConfig 
var searchconfig SearchConfig
func init() {
	configFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
		exitErrorf("Read Config Yaml File Error: ", err)
	}
	err = yaml.Unmarshal(configFile,&appconfig)
	if err != nil {
		exitErrorf("Parsing Config File Error: ", err)
	}
	searchconfig.JQ = make(map[string]*gojsonq.JSONQ)
	searchconfig.Outputs = make(map[string][]string)
	searchconfig.Fields = make(map[string][]string)
	searchconfig.JQ["Users"] = gojsonq.New().File(appconfig.UsersJSONFile)
	searchconfig.JQ["Tickets"] = gojsonq.New().File(appconfig.TicketsJSONFile)
	searchconfig.JQ["Organizations"] = gojsonq.New().File(appconfig.OrganizationsJSONFile)
	searchconfig.Outputs["Users"] = appconfig.Outputfileds.Users
	searchconfig.Outputs["Tickets"] = appconfig.Outputfileds.Tickets
	searchconfig.Outputs["Organizations"] = appconfig.Outputfileds.Organizations
	searchconfig.Fields["Users"] = getSearchFields("Users")
	searchconfig.Fields["Tickets"] = getSearchFields("Tickets")
	searchconfig.Fields["Organizations"] = getSearchFields("Organizations")
	fmt.Println("userjson:",searchconfig.Fields["Tickets"])
}