package validators

import "testing"

func TestFormatErrors_usesJSONFieldNames(t *testing.T) {
	type payload struct {
		Email string `json:"email" validate:"required,email"`
	}

	err := Get().Struct(&payload{Email: "bad"})
	got := FormatErrors(err)

	if got["email"] != "invalid email" {
		t.Fatalf("got %q, want %q", got["email"], "invalid email")
	}
}

func TestFormatErrors_nonValidationError(t *testing.T) {
	got := FormatErrors(nil)
	if got["error"] != "invalid input" {
		t.Fatalf("got %v", got)
	}
}
