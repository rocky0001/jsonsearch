package main

import (
	"reflect"
	"testing"

	prompt "github.com/c-bata/go-prompt"
)

func Test_getRegexGroups(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRegexGroups(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRegexGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executorList(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executorList(tt.args.args)
		})
	}
}

func Test_getInputValue(t *testing.T) {
	type args struct {
		args  []string
		input string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInputValue(tt.args.args, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInputValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executorQuery(t *testing.T) {
	type args struct {
		args  []string
		input string
		new   bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executorQuery(tt.args.args, tt.args.input, tt.args.new)
		})
	}
}

func Test_executorSelect(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executorSelect(tt.args.args)
		})
	}
}

func Test_executor(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor(tt.args.in)
		})
	}
}

func Test_rootListCompleter(t *testing.T) {
	tests := []struct {
		name string
		want []prompt.Suggest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rootListCompleter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootListCompleter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operatorListCompleter(t *testing.T) {
	tests := []struct {
		name string
		want []prompt.Suggest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operatorListCompleter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operatorListCompleter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subListCompleter(t *testing.T) {
	tests := []struct {
		name string
		want []prompt.Suggest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subListCompleter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subListCompleter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonListCompleter(t *testing.T) {
	tests := []struct {
		name string
		want []prompt.Suggest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonListCompleter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonListCompleter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldListCompleter(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []prompt.Suggest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fieldListCompleter(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldListCompleter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatStringList(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatStringList(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatStringList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_completer(t *testing.T) {
	type args struct {
		in prompt.Document
	}
	tests := []struct {
		name string
		args args
		want []prompt.Suggest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := completer(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("completer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_changeLivePrefix(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := changeLivePrefix()
			if got != tt.want {
				t.Errorf("changeLivePrefix() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("changeLivePrefix() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
