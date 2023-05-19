package linux

import "testing"

func TestUname(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"test", "N/A", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uname()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uname() got = %v, want %v", got, tt.want)
			}
		})
	}
}
