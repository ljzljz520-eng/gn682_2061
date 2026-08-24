package inspection

import (
	"fmt"
	"inspectionbase/internal/domain"
)

type Lifecycle struct {
	Record domain.InspectionRecord
	Events []string
}

func NewLifecycle(r domain.InspectionRecord) Lifecycle {
	return Lifecycle{Record: r, Events: []string{"created"}}
}
func (l *Lifecycle) Start() {
	l.Events = append(l.Events, "started")
	if l.Record.Status == "" {
		l.Record.Status = "pending"
	}
}
func (l *Lifecycle) Pass() {
	if l.Record.Status == "pending" {
		l.Record.Status = "passed"
		l.Events = append(l.Events, "passed")
	}
}
func (l *Lifecycle) Fail(reason string) {
	if l.Record.Status == "pending" {
		l.Record.Status = "failed"
		l.Record.Notes = reason
		l.Events = append(l.Events, "failed")
	}
}
func (l *Lifecycle) Reopen() {
	if l.Record.Status == "passed" || l.Record.Status == "failed" {
		l.Record.Status = "pending"
		l.Events = append(l.Events, "reopened")
	}
}
func (l Lifecycle) Complete() bool  { return l.Record.Status == "passed" || l.Record.Status == "failed" }
func (l Lifecycle) EventCount() int { return len(l.Events) }
func (l Lifecycle) LastEvent() string {
	if len(l.Events) == 0 {
		return ""
	}
	return l.Events[len(l.Events)-1]
}
func (l Lifecycle) HasEvent(v string) bool {
	for _, e := range l.Events {
		if e == v {
			return true
		}
	}
	return false
}
func (l Lifecycle) Summary() string {
	return fmt.Sprintf("%s:%s:%d", l.Record.ID, l.Record.Status, len(l.Events))
}
func (l *Lifecycle) AddNote(v string) {
	if v != "" {
		l.Record.Notes = v
	}
}
func (l *Lifecycle) Assign(v string) {
	if v != "" {
		l.Record.Inspector = v
	}
}
func (l *Lifecycle) SetCheckedAt(v string) { l.Record.CheckedAt = v }
func (l Lifecycle) Valid() bool            { return domain.ValidRecord(l.Record) }
func (l Lifecycle) Clone() Lifecycle {
	return Lifecycle{Record: domain.CloneRecord(l.Record), Events: append([]string{}, l.Events...)}
}
func (l *Lifecycle) MergeEvents(v []string) {
	for _, e := range v {
		if e != "" {
			l.Events = append(l.Events, e)
		}
	}
}
func (l Lifecycle) EventNames() []string { return append([]string{}, l.Events...) }
func (l Lifecycle) IsPending() bool      { return l.Record.Status == "pending" }
func (l Lifecycle) IsPassed() bool       { return l.Record.Status == "passed" }
func (l Lifecycle) IsFailed() bool       { return l.Record.Status == "failed" }
func (l *Lifecycle) SetStatus(v string) bool {
	if !domain.IsKnownStatus(v) {
		return false
	}
	l.Record.Status = v
	l.Events = append(l.Events, "status:"+v)
	return true
}
func (l Lifecycle) DeviceID() string  { return l.Record.DeviceID }
func (l Lifecycle) ID() string        { return l.Record.ID }
func (l Lifecycle) Inspector() string { return l.Record.Inspector }
func (l Lifecycle) Notes() string     { return l.Record.Notes }
func (l Lifecycle) CheckedAt() string { return l.Record.CheckedAt }
