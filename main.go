package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aphecetche/c2b/csp"
)

const (
	cmakeMessageStart = `[== "CMake Server" ==[`
	cmakeMessageEnd   = `]== "CMake Server" ==]`
)

func trimServerTags(s string) (string, error) {
	msg := strings.Trim(s, "\n")

	validStart := regexp.MustCompile(fmt.Sprintf("^%s", regexp.QuoteMeta(cmakeMessageStart)))

	if !(validStart.MatchString(msg)) {
		return "", fmt.Errorf("no start tag")
	}

	msg = strings.Trim(msg, cmakeMessageStart)

	validEnd := regexp.MustCompile(fmt.Sprintf("%s$", regexp.QuoteMeta(cmakeMessageEnd)))

	if !(validEnd.MatchString(msg)) {
		return "", fmt.Errorf("no end tag")
	}

	msg = strings.Trim(msg, cmakeMessageEnd)
	return msg, nil
}

func cmakeMsg(b []byte) (csp.Hello, error) {

	msg, err := trimServerTags(string(b))
	if err != nil {
		return csp.Hello{}, err
	}

	h, err := csp.UnmarshalHello([]byte(msg))

	return h, nil
}

func reader(r io.Reader) {
	maxsize := 1024 * 1024
	buf := make([]byte, maxsize)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		h, err := cmakeMsg(buf[0:n])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Client got:%v\n", h)
	}
}
func main() {

	c, err := net.Dial("unix", "/tmp/c2b.sock")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()

	go reader(c)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("I'm alive - please enter e to exit or any other key to continue ")
		text, _ := reader.ReadString('\n')
		fmt.Printf("text=%s\n", text)
		if strings.Trim(text, " \n") == "e" {
			fmt.Println("bye")
			break
		}
		time.Sleep(10 * time.Second)
	}
}
