package test

import (
	"plans"
	"testing"
)

func TestAddressHelper(t *testing.T) {
	inputOutput := map[string]string{
		"176 Rodriguez Junctions Apt. 773 New Jamie, UT 87198": "87198",
		"USCGC Rodriguez FPO AP 47557-8623":                    "47557",
		"3640 Pittman Station East Grant, MD 27908-3392":       "27908",
		"0566 Villegas Springs Suite 067	Josefurt, AL 15004": "15004",
		"19010": "19010",
		"":      "",
		"-":     "",
	}

	for in := range inputOutput {
		expected := inputOutput[in]
		actual := plans.GetZipFromAddress(in)
		if actual != expected {
			t.Errorf("Parsing zip from address string failed on input '%s': "+
				"Expected '%s' but got '%s' ", in, expected, actual)
		}
	}
}
