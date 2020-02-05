package main

import "github.com/cheynewallace/tabby"

func addTableLine(t *tabby.Tabby,json string, output []string, arr map[string]interface{}) {
	for h :=0; h < len(output); h++ {
		t.AddLine(json,output[h],(arr[output[h]]))
    }
}