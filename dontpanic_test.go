package dontpanic

import (
	"reflect"
	"testing"
)

func TestSliceIndex(t *testing.T) {
	type args struct {
		s []string
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{s: nil, i: 0},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceIndex(tt.args.s, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMapIndex(t *testing.T) {
	type args struct {
		m map[string]int
		k string
		v int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{m: nil, k: "foo", v: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetMapIndex(tt.args.m, tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
