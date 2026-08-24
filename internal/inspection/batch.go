package inspection

import (
	"fmt"
	"inspectionbase/internal/domain"
)

func (s *Service) CreateBatch(rs []domain.InspectionRecord) error {
	for i, r := range rs {
		if r.ID == "" {
			r.ID = fmt.Sprintf("batch-%d", i)
		}
		if e := s.Create(r); e != nil {
			return e
		}
	}
	return nil
}
func (s *Service) UpdateBatch(ids []string, status, actor string) map[string]error {
	out := map[string]error{}
	for _, id := range ids {
		out[id] = s.UpdateStatus(id, status, actor)
	}
	return out
}
func (s *Service) FindByDevice(id string) ([]domain.InspectionRecord, error) {
	return s.List(domain.QueryFilter{DeviceID: id, Limit: 100})
}
func (s *Service) Latest(id string) (domain.InspectionRecord, error) {
	v, e := s.FindByDevice(id)
	if e != nil {
		return domain.InspectionRecord{}, e
	}
	if len(v) == 0 {
		return domain.InspectionRecord{}, domain.ErrNotFound
	}
	return v[len(v)-1], nil
}
