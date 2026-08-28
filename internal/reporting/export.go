package reporting

import (
	"encoding/csv"
	"inspectionbase/internal/domain"
	"io"
)

func WriteCSV(w io.Writer, rs []domain.InspectionRecord) error {
	c := csv.NewWriter(w)
	if e := c.Write([]string{"id", "device_id", "status", "inspector"}); e != nil {
		return e
	}
	for _, r := range rs {
		if e := c.Write([]string{r.ID, r.DeviceID, r.Status, r.Inspector}); e != nil {
			return e
		}
	}
	c.Flush()
	return c.Error()
}
func GroupByStatus(rs []domain.InspectionRecord) map[string][]domain.InspectionRecord {
	out := map[string][]domain.InspectionRecord{}
	for _, r := range rs {
		out[r.Status] = append(out[r.Status], r)
	}
	return out
}
func DeviceCounts(rs []domain.InspectionRecord) map[string]int {
	out := map[string]int{}
	for _, r := range rs {
		out[r.DeviceID]++
	}
	return out
}
