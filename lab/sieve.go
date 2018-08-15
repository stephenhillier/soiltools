package lab

import "errors"

// Sieve is a type that represents a single sieve used in a test
type Sieve struct {
	Size float64 // size of opening in mm
	Mass float64 // mass retained, in grams
}

// SieveTest is a type that holds all data about a soil sample
// prepared for grain size analysis.
type SieveTest struct {
	InitialMass float64
	DryMass     float64
	WashedMass  float64
	Sieves      []Sieve
}

// PercentPassingSieve is the percentage of soil mass that passed
// through a single sieve.
type PercentPassingSieve struct {
	Size           float64
	MassPassing    float64
	PercentPassing float64
}

// AddSieve adds a new sieve, with mass, to the sieve test stack
func (t *SieveTest) AddSieve(size float64, mass float64) (err error) {

	if size <= 0 {
		err = errors.New("size must be positive and non zero")
		return err
	}
	if mass < 0 {
		err = errors.New("mass must be positive")
		return err
	}

	position := 0

	// keep the slice of sieves sorted - find position for the new sieve
	for _, sieve := range t.Sieves {
		if size > sieve.Size {
			break
		}
		position++
	}

	t.Sieves = append(t.Sieves, Sieve{})

	if position < len(t.Sieves) {
		copy(t.Sieves[position+1:], t.Sieves[position:])
	}

	t.Sieves[position] = Sieve{size, mass}

	return nil
}

// RemoveSieve removes a single sieve at a specified position
// from the set of sieves in a SieveTest
func (t *SieveTest) RemoveSieve(position int) {
	t.Sieves = append(t.Sieves[:position], t.Sieves[position+1:]...)
}

// Passing calculates the percent mass passing each sieve in a sieve test
func (t *SieveTest) Passing() (result []PercentPassingSieve, err error) {

	if len(t.Sieves) == 0 {
		err = errors.New("no sieves added to test")
		return result, err
	}

	if t.DryMass <= 0 {
		err = errors.New("dry mass must be positive and non-zero")
		return result, err
	}

	var cumulativeMass float64

	for _, sieve := range t.Sieves {
		massPassing := t.DryMass - sieve.Mass - cumulativeMass
		percentPassing := massPassing / t.DryMass
		cumulativeMass += sieve.Mass

		result = append(result, PercentPassingSieve{sieve.Size, massPassing, percentPassing})
	}

	return result, nil
}
