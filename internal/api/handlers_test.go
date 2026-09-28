package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/oumaAt/fizzbuzz-server/internal/fizzbuzz"
)

func doRequest(method, target string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(method, target, nil)
	NewRouter().ServeHTTP(rec, req)
	return rec
}

func TestFizzBuzzHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantBody   []string
		wantError  string
	}{
		{
			name:       "valid request",
			query:      "int1=2&int2=3&limit=6&str1=tic&str2=tac",
			wantStatus: http.StatusOK,
			wantBody:   []string{"1", "tic", "tac", "tic", "5", "tictac"},
		},
		{
			name:       "missing int1",
			query:      "int2=5&limit=15&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  "int1 is required",
		},
		{
			name:       "int1 is not an integer",
			query:      "int1=abc&int2=5&limit=15&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  "int1 must be an integer",
		},
		{
			name:       "int1 is zero",
			query:      "int1=0&int2=5&limit=15&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  fizzbuzz.ErrInvalidDivisor.Error(),
		},
		{
			name:       "limit is zero",
			query:      "int1=3&int2=5&limit=0&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  fizzbuzz.ErrInvalidLimit.Error(),
		},
		{
			name:       "limit above maximum",
			query:      "int1=3&int2=5&limit=10001&str1=fizz&str2=buzz",
			wantStatus: http.StatusBadRequest,
			wantError:  "limit must not exceed 10000",
		},
		{
			name:       "missing str2",
			query:      "int1=3&int2=5&limit=15&str1=fizz",
			wantStatus: http.StatusBadRequest,
			wantError:  "str1 and str2 are required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(http.MethodGet, "/fizzbuzz?"+tt.query)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("Content-Type = %q, want application/json", ct)
			}

			if tt.wantError != "" {
				var body map[string]string
				if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
					t.Fatalf("decoding error body: %v", err)
				}
				if body["error"] != tt.wantError {
					t.Errorf("error = %q, want %q", body["error"], tt.wantError)
				}
				return
			}

			var got []string
			if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
				t.Fatalf("decoding body: %v", err)
			}
			if !reflect.DeepEqual(got, tt.wantBody) {
				t.Errorf("body = %v, want %v", got, tt.wantBody)
			}
		})
	}
}

func TestFizzBuzzHandler_MethodNotAllowed(t *testing.T) {
	rec := doRequest(http.MethodPost, "/fizzbuzz")

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}