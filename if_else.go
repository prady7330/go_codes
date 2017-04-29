package main
import "fmt"
func main(){
	condition := 5>6
	if condition {
		fmt.Println("5 is greater than 6")
	}else{
		fmt.Println("5 is not greater than 6")
	}

	if 3==3 {
		fmt.Println("3 is equal to 3")
	}else {
		fmt.Println("3 is not equal to 3")
	}
}