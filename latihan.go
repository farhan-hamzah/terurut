package main
import "fmt"
func balik(x int, y int){
    for x != y{
        fmt.Print(y, x)
        fmt.Scan(&x, &y)
    }
}
func main(){
    var a, b int
    fmt.Scan(&a, &b)
    balik(a, b)
}