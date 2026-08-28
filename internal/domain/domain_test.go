package domain

import "testing"

func TestDomainValidation(t *testing.T) {
	if (InspectionRecord{}).Validate() != ErrInvalid {
		t.Fatal()
	}
	if !ValidTransition("pending", "passed") {
		t.Fatal()
	}
	if NormalizeLimit(1000) != 100 {
		t.Fatal()
	}
}
