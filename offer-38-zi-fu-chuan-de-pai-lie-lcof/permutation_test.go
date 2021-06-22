package permutation

import "testing"

func TestPermutation(t *testing.T) {
	testCases := []struct {
		desc          string
		testcase      string
		expectedValue []string
	}{
		{
			desc:          "",
			testcase:      "abc",
			expectedValue: []string{"abc"},
		},
		{
			desc:          "",
			testcase:      "aac",
			expectedValue: []string{"abc"},
		},
		{
			desc:          "",
			testcase:      "21",
			expectedValue: []string{"12", "21"},
		},
		{
			desc:          "",
			testcase:      "111",
			expectedValue: []string{"111"},
		},
	}

	for _, tC := range testCases {
		if actualValue := permutation(tC.testcase); actualValue != nil {
			t.Log(actualValue, tC.expectedValue)
		}
	}
}
