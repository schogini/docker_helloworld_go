package main
import "fmt"
import "os"
import "strconv"
func main() {
    fmt.Println("Hello World. I am Go Lang")

    if (len(os.Getenv("a")) > 0) && (len(os.Getenv("b")) > 0) {
    	firstnum, _ := strconv.ParseInt(os.Getenv("a"), 10, 0)
	    secondnum, _ := strconv.ParseInt(os.Getenv("b"), 10, 0)
	    sum := firstnum + secondnum;
	    output := fmt.Sprintf("Sum of %d & %d is %d", firstnum, secondnum, sum)
    	fmt.Println(output)
    } else {
    	fmt.Println("No parameters passed to calculate the sum.")
    }
    fmt.Println("Go Lang Bye!")
}
