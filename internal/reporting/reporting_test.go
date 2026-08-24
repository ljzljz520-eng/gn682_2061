package reporting

import (
	"inspectionbase/internal/domain"
	"testing"
)

func TestReportSummary(t *testing.T) {
	v := Build([]domain.InspectionRecord{{ID: "r", Status: "passed"}})
	if v.Metrics.Passed != 1 {
		t.Fatal()
	}
}
