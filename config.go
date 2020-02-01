package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"

)

type AppConfig struct {
	
  UsersJSONFile         string `yaml:"usersJsonFile,omitempty"`
  TicketsJSONFile       string `yaml:"ticketsJsonFile,omitempty"`
  OrganizationsJSONFile string `yaml:"organizationsJsonFile,omitempty"`  
} 


func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n",args...)
	os.Exit(1)
}
func init() {
	configFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
		exitErrorf("Read Config Yaml File Error: ", err)
	}
	err = yaml.Unmarshal(configFile,&appconfig)
	if err != nil {
		exitErrorf("Parsing Config File Error: ", err)
	}
	fmt.Println("userjson:",appconfig.UsersJSONFile)
}