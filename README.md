
# go-utils

go utils from everywhere in one place

# 部分 utils 的使用说明

## assert

一个简洁的 `if err != nil {panic(err)}` 的替代函数

```go
// 原始写法:
if err = connectDb(xx); err != nil{
    panic(fmt.Sprintf("connect db error, %s", err))
}

// 使用 assert
Assert(connectDb(xx), "connect db error")
```

## async.Runner
并发运行多个函数, 便于管理所有函数的生命周期，类似 oklog/run, 其广泛用于各种开源项目的启动函数中.
使用方式见下面的两个例子，支持使用 Stop 控制或者 配置外部的 Stop Chan 进行控制。

如果需要在任务启动时接收系统 signal，可以配合 interrupt 包使用。

```go
func TestExampleRunnerWithOuterStop(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	r := NewRunner().WithStopChan(ctx.Done())

	// add function
	r.Add(func(stop <-chan struct{}) {
		fmt.Println("func1 started")
		<-stop
		fmt.Println("func1 closed")
	})
	// add more function
	r.Add(func(stop <-chan struct{}) {
		fmt.Println("func2 started")
		<-stop
		fmt.Println("func2 closed")
	})
	r.Start()
	<-ctx.Done()
	time.Sleep(time.Millisecond * 10)
}

func TestExampleRunnerWithStopMethod(t *testing.T) {
	r := NewRunner(
		func(stop <-chan struct{}) {
			fmt.Println("func1 started")
			<-stop
			fmt.Println("func1 closed")
		}, func(stop <-chan struct{}) {
			fmt.Println("func2 started")
			<-stop
			fmt.Println("func2 closed")
		})
	r.Start()
	time.Sleep(time.Second)
	r.Stop()
	time.Sleep(time.Millisecond * 10)
}

```

## buffer
buffer 包提供了两种简单的 queue，一种 fifo（ring buffer）, 另一种 lifo (queue)，都支持自动扩容的模式。其中 queue 为线程安全。 

## cache

提供了 基于 lrucache 封装的带 ttl 特性的 cache

```go
cache := NewExpiring()

// when empty, record is stored
cache.Set("foo", record1, time.Hour)
```

## tcache

和 cache 的不同之处在于:

1. ttl 支持 soft ttl（返回数据，触发 fetch）和 hard ttl（先fetch，再返回数据）
2. 支持预先设置 fetch 函数，简化使用

```go
accountCache = cache.NewCache(60, 300, 1000, func(key string)(value interface{}, err error){
    account, err := fetchUserAccountInfoSomewhere(key)
    return account, err
} )

accountCache.Get(key)
```

## container

目的在于解决项目中的多依赖，全局变量问题。实现了自动的依赖注入，也可以手动 provide （推荐）
使用 container 能够自动解决依赖次序关系，lazy 新建，管理全局依赖项目，避免参数到处传递带来的混乱.

```go
// 使用 container 之前
clientA := newClientA()
clientB := newClientB()
...
clientZ := newClientZ()

serviceA := NewServiceA(clientA, clientB, clientC, clientD, ......clientZ)


// 使用 container 之后
container.RegisterType(&ClientA{}, func(){return newClientA(), nil})
...

// serviceA 的新建函数
func NewServiceA(){
    serviceA.clientA = container.MustProvide(&ClientA{}).(*ClientA)
}

// 自动依赖注入 使用 tag 'autowired:' 请参考 test
```

## pointer

数值转指针

```go
StringPtr(s string) *string
Float32Ptr(i float32) *float32
...
```

## retry

自动的 retry 函数封装，支持 backoff

```go
opts := wait.Backoff{Factor: 1.0, Steps: 3}

err := RetryOnConflict(opts, func() error {
    return ConflictError
})
```

## sets

类似 python set 类，支持 int，string 等多种 set

```go
a := sets.NewString()
a.Insert("a")

if a.Has("a"){
    fmt.Println("set a has a")
}
```

## wait

用于运行一些定期任务，支持多种循环和启动方式

```go
// 永久执行，每隔 period 运行 f
Forever(f func(), period time.Duration)

// Jitter 运行, 每隔 period 运行 f，直到 stop
JitterUntil(f func(), period time.Duration, jitterFactor float64, sliding bool, stopCh <-chan struct{}) 

// 立刻执行 ConditionFunc，此后执行间隔为 interval 直到 ConditionFunc 运行成功或者 timeout
PollImmediate(interval, timeout time.Duration, condition ConditionFunc)

...
```

## workqueue

内存消息队列，支持 delay，ratelimit 等等类型


### workqueue.parallelize

并发执行，支持提交到 goroutine pool
类似的包 还有 goroutinemap，这个包支持 将 goroutine 命名以防止重复的 goroutine 提交，适用于一组命名的任务提交。

```go
import (
 "github.com/panjf2000/ants/v2"
)

type TestPool struct {
	pool *ants.Pool
}

func NewTestPool(size int) *TestPool {
	p := &TestPool{}
	p.pool, _ = ants.NewPool(size)
	return p
}

func (p *TestPool) Submit(f func()) {
	_ = p.pool.Submit(f)
}


var pool = NewTestPool(10)
var l = 100
ParallelizeWithPoolUntil(context.Background(), l, l, pool, func(piece int) {
    // do something
})
ParallelizeUntil(context.Background(), l, l, func(piece int) {
   // do something
})
```


## 环境变量
获取一个环境变量，如果没有，设置默认值

```go
import "github.com/u2takey/go-utils/env"
a := env.GetEnvAsStringOrFallback("env1", "val")
```

## 常用加密库

[encrypt](./encrypt/encrypt.go)

## Json
json 包的目的是简化替换 json 库到 jsoniter 的流程, 项目中直接搜索 "encoding/json" 替换为 "github.com/u2takey/go-utils/json" 就可以


## testing
ts 库是一个另一个测试库，特点是链式测试，对简单的非数据驱动的测试比较友好，把简单的 case 用一个语句表示，测试代码看起来会比较清楚。

```go
ts.Case("show case usage 1", map[string]interface{}{"A": "b"}).PropEqual("A", "b")
ts.Case("show case usage 2", struct{ A string }{A: "b"}).PropEqual("A", "b")
ts.Case("show case usage 3", TestA{}).PropEqual("FuncB", "b")
```

## retry 
包含了一个简单的 retry 库，和功能更强大的 avast/retry-go
 

## print
print 包的目的是帮助在日志中打印结构体用于 debug 等用途。有些打印结构体使用 `%+v` 就够了，但是当结构体中有指针时，
你会发现打印出来的是指针地址，这个内容没有什么意义。另一种方式是使用 [go-spew](https://github.com/davecgh/go-spew)， 
这个库能够帮助打印出结构中的指针，以及嵌套的指针。更多时候，你可能希望打印出来的是个 json，这样无论可读性，还是从日志中
恢复出原始 case 的能力都很好。但是打印 json 存在一个问题，即：你并不知道这个 json 会有多大，对于大的结构体转换 json 和打印日志
的成本都很好，还会把日志搞得一团糟。print.Json 就是处于这个目的出现的。这个包能帮助你设置省略部分结构体的内容，只打印出
结构体中的那些关键部分。使用它，你也可以很方便的做一个 http/rpc 请求的 log middleware，而不用担心性能问题。

```go
a := struct {
		A string
		B []string
	}{
		A: "abcdf",
		B: []string{"abcdf", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a", "cdf", "", "abcdf", "adf", "abf", "bcdf", "a"},
	}
fmt.Println(print.MarshalStringIgnoreError(a))

// 输出: {"A":"abcdf","B":["abcdf","cdf","","abcdf","adf",<omitted, total length: 15>]}
```