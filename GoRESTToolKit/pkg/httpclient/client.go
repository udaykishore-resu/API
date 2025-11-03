package httpclient

import (
	"myapp/internal/util"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func GetWithRetry(url string) (*http.Response, error) {
	var resp *http.Response
	err := util.Retry(3, 1*time.Second, func() error {
		var err error
		resp, err = httpClient.Get(url)
		if err != nil || resp.StatusCode >= 500 {
			return err
		}
		return nil
	})
	return resp, err
}
