package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RestClient struct{}

func (r *RestClient) Get(url string) string {
	return "get " + url
}

type Fetcher interface {
	Url() string
	Fetch() string
	Save() error
}

type Base struct {
	RestClient //*resty.Client
	Fetcher
}

type Full struct {
	Base
}

func (f *Full) Url() string {
	return "full-url"
}

func (f *Full) Fetch() string {
	return f.RestClient.Get(f.Url())
}

func (f *Full) Save() error {
	return nil
}

// TODO: incremental
// type Incremental struct {
// 	Base
// }

func TestFull(t *testing.T) {
	full := &Full{
		Base: Base{},
	}
	assert.Equal(t, full.Url(), "full-url")
	assert.Equal(t, full.Fetch(), "get full-url")
}
