package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func StratFetcher(s FetchStrategy) string {
	return s.Url()
}

type FetchStrategy interface {
	Url() string
}

type IncrementalStrategy struct {
}

func (i *IncrementalStrategy) Url() string {
	return "incremental-url"
}

type FullStrategy struct {
}

func (f *FullStrategy) Url() string {
	return "full-url"
}

func TestStrategy(t *testing.T) {
	full := &FullStrategy{}
	assert.Equal(t, full.Url(), "full-url")
	inc := &IncrementalStrategy{}
	assert.Equal(t, inc.Url(), "incremental-url")
}

func TestFetcher(t *testing.T) {
	full := &FullStrategy{}
	assert.Equal(t, StratFetcher(full), "full-url")
}
