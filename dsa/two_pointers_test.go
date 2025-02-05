package dsa_test

import (
	"reflect"
	"testing"

	"github.com/ericbutera/code-challenges-go/dsa"
)

func TestFindConflicts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		appointments []dsa.Interval
		expected     []dsa.Interval
	}{
		{
			name: "Overlapping appointments",
			appointments: []dsa.Interval{
				{Start: 1, End: 5},
				{Start: 4, End: 6},
				{Start: 7, End: 10},
				{Start: 9, End: 11},
			},
			expected: []dsa.Interval{
				{Start: 4, End: 6},
				{Start: 9, End: 11},
			},
		},
		{
			name: "No conflicts",
			appointments: []dsa.Interval{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
				{Start: 7, End: 9},
			},
			expected: []dsa.Interval{},
		},
		{
			name: "All conflicts",
			appointments: []dsa.Interval{
				{Start: 1, End: 10},
				{Start: 2, End: 5},
				{Start: 4, End: 6},
			},
			expected: []dsa.Interval{
				{Start: 2, End: 5},
				{Start: 4, End: 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := dsa.FindConflicts(tt.appointments)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("findConflicts(%v) = %v; want %v", tt.appointments, actual, tt.expected)
			}
		})
	}
}
