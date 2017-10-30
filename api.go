package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/bakape/r-a-d.io/common"
)

var (
	// Cached data from JSON API. Must lock apiMu.
	apiData common.API

	// Hash of JSON data from API. Must lock apiMu.
	apiHash string

	apiMu sync.RWMutex
)

// Fetches, updates and caches data from the r-a-.d.io JSON API
func init() {
	go func() {
		for {
			var (
				data struct {
					Main common.API
				}
				buf     []byte
				hashBuf [16]byte
				hash    string
			)

			resp, err := http.Get("https://r-a-d.io/api")
			if err != nil {
				goto next
			}
			buf, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				goto next
			}

			hashBuf = md5.Sum(buf)
			hash = base64.URLEncoding.EncodeToString(hashBuf[:])

			err = json.Unmarshal(buf, &data)
			if err != nil {
				goto next
			}

			apiMu.Lock()
			apiHash = hash
			apiData = data.Main
			apiMu.Unlock()

		next:
			if err != nil {
				log.Printf("fetch: %s\n", err)
			}
			if resp != nil {
				resp.Body.Close()
			}
			time.Sleep(time.Second * 10)
		}
	}()
}
