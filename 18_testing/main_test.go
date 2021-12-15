package testing

import (
	"testing"
)

func TestAddition(t *testing.T){
	testcases := []struct{
		name string
		input int
		expected int
	}{
		{"positive", 1, 3},
		{"zero", 0, 2},
		{"negative", -1, 1},
		{"long", 9999, 10001},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			// Addition function is identified by the same ackage name "testing" on line 1
			if res, _ := Addition(test.input); test.expected != res{
				t.Error("Addition result is not correct!")
			}
		})
	}
}