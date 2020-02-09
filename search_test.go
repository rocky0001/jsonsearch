package main

import (
	"reflect"
	"testing"
)

func Test_createSearch(t *testing.T) {
	type args struct {
		json  string
		key   string
		op    string
		value interface{}
		isNew bool
	}
	tests := []struct {
		name string
		args args
		want JsonSearch
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSearch(tt.args.json, tt.args.key, tt.args.op, tt.args.value, tt.args.isNew); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRelatedRecords(t *testing.T) {
	type args struct {
		js    JsonSearch
		value map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []Results
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRelatedRecords(tt.args.js, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRelatedRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonSearch_Search(t *testing.T) {
	type fields struct {
		FromJson      string
		Key           string
		Operator      string
		IsNew         bool
		Value         interface{}
		RelatedSearch []SubSearch
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := &JsonSearch{
				FromJson:      tt.fields.FromJson,
				Key:           tt.fields.Key,
				Operator:      tt.fields.Operator,
				IsNew:         tt.fields.IsNew,
				Value:         tt.fields.Value,
				RelatedSearch: tt.fields.RelatedSearch,
			}
			if got := js.Search(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonSearch.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSearchResults(t *testing.T) {
	type args struct {
		js JsonSearch
	}
	tests := []struct {
		name string
		args args
		want []Results
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSearchResults(tt.args.js); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSearchResults() = %v, want %v", got, tt.want)
			}
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
