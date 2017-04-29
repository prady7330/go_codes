package main
import "fmt"
func main(){
	phrase1:="the quick brown fox"
	phrase2:=" jumps over the lazy dog"
	fmt.Println(phrase1+phrase2)
	slice := phrase1[0:2]
	fmt.Println(slice)
	fmt.Println(phrase2[1:5])
	character := phrase2[0]
	fmt.Println(character)
}