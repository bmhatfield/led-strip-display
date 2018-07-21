package frame

import (
	"reflect"
	"testing"
)

func TestHexRGB_RenderFrame(t *testing.T) {
	tests := []struct {
		name    string
		f       *HexRGB
		want    HexGRB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.RenderFrame()
			if (err != nil) != tt.wantErr {
				t.Errorf("HexRGB.RenderFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexRGB.RenderFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexRGB_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		f       HexRGB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("HexRGB.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHexRGB_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		f       HexRGB
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("HexRGB.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexRGB.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
