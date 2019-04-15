package exe

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		fmt.Println(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(filename)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	CloseFile(f, err)
	return local, n, err
}

func CloseFile(file *os.File, err error) (retErr error) {
	defer func() {
		if closeErr := file.Close(); err == nil {
			retErr = closeErr
		} else {
			retErr = err
		}
	}()
	return retErr
}
