[![Build Status](https://dev.azure.com/albertodicagno-dev/go-runningaverage_test/_apis/build/status/albertodicagno.runningaverage?branchName=master)](https://dev.azure.com/albertodicagno-dev/go-runningaverage_test/_build/latest?definitionId=4&branchName=master)
# Go Running average
![Running gopher](https://golang.org/doc/gopher/run.png)
Simple and fast Go library for performing running average and standard deviation computations.

```go
package main

import (
	"fmt"
	ravg "github.com/albertodicagno/runningaverage"
	"log"
)

const howManySamples = uint(32)

func main() {
	avg := ravg.NewRunningAverage(howManySamples)
	for i := 0; i < 1000; i++ {
		avg.AddSample(float64(i))
	}
	computedAverage, err := avg.GetAverage()
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Printf("Average: %.2f\n", computedAverage)
}
```
