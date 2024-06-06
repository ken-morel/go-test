package main
import ("fmt")


func factorial(n int) int {
  if n == 1 {
    return 1
  } else {
    return n * factorial(n - 1)
  }
}
func main() {
  name := "colorama"
  defer fmt.Println(name)
  defer fmt.Println(factorial(5))
}
