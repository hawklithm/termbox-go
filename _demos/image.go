package main

import (
	"bufio"
	"encoding/base64"
	"github.com/hawklithm/termbox-go"
	"os"
)

func getImage(src string) string {
	imgFile, _ := os.Open(src)
	defer imgFile.Close()
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	//read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	//convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64str
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	imgBase64Str := getImage("file.jpg")

	termbox.SetImageCell(1, 1, []byte(imgBase64Str))
	termbox.SetImageCell(10, 1, []byte(imgBase64Str))
	termbox.SetImageCell(16, 32, []byte(imgBase64Str))
	for i := 0; i < 36; i++ {
		ch := i
		if ch <= 9 {
			ch = ch + '0'
		} else {
			ch += 'a' - 10
		}
		termbox.SetCell(0, i, rune(ch), termbox.ColorWhite, termbox.ColorBlack)
	}

	for i := 0; i < 36; i++ {
		ch := i
		if ch <= 9 {
			ch = ch + '0'
		} else {
			ch += 'a' - 10
		}
		termbox.SetCell(i, 0, rune(ch), termbox.ColorWhite, termbox.ColorBlack)
	}

mainloop:
	for {
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
		}
		termbox.SetCell(0, 6, 'g', termbox.ColorWhite, termbox.ColorBlack)
		//termbox.SetImageCell(1, 1, []byte(imgBase64Str))
	}
}
