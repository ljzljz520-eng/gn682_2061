package inspection

import "strings"

type ChecklistItem struct {
	Code, Label string
	Required    bool
}
type Checklist struct {
	DeviceID string
	Items    []ChecklistItem
}

func DefaultChecklist(device string) Checklist {
	return Checklist{DeviceID: device, Items: []ChecklistItem{{"power", "电源", true}, {"safety", "安全", true}, {"clean", "清洁", false}, {"noise", "噪音", false}}}
}
func (c Checklist) Valid() bool {
	if c.DeviceID == "" || len(c.Items) == 0 {
		return false
	}
	for _, i := range c.Items {
		if i.Code == "" || i.Label == "" {
			return false
		}
	}
	return true
}
func (c Checklist) RequiredCodes() []string {
	v := []string{}
	for _, i := range c.Items {
		if i.Required {
			v = append(v, i.Code)
		}
	}
	return v
}
func Complete(c Checklist, answers map[string]bool) bool {
	for _, i := range c.Items {
		if i.Required && !answers[i.Code] {
			return false
		}
	}
	return true
}
func NormalizeAnswer(v string) bool {
	v = strings.ToLower(strings.TrimSpace(v))
	return v == "yes" || v == "ok" || v == "true" || v == "通过"
}
func ItemCodes(c Checklist) map[string]bool {
	m := map[string]bool{}
	for _, i := range c.Items {
		m[i.Code] = true
	}
	return m
}
func Missing(c Checklist, a map[string]bool) []string {
	v := []string{}
	for _, i := range c.Items {
		if i.Required && !a[i.Code] {
			v = append(v, i.Code)
		}
	}
	return v
}
func CountRequired(c Checklist) int {
	n := 0
	for _, i := range c.Items {
		if i.Required {
			n++
		}
	}
	return n
}
func ChecklistForDevice(id string) Checklist { return DefaultChecklist(id) }
func MergeChecklist(a, b Checklist) Checklist {
	out := Checklist{DeviceID: a.DeviceID, Items: append([]ChecklistItem{}, a.Items...)}
	for _, x := range b.Items {
		if !ItemCodes(out)[x.Code] {
			out.Items = append(out.Items, x)
		}
	}
	return out
}
