package main
import "fmt"

func fizzBuzz(n int32) {
    for i := 1; i <= int(n); i++ {
        divBy3 := (i % 3 == 0)
        divBy5 := (i % 5 == 0)

        switch {
        case divBy3 && divBy5:
            fmt.Println("FizzBuzz")
        case divBy5:
            fmt.Println("Buzz")
        case divBy3:
            fmt.Println("Fizz")
        default:
            fmt.Println(i)
        }
    }
}

func fizzBuzz2(n int32) {
    // Write your code here
    for i := 1; i <= int(n); i++ {     
        switch {
            case i%3 == 0 && i%5 == 0:
                fmt.Println("FizzBuzz")
            case i%5 == 0:
                fmt.Println("Buzz")
            case i%3 == 0:
                fmt.Println("Fizz")
            default:
                fmt.Println(i)
        }
        
    }

}