package main
/*
 * This file is for : print format 
 * How to generate execution file : go build <filename>
 */
import ("fmt")

func main(){

	exInteger()
	exStringByteRune()
}

/*
 * const and auto type assignment
 */
func exInteger(){	
	const max = 10			
	
	fmt.Println("exInteger =======================================\n")
	fmt.Println("\tSame Variable Can be displayed as different format\n\t%d\t%b\t%x\t%q\t%T (format)")
	fmt.Println("\t-----------------------------------------")
	for i := 100; i < max+100; i++ {
		fmt.Printf("\t%d\t%b\t%x\t%q\t%T\n",i,i,i,i,i)
	}	
}

/*
 * This flag causes the output to escape not only non-printable sequences, 
 * but also any non-ASCII bytes, all while interpreting UTF-8
 */
func exStringByteRune(){
	fmt.Println("exStringByteRune =======================================\n")
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println("\t%x %q %+q (format)\n")
	for i:=0; i<len(sample); i++{
		fmt.Printf("\t%x %q %+q\n",sample[i], sample[i], sample[i])
		/* expected examples: bc '¼' '\u00bc' */
	}
}





