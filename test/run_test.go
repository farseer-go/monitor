package test

import (
	"context"
	"fmt"
	"github.com/farseer-go/monitor"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	assert.Panics(t, func() {
		monitor.Run("monitorRun", 0, func(context *monitor.MonitorContext) {

		}, context.Background())
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	now := time.Now()
	monitor.Run("monitorRun", 10*time.Millisecond, func(context *monitor.MonitorContext) {
		s := time.Since(now) - 10*time.Millisecond
		if s >= 6*time.Millisecond || s < time.Nanosecond {
			t.Fatal(fmt.Sprintf("时间不对:%d", s.Milliseconds()))
		}
		now = time.Now()
	}, ctx)

	time.Sleep(500 * time.Millisecond)
}

func TestRunNow(t *testing.T) {
	assert.Panics(t, func() {
		monitor.RunNow("monitorRunNow", 0, func(context *monitor.MonitorContext) {

		}, context.Background())
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	now := time.Now().Add(-10 * time.Millisecond)
	monitor.RunNow("monitorRunNow", 10*time.Millisecond, func(context *monitor.MonitorContext) {
		s := time.Since(now) - 10*time.Millisecond
		if s >= 6*time.Millisecond || s < time.Nanosecond {
			t.Fatal(fmt.Sprintf("时间不对:%d", s.Milliseconds()))
		}
		now = time.Now()
	}, ctx)

	time.Sleep(500 * time.Millisecond)
}

func TestRunPanic(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	monitor.RunNow("monitorRun", 500*time.Millisecond, func(context *monitor.MonitorContext) {
		context.SetNextDuration(0)
	}, ctx)
	monitor.RunNow("monitorRun", 500*time.Millisecond, func(context *monitor.MonitorContext) {
		context.SetNextTime(time.Now().Add(-1 * time.Second))
	}, ctx)

	time.Sleep(50 * time.Millisecond)
}

func TestSetNextDuration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	now := time.Now().Add(20 * time.Millisecond)
	monitor.Run("monitorRun", 20*time.Millisecond, func(context *monitor.MonitorContext) {
		s := time.Since(now)
		if s >= 6*time.Millisecond || s < time.Nanosecond {
			t.Fatal(fmt.Sprintf("时间不对:%d", s.Milliseconds()))
		}
		now = time.Now().Add(10 * time.Millisecond)
		context.SetNextDuration(10 * time.Millisecond)
	}, ctx)

	time.Sleep(500 * time.Millisecond)
}

func TestSetNextTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	now := time.Now().Add(20 * time.Millisecond)
	monitor.Run("monitorRun", 20*time.Millisecond, func(context *monitor.MonitorContext) {
		s := time.Since(now)
		if s >= 6*time.Millisecond || s < time.Nanosecond {
			t.Fatal(fmt.Sprintf("时间不对:%d", s.Milliseconds()))
		}
		now = time.Now().Add(10 * time.Millisecond)
		context.SetNextTime(now)
	}, ctx)

	time.Sleep(500 * time.Millisecond)
}
