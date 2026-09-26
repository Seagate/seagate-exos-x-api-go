package mcapi

import (
	"net/http"
	"testing"
)

// Regression: log lines after ExecuteWithFailover dereferenced httpRes.Status
// unconditionally; a transport-layer failure (nil httpRes) SIGSEGV'd the driver.
func TestHTTPStatusNilSafe(t *testing.T) {
	if got := httpStatus(nil); got != "<no response>" {
		t.Errorf("httpStatus(nil) = %q, want %q", got, "<no response>")
	}
	if got := httpStatus(&http.Response{Status: "200 OK"}); got != "200 OK" {
		t.Errorf("httpStatus(non-nil) = %q, want %q", got, "200 OK")
	}
}
