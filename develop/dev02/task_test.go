package main

import "testing"

func Test_decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		err     string
	}{
		{name: "1",args: args{s: "a4bc2d5e"},want: "aaaabccddddde"},
		{name: "2",args: args{s: "abcd"},want: "abcd"},
		{name: "3",args: args{s: "45"},want: "",wantErr: true, err: "некорректная строка"},
		{name: "4",args: args{s: ""},want: ""},
		{name: "5",args: args{s: `qwe\4\5`},want: "qwe45"},
		{name: "6",args: args{s: `qwe\45`},want: "qwe44444"},
		{name: "7",args: args{s: `qwe\\5`},want: `qwe\\\\\`},
		{name: "8",args: args{s: "a04"},want: "",wantErr: true,err: "число начинается с нуля"},
		{name: "9",args: args{s: "a12"},want: "aaaaaaaaaaaa"},
		{name: "10",args: args{s: `a12\3\4\5\\6\5`},want: `aaaaaaaaaaaa345\\\\\\5`}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
