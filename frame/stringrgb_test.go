package frame

import (
	"reflect"
	"testing"
)

func TestStringRGB_RenderFrame(t *testing.T) {
	tests := []struct {
		name    string
		f       *StringRGB
		want    HexGRB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.RenderFrame()
			if (err != nil) != tt.wantErr {
				t.Errorf("StringRGB.RenderFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringRGB.RenderFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringRGBFromDisk(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *StringRGB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringRGBFromDisk(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringRGBFromDisk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringRGBFromDisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
