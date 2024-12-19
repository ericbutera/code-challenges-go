package learn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func StratFetcher(s FetchStrategy) string {
	return s.URL()
}

type FetchStrategy interface {
	URL() string
}

type IncrementalStrategy struct{}

func (i *IncrementalStrategy) URL() string {
	return "incremental-url"
}

type FullStrategy struct{}

func (f *FullStrategy) URL() string {
	return "full-url"
}

func TestStrategy(t *testing.T) {
	t.Parallel()
	full := &FullStrategy{}
	assert.Equal(t, "full-url", full.URL())
	inc := &IncrementalStrategy{}
	assert.Equal(t, "incremental-url", inc.URL())
}

func TestFetcher(t *testing.T) {
	t.Parallel()
	full := &FullStrategy{}
	assert.Equal(t, "full-url", StratFetcher(full))
}
