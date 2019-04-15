package exe

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fmt.Printf("The new Url is %s\n", url)
		}
		resp, err := http.Get(url)

		fmt.Printf("The resp code is %s\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(-1)
		}
		c, err := io.Copy(os.Stdout, resp.Body)
		//c, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(-1)
		}
		//n := bytes.IndexByte(c, 0)
		fmt.Printf("%s", string(c))
	}
}
