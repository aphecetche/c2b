package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/aphecetche/c2b/cmake"
)

func trimServerTags(b []byte) ([]byte, error) {
	s := string(b)
	msg := strings.Trim(s, "\n")

	validStart := regexp.MustCompile(fmt.Sprintf("^%s", regexp.QuoteMeta(cmake.MessageStart)))

	if !(validStart.MatchString(msg)) {
		return nil, fmt.Errorf("no start tag")
	}

	msg = strings.Trim(msg, cmake.MessageStart)

	validEnd := regexp.MustCompile(fmt.Sprintf("%s$", regexp.QuoteMeta(cmake.MessageEnd)))

	if !(validEnd.MatchString(msg)) {
		return nil, fmt.Errorf("no end tag")
	}

	msg = strings.Trim(msg, cmake.MessageEnd)
	return []byte(msg), nil
}

func read(r io.Reader, msgType cmake.MessageType, alternate *cmake.MessageType) (cmake.Message, error) {
	maxsize := 1024 * 1024
	buf := make([]byte, maxsize)
	n, err := r.Read(buf[:])
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		panic(err)
	}
	msg, err := trimServerTags(buf[0:n])
	if err != nil {
		panic(err)
	}
	h, err := cmake.NewMessage(msg, msgType)
	if err != nil {
		if alternate != nil {
			h, err = cmake.NewMessage(msg, *alternate)
		}
	}
	if err != nil {
		fmt.Printf("READ:>%s<\n", string(msg))
		fmt.Printf("READ(%d):>>%s<<\n", n, string(buf[0:n]))
		return nil, err
	}
	fmt.Println(h)
	return h, nil
}

func runserver(quit chan bool) {
	var cmd *exec.Cmd
	for {
		select {
		case <-quit:
			fmt.Println("quit received")
			exec.Command("rm", "-f", "/tmp/c2b.sock").Run()
			os.Exit(1)
		default:
			if cmd == nil {
				fmt.Println("starting server")
				cmd = exec.Command("cmake", "-E", "server", "--pipe=/tmp/c2b.sock", "--experimental")
				for _, e := range cmake.BuildConfig.Env {
					cmd.Env = append(cmd.Env, e)
				}
				cmd.Env = append(cmd.Env, os.Environ()...)
				cmd.Run()
			}
		}
	}
}

// should get hello message from server
func hello(c net.Conn) {
	read(c, cmake.HelloMsg, nil)
}

func handshake(c net.Conn) {

	handshake := cmake.NewMessageHandshake(1, cmake.BuildConfig.SourceDir, cmake.BuildConfig.BuildDir, cmake.BuildConfig.Generator)

	// send handshake
	cmake.Write(handshake, c)

	// should get a reply to the handshake
	_, err := read(c, cmake.ReplyMsg, nil)

	if err != nil {
		panic(err)
	}
}

func setGlobalSettings(c net.Conn) {
	setGlobalSettings := cmake.NewMessageSetGlobalSettings(cmake.BuildConfig.SourceDir, cmake.BuildConfig.BuildDir, cmake.BuildConfig.Generator)
	cmake.Write(setGlobalSettings, c)

	_, err := read(c, cmake.ReplyMsg, nil)
	if err != nil {
		panic(err)
	}
}

func configure(c net.Conn) {
	configure := cmake.NewMessageConfigure(cmake.BuildConfig.Configure)
	cmake.Write(configure, c)
	var err error
	p := cmake.ProgressMsg
	for {
		for err == nil {
			_, err = read(c, cmake.MessageMsg, &p)
		}
		os.Exit(3)
	}
}

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		done <- true
	}()

	go runserver(done)

	time.Sleep(1 * time.Second)

	c, err := net.Dial("unix", "/tmp/c2b.sock")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()

	hello(c)
	handshake(c)
	setGlobalSettings(c)

	configure(c)
	// compute := cmake.NewMessageCompute()
	// cmake.Write(compute, c)

	// _, err = read(c, cmake.ReplyMsg)
	// if err != nil {
	// 	panic(err)
	// }
}
