package main

import (
	"reflect"
	"testing"
)

func Test_open(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "1",args: args{s: "a\ns\nd"},want: []string{"a", "s", "d"}},
		{name: "2",args: args{s: "drwx------ 5 user user 12288 янв 15 14:59 Downloads\ndrwxr-xr-x 6 user user 4096 дек 6 14:29 Android\ndrwxr-xr-x 7 user user 4096 июн 10"+
			" 2015 Sources\ndrwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox\ndrwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks\ndrwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures"},
			want: []string{"drwx------ 5 user user 12288 янв 15 14:59 Downloads", "drwxr-xr-x 6 user user 4096 дек 6 14:29 Android", "drwxr-xr-x 7 user user 4096 июн 10 2015 Sources", "drw"+
				"xr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox", "drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks", "drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := open(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("open() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_openFile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "3",args: args{s: "test.txt"},want: []string{"computer", "mouse", "LAPTOP", "data", "RedHat", "laptop", "debian", "laptop"},wantErr: false},
		{name: "4",args: args{s: "file.txt"},want: nil,wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openFile(tt.args.s)
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

func Test_sortB(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "5",args: args{s:[]string{"a    ", "s ", "d    "}},want: []string{"a", "d", "s"}},
		{name: "6",args: args{s:[]string{"a  b  ", "s  d ", "u   f    "}},want: []string{"a b", "s d", "u f"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortB(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortC(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "7",args: args{s:[]string{"a", "b", "c"}},want: true},
		{name: "8",args: args{s:[]string{"a", "c", "b"}},want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortC(tt.args.s); got != tt.want {
				t.Errorf("sortC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortH(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "9",args: args{s:[]string{"1 12a", "2 9a", "3 7b"},n:2},want: []string{"3 7b", "2 9a", "1 12a"}},
		{name: "10",args: args{s:[]string{"1 a12", "2 9a", "3 b7"},n:2},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortH(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortH() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortK(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "11",args: args{s:[]string{"1 12a kit", "2 9a tok", "3 7b tik"},n:3},want: []string{"1 12a kit", "3 7b tik", "2 9a tok"}},
		{name: "12",args: args{s:[]string{"1 12a kit", "2 9a tok", "3 7b tik"},n:4},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortK(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortK() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortM(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "12",args: args{s:[]string{"1 12a ЯНварЬ", "2 9a МаЙ", "3 7b Март"},n:3},want: []string{"1 12a ЯНварЬ", "3 7b Март", "2 9a МаЙ"}},
		{name: "13",args: args{s:[]string{"1 12a ЯНварЬ", "2 9a МаЙ", "3 7b Март"},n:2},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortM(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortM() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortM() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortN(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "13",args: args{s:[]string{"3 12a ЯНварЬ", "2 9a МаЙ", "1 7b Март"},n:1},want: []string{"1 7b Март", "2 9a МаЙ", "3 12a ЯНварЬ"}},
		{name: "14",args: args{s:[]string{"1 12a ЯНварЬ", "2 9a МаЙ", "3 7b Март"},n:4},wantErr: true},
		{name: "15",args: args{s:[]string{"1 12a ЯНварЬ", "2 9a МаЙ", "3 7b Март"},n:3},wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortN(tt.args.s, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortR(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "16",args: args{s:[]string{"a 13 !","k 7 &","b 44 @"}},want: []string{"k 7 &","b 44 @","a 13 !"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortR(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortT(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "17",args: args{s:[]string{"a 13 !","k 7 &","b 44 @"}},want: []string{"a 13 !","b 44 @","k 7 &"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortT(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortU(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "18",args: args{s:[]string{"a 13 !","k 7 &","b 44 @","a 13 !"}},want: []string{"a 13 !","b 44 @","k 7 &"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortU(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortU() = %v, want %v", got, tt.want)
			}
		})
	}
}
