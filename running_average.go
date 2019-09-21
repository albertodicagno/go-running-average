package main

import (
	"errors"
	"math"
)

<<<<<<< HEAD
//RunningAverage structure holding average and other values
=======
//RunningAverage defines the data structure holding samples, min/max and average values, along with internal values and counters
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
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

<<<<<<< HEAD
//NewRunningAverage allocates memory and returns a pointer for a RunningAverage object with numSamples samples
=======
//NewRunningAverage returns a pointer to a new RunningAverage object with sample size of numSamples
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
func NewRunningAverage(numSamples uint) *RunningAverage {
	avg := new(RunningAverage)
	avg.SampleCount = numSamples
	avg.Samples = make([]float64, avg.SampleCount)
	return avg
}

<<<<<<< HEAD
//AddSample adds a new sample of type float64 to the RunningAverage buffer
=======
//AddSample adds a new sample to the RunningAverage struct
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
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

<<<<<<< HEAD
//Clear resets the average
=======
//Clear resets the whole RunningAverage struct
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
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

<<<<<<< HEAD
//Fill fills buffer with specified sample
=======
//Fill fills the sample buffer with the value passed as an argument
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
func (avg *RunningAverage) Fill(sample float64) {
	avg.Clear()
	for i := uint(0); i < avg.SampleCount; i++ {
		avg.AddSample(sample)
	}
}

<<<<<<< HEAD
//GetAverage returns computed average iterating through the sample buffer
=======
//GetAverage computes and returns the average iterating across the whole sample buffer
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
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

<<<<<<< HEAD
//GetFastAverage returns the last computed average value without performing iterations
=======
//GetFastAverage same as GetAverage, but faster since no iteration is made, just using the last computed values for sum and sample count
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
func (avg *RunningAverage) GetFastAverage() (float64, error) {
	if avg.counter == 0 {
		return -1, errors.New("average is not yet ready")
	}

	return avg.sum / float64(avg.counter), nil
}

<<<<<<< HEAD
//GetStandardDeviation computes and returns the stdev
=======
//GetStandardDeviation returns the stdev across the sample buffer
>>>>>>> e0aeaf130610e7e1b21376730e16e7e66be7756c
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

	for i := uint(0); i < avg.counter; i++ {
		temp += math.Pow(avg.Samples[i]-average, 2)
	}
	temp = math.Sqrt(temp / float64(avg.counter-1))
	return temp, nil
}
