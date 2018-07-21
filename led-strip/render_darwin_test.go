package ledstrip

import (
	"reflect"
	"testing"

	"github.com/bmhatfield/led-strip-display/frame"
)

func TestDarwinStrip_Init(t *testing.T) {
	type args struct {
		gpio       int
		pixels     int
		brightness int
	}
	tests := []struct {
		name    string
		s       DarwinStrip
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Init(tt.args.gpio, tt.args.pixels, tt.args.brightness); (err != nil) != tt.wantErr {
				t.Errorf("DarwinStrip.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDarwinStrip_Render(t *testing.T) {
	type args struct {
		strip frame.HexGRB
	}
	tests := []struct {
		name    string
		s       DarwinStrip
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Render(tt.args.strip); (err != nil) != tt.wantErr {
				t.Errorf("DarwinStrip.Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetStrip(t *testing.T) {
	tests := []struct {
		name string
		want Strip
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStrip(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
