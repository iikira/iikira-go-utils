package taskframework

import (
	"fmt"
	"testing"
	"time"
)

type (
	TestUnit struct {
		retry    bool
		taskInfo *TaskInfo
	}
)

func (tu *TestUnit) SetTaskInfo(taskInfo *TaskInfo) {
	tu.taskInfo = taskInfo
}

func (tu *TestUnit) OnFailed(lastRunResult *TaskUnitRunResult) {
	fmt.Printf("[%s] error: %s, failed\n", tu.taskInfo.Id(), lastRunResult.Err)
}

func (tu *TestUnit) OnSuccess(lastRunResult *TaskUnitRunResult) {
	fmt.Printf("[%s] success\n", tu.taskInfo.Id())
}

func (tu *TestUnit) OnComplete(lastRunResult *TaskUnitRunResult) {
	fmt.Printf("[%s] complete\n", tu.taskInfo.Id())
}

func (tu *TestUnit) Run() (result *TaskUnitRunResult) {
	fmt.Printf("[%s] running...\n", tu.taskInfo.Id())
	return &TaskUnitRunResult{
		//Succeed:   true,
		NeedRetry: true,
	}
}

func (tu *TestUnit) OnRetry(lastRunResult *TaskUnitRunResult) {
	fmt.Printf("[%s] prepare retry, times [%d/%d]...\n", tu.taskInfo.Id(), tu.taskInfo.Retry(), tu.taskInfo.MaxRetry())
}

func (tu *TestUnit) RetryWait() time.Duration {
	return 1 * time.Second
}

func TestTaskExecutor(t *testing.T) {
	te := NewTaskExecutor()
	te.SetParallel(2)
	for i := 0; i < 3; i++ {
		tu := TestUnit{
			retry: false,
		}
		te.Append(&tu, 2)
	}
	te.Execute()
}
