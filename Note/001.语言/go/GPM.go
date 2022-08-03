package main

import (
    "fmt"
    "os"
    "runtime"
    "runtime/trace"
    "sync"
    "time"
)

func TestGPM() {

    //创建trace文件
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    //启动trace goroutine
    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()
    wg := sync.WaitGroup{}
    for i := 0; i <= 20; i++ {
        wg.Add(1)
        go func(idx int) {
            defer wg.Done()
            time.Sleep(time.Second)
            fmt.Println(idx, "process num")

        }(i)

    }
    wg.Wait()
    //main
}

func TestGPM1() {

    for i := 0; i <= 10; i++ {
        time.Sleep(time.Second)
        fmt.Println(i, "process num")

    }
    //main
}

func TestParams() {
    fmt.Println("process num", runtime.GOMAXPROCS(-1))
    fmt.Println("process cpu", runtime.NumCPU())
    fmt.Println("process goroutine", runtime.NumGoroutine())

}
func getGOMAXPROCS() int {
    runtime.NumCPU()             // 获取机器的CPU核心数
    return runtime.GOMAXPROCS(0) // 参数为零时用于获取给GOMAXPROCS设置的值
}

func GOMAXPROCS() {
    fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())
}
func main() {
    TestGPM()
}
