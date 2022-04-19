package cron

import (
	"fmt"
	"testing"
)

func TestDemoCron(t *testing.T) {
	DemoCron()
}

func TestNewMyTicker(t *testing.T) {
	tick := NewMyTick(1, testPrint)
	tick.Start()
}

func testPrint() {
	fmt.Println(" 滴答 1 次")
}
