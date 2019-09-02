package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAverage(t *testing.T) {
	ravg := NewRunningAverage(8)
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
