package socket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeSleeper struct{}

func (f *fakeSleeper) Sleep(duration time.Duration) {
}

func Test_newProcessPercent(t *testing.T) {
	assert.Implements(t, (*Percentable)(nil), newProcessPercent(nil, &fakeSleeper{}))
}

type ppm struct {
	calledTimes int
}

func (p *ppm) SendProcessPercent(s string) {
	p.calledTimes++
}

func Test_processPercent_Add(t *testing.T) {
	p := &ppm{}
	percent := newProcessPercent(p, &fakeSleeper{})
	percent.Add()
	assert.Equal(t, 1, p.calledTimes)
}

func Test_processPercent_Current(t *testing.T) {
	p := &ppm{}
	percent := newProcessPercent(p, &fakeSleeper{})
	assert.Equal(t, int64(0), percent.Current())
	percent.Add()
	assert.Equal(t, int64(1), percent.Current())
	for i := 0; i < 100; i++ {
		percent.Add()
	}
	assert.Equal(t, int64(100), percent.Current())
	assert.Equal(t, 100, p.calledTimes)
}

func Test_processPercent_To(t *testing.T) {
	p := &ppm{}
	percent := newProcessPercent(p, &fakeSleeper{})
	percent.To(100)
	assert.Equal(t, int64(100), percent.Current())
}
