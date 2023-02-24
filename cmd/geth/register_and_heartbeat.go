
package main

// import "fmt"
import "time"
import "github.com/go-resty/resty/v2"

// func main() {		go heartBeat()  }//    time.Sleep(time.Second * 5)
func send_hearteat() {
    for range time.Tick(time.Second * 60 * 10 ) { // every 10 minutes
			client := resty.New()
//			resp, err := client.R().
			client.R().
			EnableTrace().
			Get ( "http://3.39.197.118/heartbeats" )  			
//		Get ( "http://43.200.163.88:34815/heartbeats" )  
	//			( "https://httpbin.org/get" ) //        fmt.Println("Foo")
    }
}
