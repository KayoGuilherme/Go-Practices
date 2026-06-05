package validators

import "testing"

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		cpf   string
		valid bool
	}{
		{"52998224725", true},
		{"11144477735", true},
		{"00000000000", false},
		{"12345678901", false},
		{"123", false},
		{"5299822472a", false},
	}

	for _, tt := range tests {
		if got := isValidCPF(tt.cpf); got != tt.valid {
			t.Errorf("isValidCPF(%q) = %v, want %v", tt.cpf, got, tt.valid)
		}
	}
}
