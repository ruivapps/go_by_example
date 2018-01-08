/* golang comments follow c style comments
you can write multi line comment or single line comment

golang enforce all variables, unused variables will show as error during compiling time.
you can use special _ to handle things you don't want.
unused  _ will not trigger compiler error

http package documentation: https://golang.org/pkg/net/http/
error: https://golang.org/pkg/builtin/#error
http Response: https://golang.org/pkg/net/http/#Response
*/

package main

/* golang import similar as c #include, you can import more
than one in single import call, or separate
	import "fmt"
	import "io/ioutil"
*/
import (
	"fmt"       // this is for I/O (print)
	"io/ioutil" // this is for read HTTP response
	"net/http"  // this is for http requests
	"os"        // this is for os.Args
	"time"      // need to work with time
)

func HttpGet(url string) []byte {
	// most golang function return something and error
	resp, err := http.Get(url)
	if err != nil {
		/*
			on the function, we define return type is byte
			error is interface
				type error interface {
					Error() string
				}
			so we will need to convert error.Error() into byte
		*/
		return []byte(err.Error())
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(err.Error())
	}
	resp.Body.Close()
	return bytes
}

// golang start with function main()
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: http_go url\n")
		os.Exit(1)
	}
	begin := time.Now()
	response := HttpGet(os.Args[1])
	fmt.Println(string(response))
	fmt.Println("time spent:", time.Since(begin).Seconds())
}
