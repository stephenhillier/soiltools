package lab

import "time"

// Sample represents a unit of soil collected for testing
type Sample struct {
	SampledBy   string
	SampledDate time.Time
}
