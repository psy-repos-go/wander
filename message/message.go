package message

import (
	"wander/nomad"
)

// NomadJobsMsg is a message for nomad jobs
type NomadJobsMsg []nomad.JobResponseEntry

// NomadAllocationMsg is a message for nomad allocations
type NomadAllocationMsg []nomad.AllocationRowEntry

// NomadLogsMsg is a message for nomad allocations
type NomadLogsMsg []string

// ErrMsg is an error message
type ErrMsg struct{ err error }

func (e ErrMsg) Error() string { return e.err.Error() }