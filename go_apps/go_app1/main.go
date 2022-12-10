package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 書き込み用
	fmt.Println("here is app1")
	for i := 0; i < 1000; i++ {
		f, err := os.OpenFile(fmt.Sprintf("./shareTo/pseudo_%d.txt", i), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if fi, err := f.Stat(); err == nil {
			fmt.Printf("name: %s\n", fi.Name())
			fmt.Printf("mode: %s\n", fi.Mode())
		}
		content := time.Now().Format("2006_01_02_150405_.000000000")
		data := []byte(content)
		count, err := f.Write(data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("write %d bytes\n", count)
	}
}
