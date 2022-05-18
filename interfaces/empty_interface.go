package main
 
import "fmt"
 
// declaring an empty interface
type emptyInterface interface {
}
 
// declaring a new struct type which has one field of type empty interface
type person struct {
    info interface{}
}
 
func main() {
 
    // declaring an empty interface value
    var empty interface{}
   
    // an empty interface may hold values of any type
    // storing an int in the empty interface
    empty = 5
    fmt.Println(empty) // => 5
 
    // storing a string in the empty interface
    empty = "Go"
    fmt.Println(empty) // => Go
    fmt.Println("Go", 5 , true, false, nil)
    // storing a slice in the empty interface
    empty = []int{2, 34, 4}
    fmt.Println(empty) // => [2 34 4]
 
   // fmt.Println(len(empty)) // -> error, and it doesn't work!
 
    // retrieving the dynamic value using an assertion
    fmt.Println(len(empty.([]int))) // => 3
    // declaring person value
    you := person{}
 
    // assigning any value to empty interface field
    you.info = "You name"
    you.info = 40
    you.info = []float64{4.5, 6., 8.1}
 
    fmt.Println(you.info)
}