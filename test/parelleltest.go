package main 

import ("fmt"
        "runtime"
        "sync")
        
func printGibberish(wg *sync.WaitGroup, stuff string){
    
    defer wg.Done()
    for i := 0; i< 10000; i++{
        fmt.Println(stuff)
    }
}

func main(){
    runtime.GOMAXPROCS(8)
    var wg sync.WaitGroup
    wg.Add(2)
    go printGibberish(&wg, "Sirius")
    go printGibberish(&wg, "Remus")
    wg.Wait()
    fmt.Println("Done")
}