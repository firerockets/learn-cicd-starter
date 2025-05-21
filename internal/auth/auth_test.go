package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "Bearer 123456789")

	apiKey, err := GetAPIKey(header)

	if err != nil {
		t.Fatal("Expected success. Valid header should not throw an error")
	}

	if apiKey != "123456789" {
		t.Fatalf("Expected %v, got %v", "123456789", apiKey)
	}

	header = http.Header{}

	_, err = GetAPIKey(header)

	if err == nil {
		t.Fatal("Expected error. Invalid header should throw an error")
	}

	header.Add("Authorization", "Lorem")

	_, err = GetAPIKey(header)

	if err == nil {
		t.Fatal("Expected error. Invalid header should throw an error")
	}

	header.Add("Authorization", "Lorem Ipsum")

	_, err = GetAPIKey(header)

	if err == nil {
		t.Fatal("Expected error. Invalid header should throw an error")
	}
}
