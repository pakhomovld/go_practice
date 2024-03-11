package main


import "fmt"

var num float32

func main() {
    fmt.Println("input float number: ")
    fmt.Scan(&num)
    fmt.Println(int64(num))
}
