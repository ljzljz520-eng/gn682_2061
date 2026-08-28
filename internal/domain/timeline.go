package domain

type TimelineEvent struct{ ID, RecordID, Kind, At, Actor string }

func NewTimeline(id, record, kind, at, actor string) TimelineEvent {
	return TimelineEvent{ID: id, RecordID: record, Kind: kind, At: at, Actor: actor}
}
func (e TimelineEvent) Valid() bool        { return e.ID != "" && e.RecordID != "" && e.Kind != "" }
func (e TimelineEvent) IsCreation() bool   { return e.Kind == "create" }
func (e TimelineEvent) IsTransition() bool { return e.Kind == "status" }
func (e TimelineEvent) IsExport() bool     { return e.Kind == "export" }
func (e TimelineEvent) ActorName() string {
	if e.Actor == "" {
		return "system"
	}
	return e.Actor
}
func SortTimeline(v []TimelineEvent) []TimelineEvent {
	out := append([]TimelineEvent{}, v...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].At < out[i].At {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
func TimelineFor(v []TimelineEvent, id string) []TimelineEvent {
	out := []TimelineEvent{}
	for _, e := range v {
		if e.RecordID == id {
			out = append(out, e)
		}
	}
	return out
}
func EventKinds(v []TimelineEvent) map[string]int {
	m := map[string]int{}
	for _, e := range v {
		m[e.Kind]++
	}
	return m
}
func HasTransition(v []TimelineEvent) bool {
	for _, e := range v {
		if e.IsTransition() {
			return true
		}
	}
	return false
}
func LastEvent(v []TimelineEvent) TimelineEvent {
	if len(v) == 0 {
		return TimelineEvent{}
	}
	return SortTimeline(v)[len(v)-1]
}
func TimelineCount(v []TimelineEvent) int { return len(v) }
func IsChronological(v []TimelineEvent) bool {
	for i := 1; i < len(v); i++ {
		if v[i].At < v[i-1].At {
			return false
		}
	}
	return true
}
