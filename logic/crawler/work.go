package crawler

import (
	"fmt"
	"runtime/pprof"
	"sync"
)

//定义接口 实现此接口进行实际的爬虫操作
//管理goroutine池来完成工作
type Worker interface {
	work()
}
//定义goroutine池 管理任何已经提交的work工作
type Pool struct {
	worker chan  Worker
	waitgroup sync.WaitGroup
}

//创建一个新的携程管理池
func NewPool(maxGoroutines int) *Pool {

	p := &Pool{worker : make(chan Worker)}

	p.waitgroup.Add(maxGoroutines)
	for i:=0; i< maxGoroutines; i++{
		go func() {
			fmt.Println("携程1开始执行:", i)
			for w := range p.worker{
				w.work()
				fmt.Println("实际工作开始",i)
			}
		}()
	}
	return  p
}

func (p *Pool)Run(worker Worker)  {
	p.worker <- worker
}

func (p *Pool)Shutdown()  {
	close(p.worker)
	p.waitgroup.Wait()
}