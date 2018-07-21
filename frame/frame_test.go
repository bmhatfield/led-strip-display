package frame

import (
	"reflect"
	"testing"
)

func TestRGBToHex(t *testing.T) {
	type args struct {
		colors [3]int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RGBToHex(tt.args.colors); got != tt.want {
				t.Errorf("RGBToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGBToString(t *testing.T) {
	type args struct {
		colors [3]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RGBToString(tt.args.colors); got != tt.want {
				t.Errorf("RGBToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGBToColors(t *testing.T) {
	type args struct {
		pixel uint32
	}
	tests := []struct {
		name string
		args args
		want [3]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RGBToColors(tt.args.pixel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RGBToColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
