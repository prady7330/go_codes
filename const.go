package main
import "fmt"
func main(){
	const number float64 = 1.42
	fmt.Println(number)
	const(
			first = iota + 1 + 2
			second
			third
		)
	fmt.Println(first," ",second," ",third)
}