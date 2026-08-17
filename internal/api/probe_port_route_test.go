package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"task047-portledger/internal/registry"
)

func TestProbePortRouteRequiresHostsSuffix(t *testing.T) {
	rr := httptest.NewRecorder()
	New(registry.New()).ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/ports/80", nil))
	if rr.Code != http.StatusNotFound {
		t.Fatalf("status=%d want 404; body=%s", rr.Code, rr.Body.String())
	}
}
