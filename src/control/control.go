// Package control implements the SCN-01 task state machine.
package control

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Layer string

const (
	LayerLocal      Layer = "local_configuration"
	LayerDNS        Layer = "dns"
	LayerConnection Layer = "connection_or_proxy"
	LayerTarget     Layer = "target_response"
	LayerNonNetwork Layer = "non_network"
	LayerUnknown    Layer = "insufficient_evidence"
)

type BusinessResult string

const (
	BusinessRecovered    BusinessResult = "recovered"
	BusinessNotRecovered BusinessResult = "not_recovered"
	BusinessUncertain    BusinessResult = "cannot_determine"
)

type ActionResult string

const (
	ActionNotRun    ActionResult = "not_run"
	ActionSucceeded ActionResult = "succeeded"
	ActionFailed    ActionResult = "failed"
)

type Task struct {
	ID             string
	Target         string
	Symptom        string
	ObservedAt     time.Time
	PreserveEgress bool
}

func (t Task) Validate() error {
	var missing []string
	if strings.TrimSpace(t.Target) == "" {
		missing = append(missing, "target")
	}
	if strings.TrimSpace(t.Symptom) == "" {
		missing = append(missing, "symptom")
	}
	if t.ObservedAt.IsZero() {
		missing = append(missing, "observed_at")
	}
	if !t.PreserveEgress {
		missing = append(missing, "preserve_egress")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required task fields: %s", strings.Join(missing, ", "))
	}
	return nil
}

type Evidence struct {
	Layer    Layer
	Supports bool
	Detail   string
	Source   string
}

type Diagnosis struct {
	Layer      Layer
	Confidence string
	Support    []Evidence
	Counter    []Evidence
	NextCheck  string
}

func Diagnose(evidence []Evidence) Diagnosis {
	for _, item := range evidence {
		if item.Supports && item.Layer == LayerNonNetwork {
			return Diagnosis{Layer: LayerNonNetwork, Confidence: "high", Support: []Evidence{item}, NextCheck: "handoff to platform or account support"}
		}
	}
	for _, item := range evidence {
		if item.Supports && item.Layer != LayerUnknown {
			return Diagnosis{Layer: item.Layer, Confidence: "candidate", Support: []Evidence{item}, NextCheck: "collect same-task result after constrained action"}
		}
	}
	return Diagnosis{Layer: LayerUnknown, Confidence: "low", NextCheck: "collect target response or page error evidence"}
}

type Capability struct {
	ID              string
	Available       bool
	PreservesEgress bool
	Scope           string
	Rollback        string
	BeforeEgress    string
	AfterEgress     string
}

func (c Capability) Eligible(task Task, diagnosis Diagnosis) (bool, string) {
	if err := task.Validate(); err != nil {
		return false, err.Error()
	}
	if diagnosis.Layer == LayerNonNetwork || diagnosis.Layer == LayerUnknown {
		return false, "diagnosis does not support a local connection action"
	}
	if !c.Available {
		return false, "adapter action is unavailable"
	}
	if !c.PreservesEgress || c.BeforeEgress == "" || c.AfterEgress == "" || c.BeforeEgress != c.AfterEgress {
		return false, "egress preservation is not proven"
	}
	if c.Scope != "current_profile" {
		return false, "action scope exceeds current profile"
	}
	if c.Rollback == "" {
		return false, "rollback is missing"
	}
	return true, ""
}

type Adapter interface {
	Capability() Capability
	Execute() error
}

type Probe func(Task) bool

type Record struct {
	Task          Task
	Diagnosis     Diagnosis
	Action        ActionResult
	Business      BusinessResult
	HandoffReason string
	Revisions     []string
}

func Run(task Task, evidence []Evidence, adapter Adapter, confirmed bool, probe Probe) Record {
	diagnosis := Diagnose(evidence)
	record := Record{Task: task, Diagnosis: diagnosis, Action: ActionNotRun, Business: BusinessUncertain}
	if err := task.Validate(); err != nil {
		record.HandoffReason = err.Error()
		return record
	}
	if diagnosis.Layer == LayerNonNetwork || diagnosis.Layer == LayerUnknown {
		record.HandoffReason = "no safe network action"
		return record
	}
	eligible, reason := adapter.Capability().Eligible(task, diagnosis)
	if !eligible {
		record.HandoffReason = reason
		return record
	}
	if !confirmed {
		record.HandoffReason = "action cancelled by user"
		return record
	}
	if err := adapter.Execute(); err != nil {
		record.Action = ActionFailed
		record.HandoffReason = "action failed: " + err.Error()
		return record
	}
	record.Action = ActionSucceeded
	if probe(task) {
		record.Business = BusinessRecovered
		return record
	}
	record.Business = BusinessNotRecovered
	record.HandoffReason = "action completed but same task did not recover"
	return record
}

func (r *Record) Revise(reason string) error {
	if strings.TrimSpace(reason) == "" {
		return errors.New("revision reason is required")
	}
	r.Revisions = append(r.Revisions, reason)
	return nil
}
