package inspection

import "inspectionbase/internal/domain"

type Queue struct{ items []domain.InspectionRecord }

func NewQueue() Queue                           { return Queue{items: []domain.InspectionRecord{}} }
func (q *Queue) Push(r domain.InspectionRecord) { q.items = append(q.items, r) }
func (q *Queue) Pop() (domain.InspectionRecord, bool) {
	if len(q.items) == 0 {
		return domain.InspectionRecord{}, false
	}
	r := q.items[0]
	q.items = q.items[1:]
	return r, true
}
func (q *Queue) Peek() (domain.InspectionRecord, bool) {
	if len(q.items) == 0 {
		return domain.InspectionRecord{}, false
	}
	return q.items[0], true
}
func (q Queue) Len() int    { return len(q.items) }
func (q Queue) Empty() bool { return len(q.items) == 0 }
func (q *Queue) Clear()     { q.items = []domain.InspectionRecord{} }
func (q Queue) Items() []domain.InspectionRecord {
	return append([]domain.InspectionRecord{}, q.items...)
}
func (q Queue) Contains(id string) bool {
	for _, r := range q.items {
		if r.ID == id {
			return true
		}
	}
	return false
}
func (q *Queue) Remove(id string) bool {
	for i, r := range q.items {
		if r.ID == id {
			q.items = append(q.items[:i], q.items[i+1:]...)
			return true
		}
	}
	return false
}
func (q *Queue) Replace(r domain.InspectionRecord) bool {
	for i, x := range q.items {
		if x.ID == r.ID {
			q.items[i] = r
			return true
		}
	}
	return false
}
func (q Queue) Statuses() map[string]int {
	m := map[string]int{}
	for _, r := range q.items {
		m[r.Status]++
	}
	return m
}
func (q Queue) DeviceIDs() map[string]bool {
	m := map[string]bool{}
	for _, r := range q.items {
		m[r.DeviceID] = true
	}
	return m
}
func (q Queue) Pending() []domain.InspectionRecord {
	v := []domain.InspectionRecord{}
	for _, r := range q.items {
		if r.Status == "pending" {
			v = append(v, r)
		}
	}
	return v
}
func (q Queue) Terminal() []domain.InspectionRecord {
	v := []domain.InspectionRecord{}
	for _, r := range q.items {
		if domain.IsTerminalStatus(r.Status) {
			v = append(v, r)
		}
	}
	return v
}
func (q Queue) ForInspector(id string) []domain.InspectionRecord {
	v := []domain.InspectionRecord{}
	for _, r := range q.items {
		if r.Inspector == id {
			v = append(v, r)
		}
	}
	return v
}
func (q Queue) ForDevice(id string) []domain.InspectionRecord {
	v := []domain.InspectionRecord{}
	for _, r := range q.items {
		if r.DeviceID == id {
			v = append(v, r)
		}
	}
	return v
}
func (q Queue) Valid() bool {
	for _, r := range q.items {
		if r.Validate() != nil {
			return false
		}
	}
	return true
}
func (q Queue) FirstID() string {
	if len(q.items) == 0 {
		return ""
	}
	return q.items[0].ID
}
func (q Queue) LastID() string {
	if len(q.items) == 0 {
		return ""
	}
	return q.items[len(q.items)-1].ID
}
func (q Queue) Clone() Queue { return Queue{items: q.Items()} }
func (q *Queue) Append(v []domain.InspectionRecord) {
	for _, r := range v {
		q.Push(r)
	}
}
func (q Queue) CountStatus(s string) int {
	n := 0
	for _, r := range q.items {
		if r.Status == s {
			n++
		}
	}
	return n
}
func (q Queue) AllPassed() bool  { return len(q.items) > 0 && q.CountStatus("passed") == len(q.items) }
func (q Queue) AllFailed() bool  { return len(q.items) > 0 && q.CountStatus("failed") == len(q.items) }
func (q Queue) HasPending() bool { return q.CountStatus("pending") > 0 }
func (q Queue) Summary() string  { return string(rune(len(q.items))) }
