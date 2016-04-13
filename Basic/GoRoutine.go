package main

import ("fmt"; "os";"net/http"; "io/ioutil";"time"; "strings")

type Latency struct {
		url string
		readtime float64
		sizeByte int
}

func main(){
	webs := []string{"http://www.yahoo.com",
	"www.google.com",
	"www.hgst.com",
	"http://sajenkins01.stec-inc.ad:8080/job/Pre-Commit_Kickoff/"}

	c := make(chan Latency)
	
	for _, web := range webs {
		go fetchurl(web, c)	
	}
	
	for range webs{
		result := <-c
		fmt.Println("Result: ",result)
	}
}


func checkError (e error){
	if e != nil {
		fmt.Fprintf(os.Stderr, "fetch error %v\n", e)
		os.Exit(1)
	}
}


func fetchurl(web string, c chan Latency){
	//check website name
	if strings.Contains(web, "http://") == false {
		web = "http://"+web
	}
	
	start := time.Now()  //millisecond
	
	//greeting
	resp, err := http.Get(web)
	checkError(err)
	//read
	b,err := ioutil.ReadAll(resp.Body)
	checkError(err)
	defer resp.Body.Close()
	//post process
	c <- Latency{web, time.Since(start).Seconds(), len(b)}
}