package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/api/resource"
)

var (
	argTotal         = flag.String("total", "0", "total memory to be consumed. Memory will be consumed via multiple allocations.")
	argStepSize      = flag.String("alloc-size", "4Ki", "amount of memory to be consumed in each allocation")
	argSleepDuration = flag.Duration("sleep", time.Millisecond, "duration to sleep between allocations")
	buffer           [][]byte
)

func main() {
	flag.Parse()
	total := resource.MustParse(*argTotal)
	stepSize := resource.MustParse(*argStepSize)
	glog.Infof("Allocating %q memory, in %q chunks, with a %v sleep between allocations", total.String(), stepSize.String(), *argSleepDuration)
	allocateMemory(total, stepSize)
	glog.Infof("Allocated %q memory", total.String())
	select {}
}

func allocateMemory(total, stepSize resource.Quantity) {
	for i := int64(1); i*stepSize.Value() <= total.Value(); i++ {
		newBuffer := make([]byte, stepSize.Value())
		for i := range newBuffer {
			newBuffer[i] = 0
		}
		buffer = append(buffer, newBuffer)
		time.Sleep(*argSleepDuration)
	}
}
