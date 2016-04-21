package main
/*
 * func literals let us define a function at its point of use
 */
import ("fmt"
	"time"
	)

func main(){
	runtime := time.Minute * 1
	endtime := time.Now().Add(runtime)
	f := squares()
	for i:=0; time.Now().Before(endtime); i++ {
		fmt.Printf("[%d] square : %#v\n", i, f())
		time.Sleep(time.Second << uint(i))
	}
}

func squares() func() int {
	var x int
	fmt.Println("outer : ", x)  /* it will show : outer function will be executed only first time */
	return func() int {
		fmt.Println("inner x: ", x) /* it will show: every call, inner function is called and local variable x is kept */
		x++
		return (x * x)
	}
}
