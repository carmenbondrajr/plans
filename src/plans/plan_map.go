package plans

import (
	"fmt"
	"strings"
)

// Although very wordy, this type maps zip codes to a
// set of Plan IDs without allocating unnecessary space.
type PlanMap map[string]map[string]struct{}

func (r PlanMap) GetPlansFromZip(zip string) string {
	plansSet := r[zip]

	plans := make([]string, 0)
	for key := range plansSet {
		plans = append(plans, key)
	}

	if len(plans) > 0 {
		return fmt.Sprintf("Plans Available:\n%s\n", strings.Join(plans, ", "))
	}

	return "No Plans Available."
}
