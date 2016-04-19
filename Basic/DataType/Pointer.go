package main
/*
 * POINTER
   >> *(asterisk) is represented of pointer followed by type
   >> *(asterisk) is also used to “dereference” pointer variables.
 * the & operator to find the address of a variable
 */
import ("fmt";"strconv")

type NUMBER struct {
	value int
	name string
}

//type numberlist []NUMBER

func copyReverse (from *NUMBER, to *NUMBER){		
	*to = *from	
}

func incrementValue (n *NUMBER){
	//dereference pointer variables
	n.value++
}

func printBeforeAndAfter (n *NUMBER){
	//Note: GO doesn't have dereference -> in C but all .
	fmt.Printf("%-10s = %-2d --> ",n.name, n.value)
	incrementValue(n)
	fmt.Printf("%s(++) = %d\n",n.name, n.value)
}

func main(){

	//type1
	var num1 NUMBER = NUMBER{10,"number1"}
	printBeforeAndAfter(&num1)
	
	//type 2 : new()
	num2 := new(NUMBER)
	num2.name = "number2"
	//Note: num2 is already in pointer
	printBeforeAndAfter(num2)
	
	//type 3 : make
	numArray := make([]NUMBER, 5)
	//numArray[0] = NUMBER{100,"Array[0]"}
	//numArray[1] = NUMBER{200,"Array[1]"}	
	
	fmt.Println("\n>>Incremented value in array using pointer -------- ")
	for idx, a := range numArray {
		a.name = "Array["+strconv.Itoa(idx)+"]"		
		a.value = 100 + (idx * 100)
		printBeforeAndAfter(&a)
		//Note a is local and have to store to global(?) variable
		numArray[idx] = a
	}
		
	fmt.Println("\n>>Copied the array in reverse -------- ")
	emptyArray:= make([]NUMBER, 5)
	
	datacopy := func(from *NUMBER, to *NUMBER){
		*to = *from		
	}
		
	for idx, a := range numArray{
		j := len(numArray)- idx - 1	
		datacopy(&a,&emptyArray[j])	
		emptyArray[j].name = "Copied["+strconv.Itoa(j)+"]"		
	}
	
	for _, a := range emptyArray {		
		printBeforeAndAfter(&a)
	}
}