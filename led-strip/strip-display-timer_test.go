package ledstrip

import (
	"reflect"
	"testing"
)

func TestDisplayTimer_Run(t *testing.T) {
	tests := []struct {
		name string
		dt   *DisplayTimer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dt.Run()
		})
	}
}

func TestDisplayTimer_updater(t *testing.T) {
	tests := []struct {
		name string
		dt   *DisplayTimer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dt.updater()
		})
	}
}

func TestDisplayTimer_switcher(t *testing.T) {
	tests := []struct {
		name string
		dt   *DisplayTimer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dt.switcher()
		})
	}
}

func TestDisplayTimer_shouldBeOn(t *testing.T) {
	tests := []struct {
		name string
		dt   *DisplayTimer
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.shouldBeOn(); got != tt.want {
				t.Errorf("DisplayTimer.shouldBeOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDisplayTimer(t *testing.T) {
	type args struct {
		server *RenderServer
	}
	tests := []struct {
		name string
		args args
		want *DisplayTimer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDisplayTimer(tt.args.server); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDisplayTimer() = %v, want %v", got, tt.want)
			}
		})
	}
}
