package demo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type Result struct {
	path string
	hash [md5.Size]byte
	err  error
}

func MD5Hash(rootPath string, resultCh chan Result, errCh chan error) {

	wg := &sync.WaitGroup{}

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errCh <- err
		}
		wg.Add(1)
		go func() {
			data, err := ioutil.ReadFile(path)
			resultCh <- Result{path: path, hash: md5.Sum(data), err: err}
			wg.Done()
		}()
		return nil
	})

	go func() {
		wg.Wait()
		close(resultCh)
		close(errCh)
	}()

}

func Md5Demo(root string) {

	resultCh := make(chan Result)
	errCh := make(chan error)

	MD5Hash(root, resultCh, errCh)

	m := make(map[string][md5.Size]byte)

	for res := range resultCh {
		if res.err != nil {
			fmt.Println("error:", res.err)
		}
		m[res.path] = res.hash
	}

	for err := range errCh {
		fmt.Println("err:", err)
	}

	for key, val := range m {
		fmt.Println(key, ":", hex.EncodeToString(val[:]))
	}

}
