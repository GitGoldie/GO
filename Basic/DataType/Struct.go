package main
/*
 * This package will include 
 * struct creation and intialization/ default value
 * struct extend
 * flag to get command-line input
 */
import ("fmt";"flag")

//Custom Type Definition
type myInt int
type myFloat float64
type myString string 
type KID struct {
	age myInt			//default 0
	height, weight myFloat  //default 0.0
	name myString		//default " "
}

/*
 * package level initialization and global variable
 */
var debugEn bool

func init(){
	debugEn = false	
}

// Main Function
func main(){
	debug := flag.Bool("debug",false,"enable/disable debug message")
	/* parse command-line flag */
	flag.Parse()
	
	/* How to create a object(s) */
	//method 1 : direct declaration
	Boys := []KID {{10, 2.5, 10.3, "Jimmy"},
					{10, 3.5, 9.1, "Teddy"},
					{10, 3.5, 8.2, "Alex"},
					{10, 4.5, 7.3, "Molly"}}   
		
	fmt.Println("\n----------- BOYS----------------")
	for _, kid:= range Boys {
		fmt.Printf("%-8s is %d years old (weight %2.1f(lb), height %2.1f(ft))\n", 
				kid.name, kid.age, kid.weight, kid.height)
	}
	
	//method 2: using "new"
	//Note: it will return a pointer however, GO doesn't need to use -> but everything .
	newKid := new(KID)	
	(*newKid).name = "Sera"
	newKid.weight = 11.10	//all same
	(*newKid).height = 5.5 	//all same	
	
	if (debugEn == true) || (*debug == true) {
		fmt.Printf("\n(DBG)...adding girls...%-8s is %d years old (weight %2.1f(lb), height %2.1f(ft))\n", 
				newKid.name, newKid.age, newKid.weight, newKid.height)
	}

	//method 3: make 
	Girls := make([]KID, 1, 10)  //make(type, assignment, max capacity)
	Girls[0] = *newKid
	//Extend array 
	Girls = Girls[0:2]   //Note : Girls[0:11] will go to PANIC (out of range)
	Girls[1] = KID{10, 3.5, 12, "Amy"}
	
	fmt.Println("\n----------- GIRLS ----------------")
	for _,kid := range Girls {			
	fmt.Printf("%-8s is %d years old (weight %2.1f(lb), height %2.1f(ft))\n", 
				kid.name, kid.age, kid.weight, kid.height)
				}
}

/*
 * go run <filename> -debug=true
 */
