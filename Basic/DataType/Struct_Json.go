package main
/*
 * JSON 
   Marshal/MarshalIndent(struct)
   Unmarshal (data, &struct)   
 */

import ("fmt";"encoding/json";"log";"os")

/* Comment : 
 * you will get syntax error if you use quote(') instead of ` 
 * `json:"released"' are called 'field tags' 
         1) which is often tosed to specifiy an idiomatic JSON name like total_count for a GO field named TotalCount
		 2) regarding, "color,omitempty", which indicates that no JSON output should be produced if the filed has the zero value for its type (false, here) or is otherwize empty
		    Sure enough, the JSON output for Casablanca will have no color field
*/ 
type Movie struct {
	Title 	string
	Year 	int 	`json:"released"`         //'json:"released"'   //Note: not quote(') but `
	Color 	bool 	`json:"color,omitempty"`  //'json:"color,omitempty"'
	Actors	[]string
}

type MovieList []Movie

var movies MovieList

func init(){
	movies = []Movie{
		{Title: "Casablanca", Year:1942, Color:false, 
		Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year:1967, Color:true, 
		Actors:[]string{"Paul Newman"}},
		{Title: "Bullitt", Year:1968, Color:true, 
		Actors:[]string{"Steve McQueen","Jacqueline Bisset"}},
		}
}

func checkerr(err error){
	if err != nil {
			log.Fatalf("JSON marshaling failed %s", err)
			os.Exit(1)
	}	
}

func marshalJson(movies MovieList) []byte {
	data, err := json.Marshal(movies)
	checkerr(err)
	
	fmt.Printf(">>  %s\n%s\n","with Marshal",data)
	
	data, err = json.MarshalIndent(movies, "", "	")
	checkerr(err)
	
	fmt.Printf(">>  %s\n%s\n","with MarshalIndent",data)
	return data
}

func unMarshalJson(data []byte, titles *[]struct{Title string}){
	fmt.Println("\n Unmarshal-------------------\n")
	err := json.Unmarshal(data, titles)	
	checkerr(err)
	fmt.Println(titles)
}
	
	
/* 
 * MAIN function 
 */ 
func main(){
	var data []byte
	data = marshalJson(movies)
	
	var titles []struct {Title string}
	unMarshalJson(data, &titles)	
}