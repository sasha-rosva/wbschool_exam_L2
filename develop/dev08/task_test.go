package main

import "testing"

func Test_execInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1",args: args{input: "cd"},wantErr: true},
		{name: "2",args: args{input: "cd /home"},wantErr: false},
		{name: "3",args: args{input: "ls"},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := execInput(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("execInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
