## Go并发
> Do not communicate by sharing memory; instead, share memory by communicating.

#### 并发场景
* `Channel`: 用来在不同线程之间传递数据的载体
* `WaitGroup` : 一个大任务拆成多个子任务，需要等到所有任务都完成，才停止当前大任务。
* `Context`: 给定**任务**在其他线程执行，主动通知该线程让其终止
    * 如果任务仅有少量几个，我们可以通过 `channel + select` 的方式实现
    * 如果任务是多个或者任务中也有子任务呢？我们可以通过 `context` 携带信号量的方式实现