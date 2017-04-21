package main
import "fmt"
func sum(a []int, c chan int, x int) {
	sum := 0
	for _, v := range a {
		sum += v 
	}
	c<-sum //sendsumtoc 
	fmt.Println(x);
}
func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c, 11)
	go sum(a[len(a)/2:], c, 22)
//	x := <-c
//	fmt.Println("in2");
//	y := <-c
//	fmt.Println("in3");
//	fmt.Println(x);
//	fmt.Println(y);
	x, y := <-c, <-c // receive from c
	//fmt.Println("in2");
	fmt.Println(x, y, x + y)
}

