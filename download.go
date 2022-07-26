package main

import (
	"fmt"
	"os/exec"
	"sync"
)

const (
	URL  = "https://pages.cs.wisc.edu/~remzi/OSTEP/Chinese/%d.pdf"
	URL1 = "https://pages.cs.wisc.edu/~remzi/OSTEP/Chinese/0%d.pdf"
	URL2 = "https://pages.cs.wisc.edu/~remzi/OSTEP/Chinese/%s.pdf"
)

var (
	nonTypeURL = []string{
		"fla",
		"flb",
		"flc",
		"fld",
		"fle",
		"flf",
	}
)

func main() {
	var url string
	var wg sync.WaitGroup
	for i := 0; i < 51; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer fmt.Println("第", i, "个文件下载完毕")
			if i < 10 {
				url = fmt.Sprintf(URL1, i)
			} else {
				url = fmt.Sprintf(URL, i)
			}
			cmd := exec.Command("wget", url)
			e := cmd.Run()
			if e != nil {
				fmt.Println(e)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("开始下载附录")
	for k, v := range nonTypeURL {
		nonTypeURL[k] = fmt.Sprintf(URL2, v)
	}
	nonTypeURL = append(nonTypeURL, fmt.Sprintf(URL2, "preface"), fmt.Sprintf(URL2, "toc"))
	downNonTypeUrl(nonTypeURL)

}

func downNonTypeUrl(urls []string) {
	for _, url := range urls {
		cmd := exec.Command("wget", url)
		cmd.Run()
	}
}
