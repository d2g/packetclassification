package l7filter

import (
	"testing"
)

func Test_l7filterload(t *testing.T) {
	patterns, err := PatternsFromDirectory("patterns")
	if err != nil {
		t.Log("Error:" + err.Error())
		t.Fail()
	}
	if len(patterns) == 0 {
		t.Log("Patterns Not Loaded?")
		t.Fail()
	}

	for i := range patterns {
		if patterns[i].Name == "" || patterns[i].RegularExpresion == "" {
			if patterns[i].Name == "" {
				t.Log("Empty Pattern Found??")
				t.Fail()
			}

			if patterns[i].RegularExpresion == "" {
				t.Log("Empty Regex for Pattern:" + patterns[i].RegularExpresion)
				t.Fail()
			}
		}
	}
}
