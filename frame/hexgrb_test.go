package frame

import (
	"reflect"
	"testing"
)

func TestHexGRB_RenderFrame(t *testing.T) {
	tests := []struct {
		name    string
		f       *HexGRB
		want    HexGRB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.RenderFrame()
			if (err != nil) != tt.wantErr {
				t.Errorf("HexGRB.RenderFrame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexGRB.RenderFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexGRB_HexRGB(t *testing.T) {
	tests := []struct {
		name    string
		f       *HexGRB
		want    HexRGB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.HexRGB()
			if (err != nil) != tt.wantErr {
				t.Errorf("HexGRB.HexRGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexGRB.HexRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexGRB_StringRGB(t *testing.T) {
	tests := []struct {
		name    string
		f       *HexGRB
		want    StringRGB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.StringRGB()
			if (err != nil) != tt.wantErr {
				t.Errorf("HexGRB.StringRGB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexGRB.StringRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexGRB_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		f       HexGRB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("HexGRB.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHexGRB_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		f       HexGRB
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("HexGRB.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexGRB.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
