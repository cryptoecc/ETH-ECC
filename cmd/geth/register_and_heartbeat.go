
package main

// import "fmt"
import ( "time"
	"fmt"
	"net/http"
//	"io/ioutil"
)
// import "github.com/go-resty/resty/v2"

// func main() {		go heartBeat()  }//    time.Sleep(time.Second * 5)
func send_hearteat() {
	fmt.Printf("____________________hello world____________________\n"  )
	_, err := 	http.Get( "http://3.39.197.118:34815/heartbeats" )
	if err != nil {
		panic(err)
	}
//	fmt.Printf("%s\n", ioutil.ReadAll(resp))

	// clienttmp := resty.New() //			resp, err := client.R().
	// clienttmp.R().
	// EnableTrace().
	// Get ( "http://3.39.197.118:34815/heartbeats" )
//				 http://3.39.197.118:34815/heartbeats
	for range time.Tick( time.Second * 10 * 1 ) { // every 10 seconds
//	for range time.Tick(time.Second * 60 * 10 ) { // every 10 minutes
		http.Get( "http://3.39.197.118:34815/heartbeats" )
			// client := resty.New() //			resp, err := client.R().
			// client.R().
			// EnableTrace().
			// Get ( "http://3.39.197.118:34815/heartbeats" )
//		Get ( "http://43.200.163.88:34815/heartbeats" )  
	//			( "https://httpbin.org/get" ) //        fmt.Println("Foo")
    }
}
