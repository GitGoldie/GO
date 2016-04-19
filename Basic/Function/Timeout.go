package main

import ("fmt"
	"net/http"
	"time"
	"os"
)

func main(){
	if err := waitForResponse("http://www.abc-test.com"); err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v", err)
		os.Exit(99)
	}
}


func waitForResponse(url string) error {
	const timeout = time.Minute * 1  //in minute unit
	timeend := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(timeend); tries++ {
		_, err := http.Get(url)
		if (err == nil) { return nil }
		fmt.Printf("[%d] server %s doesn't response %v\n", tries, url, time.Now())
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to response after %s",url,timeout)
}



