package main

import (
	"strings"
	"testing"
)

func TestParseCreateTaskRequest_ValidInput(t *testing.T) {
	body := strings.NewReader(`{"user_id":1,"title":"walk the dog","date":"2026-07-17T00:00:00Z"}`)

	request, err := parseCreateTaskRequest(body)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if request.Title != "walk the dog" {
		t.Errorf("expected title 'walk the dog', got %q", request.Title)
	}
	if request.UserID != 1 {
		t.Errorf("expected user_id 1, got %d", request.UserID)
	}
}

func TestParseCreateTaskRequest_InvalidJSON(t *testing.T) {
	body := strings.NewReader(`{invalid json}`)

	_, err := parseCreateTaskRequest(body)

	if err == nil {
		t.Errorf("expected an error for invalid JSON, got nil")
	}
}