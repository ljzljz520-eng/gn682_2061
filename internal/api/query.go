package api

import (
	"inspectionbase/internal/domain"
	"inspectionbase/internal/reporting"
	"net/http"
	"strconv"
)

func parseFilter(r *http.Request) domain.QueryFilter {
	q := r.URL.Query()
	n, _ := strconv.Atoi(q.Get("limit"))
	return domain.QueryFilter{DeviceID: q.Get("device"), Status: q.Get("status"), Inspector: q.Get("inspector"), Limit: n}
}
func writeItems(w http.ResponseWriter, v interface{}) { writeJSON(w, http.StatusOK, v) }
func acceptsJSON(r *http.Request) bool {
	return r.Header.Get("Accept") == "" || r.Header.Get("Accept") == "application/json"
}
func methodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
func pathID(path string) string {
	for len(path) > 0 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
func statusCode(err error) int {
	if err == nil {
		return 200
	}
	return 500
}
func safeLimit(v int) int {
	if v < 1 {
		return 20
	}
	if v > 100 {
		return 100
	}
	return v
}
func jsonContent(w http.ResponseWriter)                             { w.Header().Set("Content-Type", "application/json") }
func noCache(w http.ResponseWriter)                                 { w.Header().Set("Cache-Control", "no-store") }
func summaryResponse(v []domain.InspectionRecord) reporting.Summary { return reporting.Build(v) }
func queryDevice(r *http.Request) string                            { return r.URL.Query().Get("device") }
func queryStatus(r *http.Request) string                            { return r.URL.Query().Get("status") }
func queryInspector(r *http.Request) string                         { return r.URL.Query().Get("inspector") }
func isMutation(r *http.Request) bool {
	return r.Method == "POST" || r.Method == "PATCH" || r.Method == "PUT" || r.Method == "DELETE"
}
func isRead(r *http.Request) bool                { return r.Method == "GET" || r.Method == "HEAD" }
func responseMessage(v string) map[string]string { return map[string]string{"message": v} }
func errorMessage(v string) map[string]string    { return map[string]string{"error": v} }
func statusText(code int) string {
	if code >= 500 {
		return "server_error"
	}
	if code >= 400 {
		return "client_error"
	}
	return "ok"
}
