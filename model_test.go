package jsonstatus_test

import (
	"github.com/bonsai-oss/jsonstatus"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestStatus_Encode(t *testing.T) {
	for _, tt := range []struct {
		name string
		s    *jsonstatus.Status
	}{
		{
			name: "success",
			s: &jsonstatus.Status{
				Code:    http.StatusOK,
				Message: "OK",
			},
		},
		{
			name: "error",
			s: &jsonstatus.Status{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				jsonstatus.Status{Code: tt.s.Code, Message: tt.s.Message}.Encode(w)

			}))
			defer svr.Close()

			req, err := http.NewRequest("GET", svr.URL, nil)
			if err != nil {
				t.Fatal(err)
			}
			rsp, _ := http.DefaultClient.Do(req)
			if rsp.StatusCode != tt.s.Code {
				t.Errorf("got %v, want %v", rsp.StatusCode, tt.s.Code)
			}

			if rsp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("got %v, want %v", rsp.Header.Get("Content-Type"), "application/json")
			}

			if responseStatus, decodeError := jsonstatus.Decode(rsp.Body); responseStatus != nil {
				if !reflect.DeepEqual(responseStatus, tt.s) {
					t.Errorf("got %v, want %v", responseStatus, tt.s)
				}
			} else {
				t.Error(decodeError)
			}
		})
	}
}
