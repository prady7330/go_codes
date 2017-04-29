package main
import "fmt"
func main(){
	var i int
	for i=0 ; i<=10 ; i++ {
		fmt.Println(i)
	}
	name := "pradeep"
	for char := range name {
		fmt.Print(string(char)+" ")
	}
	i=10;
	for i>0{
		fmt.Println(i)
		i--;
	}
}