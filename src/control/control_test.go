package control

import (
	"errors"
	"testing"
	"time"
)

type fakeAdapter struct {
	capability Capability
	err        error
	calls      int
}

func (a *fakeAdapter) Capability() Capability { return a.capability }
func (a *fakeAdapter) Execute() error         { a.calls++; return a.err }

func validTask() Task {
	return Task{ID: "task-1", Target: "https://demo.invalid/orders", Symptom: "timeout", ObservedAt: time.Date(2026, 9, 10, 9, 0, 0, 0, time.UTC), PreserveEgress: true}
}
func validAdapter() *fakeAdapter {
	return &fakeAdapter{capability: Capability{ID: "reconnect_current_egress", Available: true, PreservesEgress: true, Scope: "current_profile", Rollback: "reload prior profile", BeforeEgress: "egress-a", AfterEgress: "egress-a"}}
}
func connectionEvidence() []Evidence {
	return []Evidence{{Layer: LayerConnection, Supports: true, Detail: "simulated session failure", Source: "simulation"}}
}

func TestTaskRequiresSpecificInput(t *testing.T) {
	record := Run(Task{}, nil, validAdapter(), true, func(Task) bool { return true })
	if record.HandoffReason == "" || record.Action != ActionNotRun {
		t.Fatalf("expected validation handoff, got %#v", record)
	}
}

func TestNonNetworkEvidenceDoesNotExecute(t *testing.T) {
	adapter := validAdapter()
	record := Run(validTask(), []Evidence{{Layer: LayerNonNetwork, Supports: true, Detail: "401", Source: "simulation"}}, adapter, true, func(Task) bool { return true })
	if adapter.calls != 0 || record.Diagnosis.Layer != LayerNonNetwork {
		t.Fatalf("unexpected execution: %#v", record)
	}
}

func TestIneligibleEgressDoesNotExecute(t *testing.T) {
	adapter := validAdapter()
	adapter.capability.AfterEgress = "egress-b"
	record := Run(validTask(), connectionEvidence(), adapter, true, func(Task) bool { return true })
	if adapter.calls != 0 || record.HandoffReason != "egress preservation is not proven" {
		t.Fatalf("expected egress rejection: %#v", record)
	}
}

func TestCancellationDoesNotExecute(t *testing.T) {
	adapter := validAdapter()
	record := Run(validTask(), connectionEvidence(), adapter, false, func(Task) bool { return true })
	if adapter.calls != 0 || record.HandoffReason != "action cancelled by user" {
		t.Fatalf("expected cancellation: %#v", record)
	}
}

func TestActionSuccessDoesNotImplyRecovery(t *testing.T) {
	adapter := validAdapter()
	record := Run(validTask(), connectionEvidence(), adapter, true, func(Task) bool { return false })
	if record.Action != ActionSucceeded || record.Business != BusinessNotRecovered {
		t.Fatalf("expected separate action and business results: %#v", record)
	}
}

func TestActionFailureIsRecorded(t *testing.T) {
	adapter := validAdapter()
	adapter.err = errors.New("simulated adapter failure")
	record := Run(validTask(), connectionEvidence(), adapter, true, func(Task) bool { return true })
	if record.Action != ActionFailed || record.Business != BusinessUncertain {
		t.Fatalf("expected failed action: %#v", record)
	}
}

func TestRevisionRequiresReason(t *testing.T) {
	record := Record{}
	if err := record.Revise(""); err == nil {
		t.Fatal("expected revision validation")
	}
	if err := record.Revise("corrected source"); err != nil || len(record.Revisions) != 1 {
		t.Fatalf("unexpected revision: %#v, %v", record, err)
	}
}
