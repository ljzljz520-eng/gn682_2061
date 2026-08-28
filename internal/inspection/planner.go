package inspection

import (
	"inspectionbase/internal/domain"
	"sort"
)

type Plan struct {
	DeviceID string
	Due      string
	Priority int
	Status   string
}

func BuildPlan(ds []domain.Device, due string) []Plan {
	out := []Plan{}
	for _, d := range ds {
		if d.Active {
			out = append(out, Plan{DeviceID: d.ID, Due: due, Priority: 1, Status: "scheduled"})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].DeviceID < out[j].DeviceID })
	return out
}
func Prioritize(p []Plan) []Plan {
	out := append([]Plan{}, p...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Priority > out[j].Priority })
	return out
}
func MarkStarted(p Plan) Plan   { p.Status = "started"; return p }
func MarkCompleted(p Plan) Plan { p.Status = "completed"; return p }
func CancelPlan(p Plan) Plan    { p.Status = "cancelled"; return p }
func IsDue(p Plan, today string) bool {
	return p.Due != "" && p.Due <= today && p.Status != "completed" && p.Status != "cancelled"
}
func ActivePlans(p []Plan) []Plan {
	v := []Plan{}
	for _, x := range p {
		if x.Status != "cancelled" && x.Status != "completed" {
			v = append(v, x)
		}
	}
	return v
}
func PlanCount(p []Plan) int { return len(p) }
func DevicePlanned(p []Plan, id string) bool {
	for _, x := range p {
		if x.DeviceID == id {
			return true
		}
	}
	return false
}
func Reschedule(p Plan, due string) Plan { p.Due = due; return p }
