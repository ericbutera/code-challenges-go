package dsa

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func FindConflicts(appointments []Interval) []Interval {
	// Sort intervals by start time
	sort.Slice(appointments, func(i, j int) bool {
		return appointments[i].Start < appointments[j].Start
	})

	conflicts := []Interval{}

	// Two pointers
	i, j := 0, 1
	for j < len(appointments) {
		// Check for overlap between i and j
		if appointments[j].Start < appointments[i].End {
			// Record the conflicting interval
			conflicts = append(conflicts, appointments[j])
			// Move pointer j to check further
			j++
		} else {
			// No overlap, move pointer i to j, and j to the next interval
			i = j
			j++
		}
	}

	return conflicts
}
