package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var f []fs.DirEntry
	var err error
	if f, err = os.ReadDir("c:\\png"); err == nil {
		fmt.Println(err)
	}
	for _, i := range f {
		if i == nil {
			continue
			//break?
		}
		nn := strconv.Itoa(int(time.Now().Local().UnixMilli()))
		time.Sleep(1 * time.Second / 5)
		if strings.Contains(i.Name(), "png") {
			nn = nn + ".png"
		} else {
			nn = nn + ".jpg"
		}
		// nn = nn + i.Name()[len(i.Name())-4:]

		if err := os.Rename("c:\\png\\"+i.Name(), "c:\\png\\"+nn); err != nil {

			fmt.Print(i.Name() + " ->  ")
			fmt.Println(nn)
		}
	}
}
