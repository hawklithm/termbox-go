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

	termbox.SetImageCell(10, 1, []byte(imgBase64Str))
	termbox.SetImageCell(14, 1, []byte(imgBase64Str))
	termbox.SetImageCell(18, 1, []byte(imgBase64Str))
	termbox.SetCell(1, 1, 'z', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(2, 1, 'z', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(3, 1, 'z', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(3, 2, 't', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(5, 1, 'z', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, 1, 'z', termbox.ColorWhite, termbox.ColorBlack)
	//termbox.SetImageCell(10, 10, []byte(imgBase64Str))
	termbox.SetCell(10, 10, 'z', termbox.ColorWhite, termbox.ColorBlack)

mainloop:
	for {
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
		}
	}
}
