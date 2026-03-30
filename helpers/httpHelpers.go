package helpers

import (
	"io"
	"net/http"
	"os"
)

func HttpDownloadFile(url string, targetFile string) error {
	out, _ := os.Create(targetFile)
	defer out.Close()
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)

	return err
}
