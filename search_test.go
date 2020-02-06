package main

import (
	"testing"

	"github.com/cheynewallace/tabby"
)

func Test_searchRelatedRecords(t *testing.T) {
	type args struct {
		file  string
		json  string
		key   string
		value interface{}
		t     *tabby.Tabby
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchRelatedRecords(tt.args.file, tt.args.json, tt.args.key, tt.args.value, tt.args.t)
		})
	}
}

func TestCurrentJsonSearch_Search(t *testing.T) {
	type fields struct {
		CurrentJson string
		Key         string
		Operator    string
		Value       interface{}
	}
	type args struct {
		new bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := &CurrentJsonSearch{
				CurrentJson: tt.fields.CurrentJson,
				Key:         tt.fields.Key,
				Operator:    tt.fields.Operator,
				Value:       tt.fields.Value,
			}
			js.Search(tt.args.new)
		})
	}
}

func Test_isInputValid(t *testing.T) {
	type args struct {
		input string
		ops   []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInputValid(tt.args.input, tt.args.ops); got != tt.want {
				t.Errorf("isInputValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
