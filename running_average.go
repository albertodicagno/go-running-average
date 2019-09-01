package main

import (
	"errors"
	"math"
)

type RunningAverage struct {
	SampleCount uint
	Samples     []float64
	Average     float64
	Min         float64
	Max         float64
	counter     uint
	cursor      uint
	sum         float64
}

func NewRunningAverage(numSamples uint) *RunningAverage {
	avg := new(RunningAverage)
	avg.SampleCount = numSamples
	return avg
}

func (avg *RunningAverage) AddSample(sample float64) {
	avg.sum -= avg.Samples[avg.cursor]
	avg.Samples[avg.cursor] = sample
	avg.sum += avg.Samples[avg.cursor]
	avg.cursor++

	if avg.cursor == avg.SampleCount {
		avg.cursor = 0
	}

	if avg.counter == 0 {
		avg.Min = sample
		avg.Max = sample
	} else if sample < avg.Min {
		avg.Min = sample
	} else if sample > avg.Max {
		avg.Max = sample
	}

	if avg.counter < avg.SampleCount {
		avg.counter++
	}
}

func (avg *RunningAverage) Clear() {
	avg.Min = 0.00
	avg.Max = 0.00
	avg.sum = 0
	avg.counter = 0
	avg.cursor = 0

	for i := uint(0); i < avg.SampleCount; i++ {
		avg.Samples[i] = 0.0
	}
}

func (avg *RunningAverage) Fill(sample float64) {
	avg.Clear()
	for i := uint(0); i < avg.SampleCount; i++ {
		avg.AddSample(sample)
	}
}

func (avg *RunningAverage) GetAverage() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not yet ready")
	}

	avg.sum = 0.0

	for i := uint(0); i < avg.counter; i++ {
		avg.sum += avg.Samples[i]
	}
	return avg.sum / float64(avg.counter), nil
}


func (avg *RunningAverage) GetFastAverage() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not yet ready")
	}

	return avg.sum / float64(avg.counter), nil
}

func (avg *RunningAverage) GetStandardDeviation() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not yet ready")
	}
	var temp float64
	var average float64
	var err error

	average, err = avg.GetFastAverage()

	if err != nil {
		return -1, err
	}

	for i := uint(0); i < avg.counter;  i++{
		temp += math.Pow(avg.Samples[i] - average, 2)
	}
	temp = math.Sqrt(temp / float64(avg.counter - 1))
	return temp, nil
}