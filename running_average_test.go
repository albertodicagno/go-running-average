package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ravg *RunningAverage

func TestGetAverage(t *testing.T) {
	ravg = NewRunningAverage(8)
	ravg.AddSample(1)
	ravg.AddSample(2)
	ravg.AddSample(3)
	ravg.AddSample(4)
	ravg.AddSample(1)
	ravg.AddSample(2)
	ravg.AddSample(3)
	ravg.AddSample(4)

	_, err := ravg.GetAverage()
	assert.Nil(t, err)
}

func TestClear(t *testing.T) {
	ravg.Clear()
	assert.Equal(t, 0.0, ravg.Samples[4])
	assert.Equal(t, 0.0, ravg.Min)
	assert.Equal(t, 0.0, ravg.Max)
	assert.Equal(t, 0.0, ravg.Average)
}
