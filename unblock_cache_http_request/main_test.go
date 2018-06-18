package main

import (
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func testURL() []string {
	return []string{
		"http://www.weibo.com",
		"http://www.weibo.com",
		"http://www.weibo.com",
		"http://www.baidu.com",
		"http://www.baidu.com",
		"http://www.163.com",
		"http://www.sina.com.cn",
		"http://www.hao123.com",
		"http://www.12306.cn",
	}
}

func httpGetFunc(url string) (*Result, error) {
	res := &Result{}

	resp, err := http.Get(url)
	res.Err = err
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode

	if err != nil {
		return res, res.Err
	}
	defer resp.Body.Close()

	res.Body, res.Err = ioutil.ReadAll(resp.Body)
	return res, res.Err
}

func TestRequestAll(t *testing.T) {
	cache := NewCache()

	var wg sync.WaitGroup
	wg.Add(len(testURL()))

	begin := time.Now()
	for _, url := range testURL() {
		url := url
		go func() {
			res, err := cache.Get(url, httpGetFunc)
			if err != nil {
				t.Logf("request cache error: %v", err)
			}
			t.Logf("request result: %s,%+v\n", url, res.Status)
			wg.Done()
		}()
	}
	wg.Wait()

	t.Logf("used_time: %v\n", time.Since(begin))
}
