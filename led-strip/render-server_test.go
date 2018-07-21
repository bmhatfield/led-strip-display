package ledstrip

import (
	"reflect"
	"testing"
)

func TestRenderServer_render(t *testing.T) {
	tests := []struct {
		name string
		r    *RenderServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.render()
		})
	}
}

func TestNewRenderServer(t *testing.T) {
	type args struct {
		strip Strip
	}
	tests := []struct {
		name string
		args args
		want *RenderServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRenderServer(tt.args.strip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRenderServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
