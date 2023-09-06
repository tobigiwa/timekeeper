package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetAsciiArt(n int) string {

	if n > 1 {
		os.Exit(1)
	}

	var (
		colorArr  = [7]string{"\u001b[33m", "\u001b[33m", "\u001b[33m", "\u001b[31m", "\u001b[35m", "\u001b[35m", "\u001b[33m"}
		turnoff   = "\u001b[0m\n"
		ch        = make(chan string)
		fileNames = [2]string{"asciiArtLogo.txt", "asciiArt.txt"}
		container = make([]string, 0, 2)
		count     = 0
		size      = len(colorArr)
	)

	for _, v := range fileNames {
		fileName := v

		go func(fileName string, ch chan<- string) {
			var (
				buf   strings.Builder
				bytes []byte
				err   error
			)
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)

			if bytes, err = os.ReadFile(fileName); err != nil {
				ch <- err.Error()
				return
			}

			strSlice := strings.Split(strings.TrimSpace(string(bytes)), "\n")
			for _, art := range strSlice {
				buf.WriteString(fmt.Sprintf("%s %s %s", colorArr[r.Intn(size)], art, turnoff))
			}
			ch <- buf.String()

		}(fileName, ch)
	}

	for v := range ch {
		count++
		container = append(container /*strings.TrimSpace(v)*/, v)
		if count >= 2 {
			close(ch)
			fmt.Println("channel closed successfully")
			break
		}
	}

	return container[n]

}
