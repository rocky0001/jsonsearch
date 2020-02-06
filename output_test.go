package main

import (
	"testing"

	"github.com/cheynewallace/tabby"
)

func Test_addTableLine(t *testing.T) {
	type args struct {
		t      *tabby.Tabby
		json   string
		output []string
		arr    map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addTableLine(tt.args.t, tt.args.json, tt.args.output, tt.args.arr)
		})
	}
}
