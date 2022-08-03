package main

import (
    "fmt"
    "time"
)

func TestBaseChanel() {
    var ca = make(chan int)
    defer close(ca)
    go func() {
        // time.Sleep(time.Second * 2)
        ca <- 3 + 4
    }()
    fmt.Println("---------")
    num := <-ca
    fmt.Println(num)
}

func TestPanicChanel() {
    var ca = make(chan int)
    defer close(ca)
    go func() {
        close(ca)
        ca <- 3 + 4
    }()
    fmt.Println("---------")
    num := <-ca
    fmt.Println(num)
}

func TestCloseChanel() {
    var ca = make(chan int)
    defer close(ca)
    go func() {
        ca <- 3 + 4
    }()
    close(ca)
    fmt.Println("---------")
    num := <-ca
    fmt.Println(num)
}

func TestnilChanel() {
    var ca = make(chan int)
    defer close(ca)
    go func() {
        ca = nil
    }()
    fmt.Println("---------")
    num := <-ca
    fmt.Println(num)
}

func TestWaitChanel() {
    var ca = make(chan int)
    defer close(ca)
    go func() {
        time.Sleep(time.Second * 3)
        ca <- 3 + 4
    }()
    go func() {
        fmt.Println("---------")
        time.Sleep(time.Second * 3)

        num := <-ca
        fmt.Println(num)
    }()

}

func TestLenChanel() {
    var ca = make(chan int)
    //var ca = make(chan int, 2)

    defer close(ca)
    go func() {
        for i := 1; i < 6; i++ {
            fmt.Println(len(ca), "-----x------", i)
            ca <- i
        }
    }()
    for i := 1; i < 6; i++ {
        time.Sleep(time.Second * 2)
        fmt.Println(len(ca), "-----y-----", <-ca)
    }
}

func TestForChanel() {
    var ca = make(chan int)
    //var ca = make(chan int, 2)

    go func() {
        for i := 1; i < 6; i++ {
            fmt.Println(len(ca), "-----x------", i)
            ca <- i
        }
        close(ca)
    }()
    for i := range ca {
        fmt.Println(i, "----", len(ca), "---", <-ca)
    }
}

func TestSelectChanel() {
    var ca = make(chan int)

    go func() {
        time.Sleep(time.Second * 1)

        ca <- 1
        time.Sleep(time.Second * 2)
        ca <- 2
        time.Sleep(time.Second * 2)
        close(ca)
    }()
    idx := 0
    for {
        if idx >= 20 {
            break
        }
        idx++
        select {
        case x, ok := <-ca:
            time.Sleep(time.Second)

            if ok {
                fmt.Println("received ca ", x, " from ")
            } else {
                fmt.Println("ca is closed")
            }
        case i := <-ca:
            fmt.Println("send ca-----", i)

        default:
            fmt.Println("no communication")
        }
        time.Sleep(time.Second)

    }

}

func main1() {
    // TestBaseChanel()
    // TestPanicChanel()
    // TestWaitChanel()
    // TestCloseChanel()
    // TestLenChanel()
    //TestForChanel()
    // TestSelectChanel()
}
