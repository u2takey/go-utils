package shellquote

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantWords []string
		wantErr   bool
	}{
		{
			"test1",
			args{"--arg1=1 arg2=2 --args2 3"},
			[]string{"--arg1=1", "arg2=2", "--args2", "3"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWords, err := Split(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("Split() = %v, want %v", gotWords, tt.wantWords)
			}
		})
	}
}
