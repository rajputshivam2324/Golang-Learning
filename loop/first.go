package main
import "fmt"
func main(){
	i:=1
	for i<=3 {
		fmt.Println(i)
		i=i+1
	}
	for j:=1; j<=10 ;j++ {
		fmt.Println(i)
	}
	for i :=range 3{
		fmt.Println( "Print Range :",i)
	}
	for {
		fmt.Println("Loop")
		break
	}
	for i:= range 10{
		if i%2==0{
			continue
		}
		fmt.Println(i)
	}
}