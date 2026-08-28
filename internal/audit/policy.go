package audit

import "inspectionbase/internal/domain"

func Allowed(action string) bool {
	switch action {
	case "create", "status", "delete", "export":
		return true
	default:
		return false
	}
}
func Normalize(a domain.AuditEntry) domain.AuditEntry {
	if a.Actor == "" {
		a.Actor = "system"
	}
	return a
}
func Actions() []string { return []string{"create", "status", "delete", "export"} }
