package main

import (
    // Local imports
    "os"
    "fmt"
    "sync"
    "time"
    "strings"
    "io/ioutil"
    "math/rand"

    // Extern imports
    "github.com/sheerun/queue"
)

type threadInfo struct {
    queueList queue.Queue
    queueMaxSize int
}

var (
    runningThreads int = 10
    workersPerThread int = 100

    threadGroup sync.WaitGroup
    taskGroup sync.WaitGroup

    threadList = make(map[int]*threadInfo)
)

func spawnThread(threadID int) {
    var emptyCount int = 0
    threadList[threadID] = &threadInfo{*queue.New(), workersPerThread}

    fmt.Printf("Thread[%d]: Thread started\r\n", threadID)

    for {
        if (threadList[threadID].queueList.Length() <= 0) {
            if (emptyCount >= 5) {
                break
            }

            time.Sleep(1 * time.Second)
            emptyCount++
            continue
        }

        // Run task
        fmt.Printf("%v\r\n", threadList[threadID].queueList.Pop())
    }

    threadGroup.Done()
}

func queueTask(addr string) {

    for {
        selectedID := rand.Intn(runningThreads)

        if threadList[selectedID].queueList.Length() == workersPerThread {
            continue
        }

        threadList[selectedID].queueList.Append(addr)
        break
    }

    taskGroup.Done()
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("./" + os.Args[0] + ": file_name.txt")
        return
    }

    fileBuf, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("./" + os.Args[0] + ": File not found")
        return
    }

    fileSplit := strings.Split(string(fileBuf), "\n")
    if len(fileSplit) <= 0 {
        fmt.Println("./" + os.Args[0] + ": Empty file")
        return
    }

    for i := 0; i < runningThreads; i++ {
        threadGroup.Add(1)
        go spawnThread(i)
    }

    time.Sleep(1 * time.Second)

    for i := 0; i < len(fileSplit); i++ {
        taskGroup.Add(1)
        go queueTask(fileSplit[i])
    }

    threadGroup.Wait()
    taskGroup.Wait()
}
