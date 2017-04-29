package main
import "fmt"
func main() {
	var1,var2 := 1,2
	fmt.Println("before modification")
	fmt.Println(var1)
	//fmt.Println("\n")
	fmt.Println(var2)
	var1,var2 = 3,4
	fmt.Println("after modification")
	fmt.Println(var1)
	fmt.Println(var2)
	//fmt.Println("\n")
}