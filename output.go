package main

import "github.com/cheynewallace/tabby"

func showResults(res []Results) {
    if res != nil && len(res) >0 {
		t := tabby.New()
		t.AddHeader("--", "---Key---","---Value---")
		for i :=0; i < len(res); i++ {
			t.AddLine(res[i].Note,res[i].Key,res[i].Value)
		}
		t.Print()
	}
	
}