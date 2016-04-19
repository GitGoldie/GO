package main

/*
 * Struct/Doubly Linked List
   Add
   Sort
 * Note
   1) should start with Capital letter for exported field
   2) a named struct type S cannot declrae a field of the same type S (cannot contain itself)
      but, it can have a field of the pointer type *S (recursive data structure)
 */
import ("fmt")

type tree struct{
	value int
	left, right *tree
}

func main(){
	var root *tree
	values := []int{77,55,88,33,66,99,11}
	
	for _, v := range values {
		fmt.Println("root",root)	
		root = addToList(root, v)		
	}
	
	SortList(root, values[:0])

	displayList(root)
	
}


func addToList(t *tree, val int) *tree{
	
	if t == nil {
		root := new(tree)
		root.value = val
		t = root
		fmt.Println("first",t)
	} else {
		
		if (val > t.value){
			t.right = addToList(t.right, val)
		} else	{
			t.left = addToList(t.left, val)
		}
	}
	
	return t
}

func SortList(t *tree, values []int) []int{
	if t!= nil{
		values = SortList(t.left, values)
		values = append(values, t.value)
		values = SortList(t.right, values)
	}
	
	return values
}

func displayList(t *tree){
	//start from now
	first := t
	for {
		if t.left == nil { break }	
		first = t.left
	}
	
	fmt.Println("MY LIST ----------- \n")
	for {
		if (first.right == nil) { break }
		fmt.Printf("\t%d",t.value)
		first = t.right
	}
}


