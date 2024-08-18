package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("test getting API Key", func(t *testing.T) {
		auth := map[string][]string{"Authorization": []string{"000", "oiwjd"}}

		_, err := GetAPIKey(auth)
		if err == nil {
			t.Fatal("expected error")
		}

		want := "malformed authorization header"

		if err.Error() != want {
			t.Fatalf("expected: %v, got: %v", want, err)
		}

	})
}
