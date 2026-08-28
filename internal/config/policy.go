package config

type Policy struct {
	RequireInspector, RequireNotesForFailure, AllowReopen, AuditEnabled bool
	MaxBatch                                                            int
}

func DefaultPolicy() Policy {
	return Policy{RequireInspector: true, RequireNotesForFailure: true, AllowReopen: true, AuditEnabled: true, MaxBatch: 100}
}
func (p Policy) Validate() bool                  { return p.MaxBatch > 0 && p.MaxBatch <= 1000 }
func (p Policy) CanCreate() bool                 { return true }
func (p Policy) CanUpdate() bool                 { return true }
func (p Policy) CanReopen() bool                 { return p.AllowReopen }
func (p Policy) NeedsAudit() bool                { return p.AuditEnabled }
func (p Policy) CheckBatch(n int) bool           { return n >= 0 && n <= p.MaxBatch }
func (p Policy) CheckInspector(v string) bool    { return !p.RequireInspector || v != "" }
func (p Policy) CheckFailureNotes(v string) bool { return !p.RequireNotesForFailure || v != "" }
func (p Policy) WithMaxBatch(n int) Policy {
	if n > 0 {
		p.MaxBatch = n
	}
	return p
}
func (p Policy) DisableAudit() Policy  { p.AuditEnabled = false; return p }
func (p Policy) EnableReopen() Policy  { p.AllowReopen = true; return p }
func (p Policy) DisableReopen() Policy { p.AllowReopen = false; return p }
func (p Policy) Summary() map[string]bool {
	return map[string]bool{"inspector": p.RequireInspector, "notes": p.RequireNotesForFailure, "reopen": p.AllowReopen, "audit": p.AuditEnabled}
}
func (p Policy) Flags() []string {
	v := []string{}
	if p.RequireInspector {
		v = append(v, "inspector")
	}
	if p.RequireNotesForFailure {
		v = append(v, "notes")
	}
	if p.AllowReopen {
		v = append(v, "reopen")
	}
	if p.AuditEnabled {
		v = append(v, "audit")
	}
	return v
}
func MergePolicy(a, b Policy) Policy {
	p := a
	if b.MaxBatch > 0 {
		p.MaxBatch = b.MaxBatch
	}
	p.RequireInspector = p.RequireInspector || b.RequireInspector
	p.RequireNotesForFailure = p.RequireNotesForFailure || b.RequireNotesForFailure
	p.AllowReopen = p.AllowReopen && b.AllowReopen
	p.AuditEnabled = p.AuditEnabled && b.AuditEnabled
	return p
}
func IsStrict(p Policy) bool { return p.RequireInspector && p.RequireNotesForFailure && p.AuditEnabled }
func RelaxedPolicy() Policy  { return Policy{MaxBatch: 1000, AllowReopen: true} }
func PolicyName(p Policy) string {
	if IsStrict(p) {
		return "strict"
	}
	return "custom"
}
