package lab

import "testing"

func TestAddSieve(t *testing.T) {
	newTest := SieveTest{}
	testMass := 215.8
	newTest.AddSieve(19, testMass)

	if newTest.Sieves[0].Mass != testMass {
		t.Errorf("Mass of new sieve was incorrect. got: %v, want: %v", newTest.Sieves[0].Mass, testMass)
	}
}

func TestRemoveSieve(t *testing.T) {
	newTest := SieveTest{}
	newTest.AddSieve(19, 215.8)
	if len(newTest.Sieves) != 1 {
		t.Error("Trying to test removing sieves but a sieve was not added to set")
	}

	newTest.RemoveSieve(0)

	if len(newTest.Sieves) != 0 {
		t.Error("Sieve was not removed")
	}
}

func TestCalculatePassing(t *testing.T) {
	newTest := SieveTest{2100.0, 2000.0, 1900.0, []Sieve{}}

	// define sieves with a size, mass, and the expected results
	cases := []struct {
		size               float64
		mass               float64
		wantMassPassing    float64
		wantPercentPassing float64
	}{
		{20, 0, 2000, 1.00},  // 20 mm sieve, 0 mass, 400 grams passing and 100 percent passing
		{16, 100, 1900, .95}, // 16 mm sieve with 100g mass
		{12, 100, 1800, .90},
		{5, 100, 1700, .85},
		{2, 100, 1600, .80},
	}

	// add sieves to test
	for _, sieve := range cases {
		newTest.AddSieve(sieve.size, sieve.mass)
	}

	result, err := newTest.Passing()

	if err != nil {
		t.Error("Error calculating percent passing:", err)
	}

	for i, testCase := range cases {
		if result[i].MassPassing != testCase.wantMassPassing {
			t.Errorf("Mass passing on %v mm sieve incorrect. tested size: %v, got: %v, want: %v", testCase.size, result[i].Size, result[i].MassPassing, testCase.wantMassPassing)
		}
		if result[i].PercentPassing != testCase.wantPercentPassing {
			t.Errorf("Percent passing on %v mm sieve incorrect. got: %v, want: %v", testCase.size, result[i].PercentPassing, testCase.wantPercentPassing)
		}
	}
	if result[0].MassPassing != 400 {
	}
}
