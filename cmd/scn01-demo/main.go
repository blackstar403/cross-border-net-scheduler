package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/liutingyu/cross-border-net-scheduler/src/control"
)

type simulatedAdapter struct{}

func (simulatedAdapter) Capability() control.Capability {
	return control.Capability{ID: "reconnect_current_egress", Available: true, PreservesEgress: true, Scope: "current_profile", Rollback: "restore simulated session", BeforeEgress: "demo-egress", AfterEgress: "demo-egress"}
}
func (simulatedAdapter) Execute() error { return nil }

func main() {
	task := control.Task{ID: "demo-scn-01", Target: "https://demo.invalid/backend", Symptom: "simulated proxy session timeout", ObservedAt: time.Now().UTC(), PreserveEgress: true}
	record := control.Run(task, []control.Evidence{{Layer: control.LayerConnection, Supports: true, Detail: "simulated connection failure", Source: "simulation"}}, simulatedAdapter{}, true, func(control.Task) bool { return true })
	output, _ := json.MarshalIndent(record, "", "  ")
	fmt.Println(string(output))
}
