package helpers

import (
	"io"
	"net/http"
	"os"
)

func HttpDownloadFile(url string, targetFile string) {
	out, _ := os.Create(targetFile)
	defer out.Close()
	resp, err := http.Get(url)

	if err != nil {
		panic("AAAAAHHHHH!") // TODO: remove
	}

	defer resp.Body.Close()
	io.Copy(out, resp.Body)
}
