package _interface

import (
	"reflect"
	"testing"
)

func Test_testInterface1(t *testing.T) {
	type args struct {
		testabc []abc
		testdef []def
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "base test",
			args: args{
				testabc: []abc{"dfasdf", "aasdfadf", "asdfasdf"},
				testdef: []def{"1324", "1324", "123412"},
			},
			want: []interface{{"dfasdf", "aasdfadf", "asdfasdf"}, {"1324", "1324", "123412"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testInterface(tt.args.testabc, tt.args.testdef); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}