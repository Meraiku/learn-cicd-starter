package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("test getting API Key with invalid API Key", func(t *testing.T) {
		auth := map[string][]string{"Authorization": []string{"000"}}

		_, err := GetAPIKey(auth)
		if err == nil {
			t.Fatal("expected error")
		}

		want := "malformed authorization header"

		if err.Error() != want {
			t.Fatalf("expected: %v, got: %v", want, err)
		}

	})

	t.Run("test getting API Key without header", func(t *testing.T) {
		auth := map[string][]string{"Auth": []string{}}

		_, err := GetAPIKey(auth)
		if err == nil {
			t.Fatal("expected error")
		}

		want := "no authorization header included"

		if err.Error() != want {
			t.Fatalf("expected: %v, got: %v", want, err)
		}

	})

	t.Run("test getting API Key with ApiKey", func(t *testing.T) {
		auth := map[string][]string{"Authorization": []string{"ApiKey ayaya"}}

		got, _ := GetAPIKey(auth)
		want := "ayaya"

		if got != want {
			t.Fatalf("expected: %v, got: %v", want, got)
		}

	})
}
