package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err == nil {
		fmt.Println(string(dateOut))
	}

	//
	grepCmd := exec.Command("grep", "hello")
	in, err := grepCmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	out, err := grepCmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	err = grepCmd.Start()
	if err != nil {
		panic(err)
	}
	in.Write([]byte("hello grep\ngoodbye grep"))
	in.Close()
	grepByte, _ := ioutil.ReadAll(out)
	grepCmd.Wait()
	fmt.Println(string(grepByte))
}
