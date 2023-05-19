package linux

import "testing"

func TestUser(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"na", "N/A", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := User()
			if (err != nil) != tt.wantErr {
				t.Errorf("User() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User() got = %v, want %v", got, tt.want)
			}
		})
	}
}
