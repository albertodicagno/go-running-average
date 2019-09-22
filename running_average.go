package runningaverage

import (
	"errors"
	"math"
)

//RunningAverage defines the data structure holding samples, min/max and average values, along with internal values and counters
type RunningAverage struct {
	SampleCount uint64
	Samples     []float64
	Average     float64
	Min         float64
	Max         float64
	counter     uint64
	cursor      uint64
	sum         float64
}

//NewRunningAverage returns a pointer to a new RunningAverage object with sample size of numSamples
func NewRunningAverage(numSamples uint64) *RunningAverage {
	avg := new(RunningAverage)
	avg.SampleCount = numSamples
	avg.Samples = make([]float64, avg.SampleCount)
	return avg
}

//AddSample adds a new sample to the RunningAverage struct
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

//Clear resets the whole RunningAverage struct
func (avg *RunningAverage) Clear() {
	avg.Min = 0.00
	avg.Max = 0.00
	avg.sum = 0
	avg.counter = 0
	avg.cursor = 0

	for i := uint64(0); i < avg.SampleCount; i++ {
		avg.Samples[i] = 0.0
	}
}

//Fill fills the sample buffer with the value passed as an argument
func (avg *RunningAverage) Fill(sample float64) {
	avg.Clear()
	for i := uint64(0); i < avg.SampleCount; i++ {
		avg.AddSample(sample)
	}
}

//GetAverage computes and returns the average iterating across the whole sample buffer
func (avg *RunningAverage) GetAverage() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not yet ready")
	}

	avg.sum = 0.0

	for i := uint64(0); i < avg.counter; i++ {
		avg.sum += avg.Samples[i]
	}
	return avg.sum / float64(avg.counter), nil
}

//GetFastAverage same as GetAverage, but faster since no iteration is made, just using the last computed values for sum and sample count
func (avg *RunningAverage) GetFastAverage() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not yet ready")
	}

	return avg.sum / float64(avg.counter), nil
}

//GetStandardDeviation returns the stdev across the sample buffer
func (avg *RunningAverage) GetStandardDeviation() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not ready")
	}
	var temp float64
	var average float64
	var err error

	average, err = avg.GetFastAverage()

	if err != nil {
		return -1, err
	}

	for i := uint64(0); i < avg.counter; i++ {
		temp += math.Pow(avg.Samples[i]-average, 2)
	}
	temp = math.Sqrt(temp / float64(avg.counter-1))
	return temp, nil
}
