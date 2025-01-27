// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2023 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package aa

type Ptrace struct {
	Qualifier
	Access string
	Peer   string
}

func PtraceFromLog(log map[string]string) ApparmorRule {
	return &Ptrace{
		Qualifier: NewQualifierFromLog(log),
		Access:    maskToAccess[log["requested_mask"]],
		Peer:      log["peer"],
	}
}

func (r *Ptrace) Less(other any) bool {
	o, _ := other.(*Ptrace)
	if r.Qualifier.Equals(o.Qualifier) {
		if r.Access == o.Access {
			return r.Peer == o.Peer
		}
		return r.Access < o.Access
	}
	return r.Qualifier.Less(o.Qualifier)
}

func (r *Ptrace) Equals(other any) bool {
	o, _ := other.(*Ptrace)
	return r.Access == o.Access && r.Peer == o.Peer &&
		r.Qualifier.Equals(o.Qualifier)
}
