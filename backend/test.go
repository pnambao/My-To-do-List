package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTaskHandler(t *testing.T) {

	body := strings.NewReader(`{invalid json}`)

	req := httptest.NewRequest(http.MethodPost, "/tasks", body)

	rr := httptest.NewRecorder()

	createTaskHandler(rr, req)

	if rr.Code != http.StatusBadRequest {

		t.Errorf("expected 400 but got %d", rr.Code)

	}

}