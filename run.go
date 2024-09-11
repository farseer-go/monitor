package monitor

import (
	"context"
	"github.com/farseer-go/fs/asyncLocal"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/stopwatch"
	"github.com/farseer-go/fs/trace"
	"time"
)

// Run 运行一个监控，运行前先休眠
// interval:监控运行的间隔时间
// monitorFn:要运行的监控
func Run(monitorName string, interval time.Duration, monitorFn func(context *MonitorContext), ctx context.Context) {
	// 不立即运行，则先休眠interval时间
	if interval <= 0 {
		panic("interval参数，必须大于0")
	}

	go func() {
		taskInterval := interval
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(taskInterval):
				taskInterval = runMonitor(monitorName, interval, monitorFn)
			}
		}
	}()
}

// RunNow 运行一个监控
// interval:监控运行的间隔时间
// monitorFn:要运行的监控

func RunNow(monitorName string, interval time.Duration, monitorFn func(context *MonitorContext), ctx context.Context) {
	// 立即执行
	monitorFn(&MonitorContext{
		sw: stopwatch.StartNew(),
	})
	Run(monitorName, interval, monitorFn, ctx)
}

// 运行监控
func runMonitor(taskName string, interval time.Duration, monitorFn func(context *MonitorContext)) (nextInterval time.Duration) {
	// 这里需要提前设置默认的间隔时间。如果发生异常时，不提前设置会=0
	nextInterval = interval
	entryTask := container.Resolve[trace.IManager]().EntryTask(taskName)
	try := exception.Try(func() {
		MonitorContext := &MonitorContext{
			sw: stopwatch.StartNew(),
		}
		monitorFn(MonitorContext)
		flog.ComponentInfof("monitor", "%s，耗时：%s", taskName, MonitorContext.sw.GetMillisecondsText())
		if MonitorContext.nextRunAt.Year() >= 2022 {
			nextInterval = MonitorContext.nextRunAt.Sub(time.Now())
		}
	})
	try.CatchException(func(exp any) {
		entryTask.Error(flog.Errorf("[%s] throw exception：%s", taskName, exp))
	})
	entryTask.End()
	asyncLocal.Release()
	return
}
