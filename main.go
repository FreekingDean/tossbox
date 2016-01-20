package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/jzelinskie/geddit"
	"os"
)

func main() {
	r, err := geddit.NewLoginSession("", "", "FuckTheCowboysBot/v0.2 /u/bananaboydean /u/harkins")
	if err != nil {
		fmt.Println(err)
	}
	_, err = r.AboutSubreddit("tossboxtest")
	if err != nil {
		fmt.Println(err)
	}
	cpt, err := r.NewCaptchaIden()
	if err != nil {
		fmt.Println(err)
	}
	capt := &geddit.Captcha{cpt, cpt}
	file, err := os.Open("test.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fInfo, _ := file.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)
	fReader := bufio.NewReader(file)
	fReader.Read(buf)
	fileBase64Str := base64.StdEncoding.EncodeToString(buf)
	fmt.Println(fInfo.Name())
	fmt.Println(len(fileBase64Str))
	r.Submit(geddit.NewTextSubmission("tossboxtest", fInfo.Name(), fileBase64Str[0:35000], false, capt))
}
