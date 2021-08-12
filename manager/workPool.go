package manager

import "fmt"

// --------------------------- Job ---------------------
type Job interface {
	Do()
}

// --------------------------- Worker ---------------------
type Worker struct {
	JobQueue chan Job
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}
func (w Worker) Run(wq chan chan Job) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			}
		}
	}()
}

// --------------------------- WorkerPool ---------------------
type WorkerPool struct {
	workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}
func (wp *WorkerPool) Run() {
	fmt.Println("初始化worker")
	//初始化worker
	for i := 0; i < wp.workerlen; i++ {
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	// 循环获取可用的worker,往worker中写job
	go func() {
		for {
			select {
			case job := <-wp.JobQueue:
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}

// --------------- 使用 --------------------
/*
type Score struct {
    Num int
}
func (s *Score) Do() {
    fmt.Println("num:", s.Num)
    time.Sleep(1 * 1 * time.Second)
}
func main() {
    num := 100 * 100 * 20
    // debug.SetMaxThreads(num + 1000) //设置最大线程数
    // 注册工作池，传入任务
    // 参数1 worker并发个数
    p := NewWorkerPool(num)
    p.Run()
    datanum := 100 * 100 * 100 * 100
    go func() {
        for i := 1; i <= datanum; i++ {
            sc := &Score{Num: i}
            p.JobQueue <- sc
        }
    }()
    for {
        fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
        time.Sleep(2 * time.Second)
    }
}
*/
