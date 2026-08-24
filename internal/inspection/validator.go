package inspection

import (
	"inspectionbase/internal/domain"
	"strings"
)

func ValidateQuery(id string) error {
	if strings.TrimSpace(id) == "" {
		return domain.ErrInvalid
	}
	return nil
}
func SanitizeNotes(v string) string { return strings.TrimSpace(v) }
func EnsureDevice(st interface {
	GetDevice(string) (domain.Device, error)
}, id string) error {
	_, e := st.GetDevice(id)
	return e
}
