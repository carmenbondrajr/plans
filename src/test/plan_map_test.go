package test

import (
	"plans"
	"testing"
)

func TestPlanMap(t *testing.T) {
	emptyStruct := struct{}{}

	planMap := plans.PlanMap{
		"12345": map[string]struct{}{
			"abc": emptyStruct,
			"def": emptyStruct,
		},
		"92929": map[string]struct{}{
			"fff": emptyStruct,
			"abc": emptyStruct,
		},
		"44444": map[string]struct{}{
			"hello":  emptyStruct,
			"world!": emptyStruct,
		},
	}

	inputOutput := map[string]string{
		"12345": "Plans Available:\nabc, def\n",
		"92929": "Plans Available:\nfff, abc\n",
		"44444": "Plans Available:\nhello, world!\n",
		"90210": "No Plans Available.",
	}

	for in := range inputOutput {
		expected := inputOutput[in]
		actual := planMap.GetPlansFromZip(in)
		if actual != expected {
			t.Errorf("Formatting Plans from map failed for Zip '%s': "+
				"Expected '%s' but got '%s' ", in, expected, actual)
		}
	}

}
