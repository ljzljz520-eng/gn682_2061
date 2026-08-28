package domain

import (
	"strings"
	"unicode"
)

func CleanID(v string) string       { return strings.TrimSpace(v) }
func CleanName(v string) string     { return strings.TrimSpace(v) }
func CleanLocation(v string) string { return strings.TrimSpace(v) }
func CleanNotes(v string) string    { return strings.TrimSpace(v) }
func IsASCII(v string) bool {
	for _, r := range v {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}
func HasSpace(v string) bool     { return strings.ContainsAny(v, " \t\n") }
func SplitID(v string) []string  { return strings.Split(CleanID(v), "-") }
func Prefix(v, p string) bool    { return strings.HasPrefix(v, p) }
func Suffix(v, p string) bool    { return strings.HasSuffix(v, p) }
func JoinID(v []string) string   { return strings.Join(v, "-") }
func EqualFold(a, b string) bool { return strings.EqualFold(a, b) }
func NonEmpty(v ...string) bool {
	for _, x := range v {
		if strings.TrimSpace(x) == "" {
			return false
		}
	}
	return true
}
func LimitText(v string, n int) string {
	if n < 1 {
		return ""
	}
	r := []rune(v)
	if len(r) > n {
		return string(r[:n])
	}
	return v
}
func CountWords(v string) int           { return len(strings.Fields(v)) }
func NormalizeLocation(v string) string { return strings.Join(strings.Fields(v), " ") }
func NormalizeName(v string) string     { return strings.TrimSpace(v) }
func CanonicalStatus(v string) string   { return strings.ToLower(strings.TrimSpace(v)) }
func CanonicalActor(v string) string    { return strings.TrimSpace(v) }
func CanonicalDevice(d Device) Device {
	d.ID = CleanID(d.ID)
	d.Name = CleanName(d.Name)
	d.Location = CleanLocation(d.Location)
	return d
}
func CanonicalRecord(r InspectionRecord) InspectionRecord {
	r.ID = CleanID(r.ID)
	r.DeviceID = CleanID(r.DeviceID)
	r.Status = CanonicalStatus(r.Status)
	r.Inspector = CanonicalActor(r.Inspector)
	r.Notes = CleanNotes(r.Notes)
	return r
}
