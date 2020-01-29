// Our first program to use variables
import "fmt"

func arithmetic_sum(var a int, var d int, var n int){
    var sum int = 0
    sum = 2 * a + (n - 1) * d
    return sum
}

func main() {
    fmt.Println("Sum of 10 numbers: = ",arithmetic_sum(1, 2, 10))
}