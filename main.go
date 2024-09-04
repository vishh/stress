package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	argMemTotal         = flag.String("mem-total", "0", "total memory to be consumed. Memory will be consumed via multiple allocations.")
	argMemStepSize      = flag.String("mem-alloc-size", "4Ki", "amount of memory to be consumed in each allocation")
	argMemSleepDuration = flag.Duration("mem-alloc-sleep", time.Millisecond, "duration to sleep between allocations")
	argCpus             = flag.Int("cpus", 0, "total number of CPUs to utilize")
	buffer              [][]byte
)

func main() {
	flag.Parse()
	total := resource.MustParse(*argMemTotal)
	stepSize := resource.MustParse(*argMemStepSize)
	glog.Infof("Allocating %q memory, in %q chunks, with a %v sleep between allocations", total.String(), stepSize.String(), *argMemSleepDuration)
	burnCPU()
	allocateMemory(total, stepSize)
	glog.Infof("Allocated %q memory", total.String())
	select {}
}

func burnCPU() {
	src, err := os.Open("/dev/zero")
	if err != nil {
		glog.Fatalf("failed to open /dev/zero")
	}
	for i := 0; i < *argCpus; i++ {
		glog.Infof("Spawning a thread to consume CPU")
		go func() {
			_, err := io.Copy(ioutil.Discard, src)
			if err != nil {
				glog.Fatalf("failed to copy from /dev/zero to /dev/null: %v", err)
			}
		}()
	}
}

func allocateMemory(total, stepSize resource.Quantity) {
	for i := int64(1); i*stepSize.Value() <= total.Value(); i++ {
		newBuffer := make([]byte, stepSize.Value())
		for i := range newBuffer {
			newBuffer[i] = 0
		}
		buffer = append(buffer, newBuffer)
		time.Sleep(*argMemSleepDuration)
	}
}
