package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_checkFlags(t *testing.T) {
	n1:=0
	n2:=2
	type args struct {
		grep *grep
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "1",args: args{grep: &grep{flagsN: []*int{&n2,&n1,&n1}}},want: 1},
		{name: "2",args: args{grep: &grep{flagsN: []*int{&n2,&n2,&n1}}},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkFlags(tt.args.grep)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkFlags() got = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_newt(t *testing.T) {
	n1:=0
	n2:=1
	sss := new(grep)
	sss.flagsN = append([]*int{}, &n1, &n2, &n2)
	sss.flagsB = append([]*bool{}, c, i, v, f, n)
	fmt.Println(*sss.flagsB[0])
	type args struct {
		arr  []string
		text string
		grep *grep
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "3",args: args{arr: []string{"computer","mouse","LAPTOP,time!","data","RedHat","laptoptime?",
			"debian","laptop","time","gggg","kjkjk","fgh","Timegggg","fgdhftimf","uuu","jjj",""},text: "time",
		grep: sss},want: []int{2,5,8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newt(tt.args.arr, tt.args.text, tt.args.grep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "4",args: args{file: "test.txt"},want: []string{"computer","mouse","LAPTOP,time!","data","RedHat","laptoptime?",
			"debian","laptop","time","gggg","kjkjk","fgh","Timegggg","fgdhftimf","uuu","jjj",""}},
		{name: "5",args: args{file: "file.txt"},want: nil,wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
