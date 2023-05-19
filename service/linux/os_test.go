package linux

import "testing"

func TestOs(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"test", "linux", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Os()
			if (err != nil) != tt.wantErr {
				t.Errorf("Os() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Os() got = %v, want %v", got, tt.want)
			}
		})
	}
}
