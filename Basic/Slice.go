package main
/*
 * Slice Basic (creation, reverse, copy, comparison, append/extend)
   1. append(slice, value)
   2. Insert(slice, idx,value)
   3. copy(sliceTo, sliceFrom)  //note, length of copy is shortest len(either)
 
 * Note
   1. idx, value <= range <slice>
   2. to enable flag in this example:   go run <filename> -debugEn=true
   
 * Reference: https://blog.golang.org/slices 
 */

import ("fmt";"unsafe";"math/rand";"flag")

func reverse(s []string){	
	for i,j:=0,len(s)-1; i<j; i,j=i+1,j-1 {
		s[i],s[j] = s[j],s[i]
	}
}

func appendInt(arr []int, val int) []int {
	debugEn := flag.Bool("debugEn",false,"Enable debug print")
	flag.Parse()
	//check capacity
	//if memory allocated, reslice and update
	//if memory NOT allocated, double size and update
	var temp []int
	if (*debugEn == true){
		fmt.Println(arr,"len(x)",len(arr),"cap(x)",cap(arr))	
	}
	
	if (len(arr) < cap(arr)){
		//reslice (+1)
		temp = arr[:len(arr)+1]
	} else {
		//make new one with one more element
		temp = make([]int, len(arr)+1, len(arr)*2)
		copy(temp, arr)
	}

	//update
	temp[len(arr)] = val	
	return temp
}


func appendIntMult(arr []int, values ... int) []int {
	curLen := len(arr)
	newLen := curLen + len(values)
	curCap := cap(arr)
	
	var temp []int
	
	if (newLen <= curCap){
		temp = arr[:newLen]
	} else {
		temp = make([]int, newLen)
		copy(temp,arr)
	}
	
	//add
	copy(temp[curLen:],values)
	return temp
}

func removeArray(strings []string, option string, idx int) []string {
	
	i:=len(strings)
	switch (option) {	
	case "empty":
		i = 0
		for _,s := range strings{
			if s!=""{
				strings[i] = s
				i++
			}		
		}
	case "index":
		copy(strings[idx:],strings[idx+1:])
		i = len(strings)-1
	}
	
	return strings[:i]
}


func main(){	
	
	c0 := 'c'
	s0 := "c"
	s1 := "fooStr"
	fmt.Println(">>------ sizeof integer vs. strings")
	fmt.Printf("sizeof(%T):%d  sizeof(%T):%d  sizeof(%T):%d\n",c0,unsafe.Sizeof(c0), s0,unsafe.Sizeof(s0), s1, unsafe.Sizeof(s1))	
		
	// Note : this generate 13 element insstead of 12
	months := [...]string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"Jun",7:"Jul",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}
	//reverse array
	reverse(months[:]) //(error) reversa(months) 	//it will generate error "type [13]string []string"	
	//slices
	spring := months[3:6]
	fmt.Println(">>------ slice and comparison")
	Q1 := months[1:4]	
	for _, s := range spring {
		for _, q := range Q1 {
			if s == q {
				fmt.Printf("%s belongs to Spring and Q1\n",s)
			}
		}
	}
	
	//append considering capacity
	array := make([]int, 1, 3)  //syntax(type,len,cap) and cap is ommitted ==> cap <= len
	//array = []int{1,2,3} //it overwrite cap 
	array[0] = 1
	fmt.Println(">>------ append/extend(reslice): ",array,"len:",len(array),",cap:",cap(array))
	array = appendInt(array,rand.Intn(100))
	fmt.Println("[Custom:Append] Array=",array,",len: ",len(array), "cap: ",cap(array))
	array = appendIntMult(array,rand.Intn(100),rand.Intn(100),rand.Intn(100))
	fmt.Println("[Custom:Append] Array=",array,",len: ",len(array), "cap: ",cap(array))
	
	
	//built-int append
	for i:=0; i<7; i++ {
		array = append(array, rand.Intn(1000))
	}
	fmt.Println("[Built-in:Append] Array=",array,",len: ",len(array), "cap: ",cap(array))
	
	array = append(array, 1,2,3,4,5,6,7)
	fmt.Println("[Built-in:Append] Array=",array,",len: ",len(array), "cap: ",cap(array))
	
	//string array and remove empty element & shift 
	strArray := []string{"cat","","dog","","bluejay"}
	strArray = removeArray(strArray,"empty",0)
	fmt.Println("[Custom:Remove] strArray=",strArray,",len: ",len(strArray), "cap: ",cap(strArray))
	strArray = removeArray(strArray,"index",1)
	fmt.Println("[Custom:Remove] strArray=",strArray,",len: ",len(strArray), "cap: ",cap(strArray))
	
	slice1 := []int{1,2,3}
	slice2 := []int{11,22,33} 
	slice12 := append(slice1,slice2...)
	slice3 := append([]int(nil), slice12...)  //Note : ... (must)
	fmt.Println(slice12,"\n",slice3)
}