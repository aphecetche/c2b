package main

import (
	"bufio"
	"bytes"
	"flag"
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

func read(r io.Reader) ([]cmake.Message, error) {

	var p []byte
	const maxsize = 8192

	var line string
	started := false
	var payload string

	one := make([]byte, maxsize)

	for {
		n, err := r.Read(one)
		if err != nil && err != io.EOF {
			panic(err)
		}
		p = append(p, one...)
		if n < maxsize {
			break
		}
	}

	p = append(p, '\n')

	messages := []cmake.Message{}

	reader := bufio.NewReader(bytes.NewReader(p))

	var err error
	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		line = strings.Trim(line, "\n")

		switch {

		case line == cmake.MessageStart:
			payload = ""
			started = true

		case line == cmake.MessageEnd && started:
			m, err := cmake.NewMessage([]byte(payload))
			if err != nil {
				fmt.Println("err=", err)
				break
			}

			messages = append(messages, m)
			started = false
		default:
			payload += line
		}

	}
	if err != io.EOF {
		return nil, err
	}
	return messages, nil
}

func runserver(quit chan bool, config cmake.BuildConfig) {
	var cmd *exec.Cmd
	for {
		select {
		case <-quit:
			fmt.Println("quit received")
			exec.Command("rm", "-f", config.SocketName()).Run()
			os.Exit(1)
		default:
			if cmd == nil {
				fmt.Printf("starting server using socket %s\n", config.SocketName())
				cmd = exec.Command("cmake", "-E", "server", fmt.Sprintf("--pipe=%s", config.SocketName()), "--experimental")
				for _, e := range config.Env {
					cmd.Env = append(cmd.Env, e)
				}
				cmd.Env = append(cmd.Env, os.Environ()...)
				cmd.Run()
			}
		}
	}
}

func readOne(r io.Reader, msgType cmake.MessageType) error {
	messages, err := read(r)
	if err != nil {
		return err
	}
	if len(messages) != 1 {
		return fmt.Errorf("Expected only one message but got %d", len(messages))
	}
	if messages[0].Type() != msgType {
		fmt.Println(messages[0])
		return fmt.Errorf("Expected message of type %s but got %s",
			msgType.String(), messages[0].Type().String())
	}
	return nil
}

// should get hello message from server
func hello(r io.Reader) {
	err := readOne(r, cmake.HelloMsg)
	if err != nil {
		panic(err)
	}
}

func handshake(rw io.ReadWriter, config cmake.BuildConfig) {

	handshake := cmake.NewMessageHandshake(1, config.SourceDir, config.BuildDir, config.Generator)

	// send handshake
	cmake.Write(handshake, rw)

	// should get a reply to the handshake
	err := readOne(rw, cmake.ReplyMsg)

	if err != nil {
		panic(err)
	}
}

func setGlobalSettings(rw io.ReadWriter, config cmake.BuildConfig, coldStart bool) {
	var m cmake.Message
	if coldStart {
		m = cmake.NewMessageSetGlobalSettingsColdStart(config.SourceDir, config.BuildDir, config.Generator)
	} else {
		m = cmake.NewMessageSetGlobalSettingsWarmStart(config.BuildDir)
	}
	cmake.Write(m, rw)
	err := readOne(rw, cmake.ReplyMsg)
	if err != nil {
		panic(err)
	}
}

func writeAndReadReply(rw io.ReadWriter, m cmake.Message) (cmake.Message, error) {
	cmake.Write(m, rw)
	var err error
	var h []cmake.Message
	for err == nil {
		h, err = read(rw)
		for _, m := range h {
			switch m.Type() {
			case cmake.ReplyMsg:
				return m, nil
			case cmake.ErrorMsg:
				return nil, fmt.Errorf(m.String())
			case cmake.ProgressMsg:
			default:
				fmt.Println(m)
			}
		}
	}
	return nil, nil
}

func configure(rw io.ReadWriter, config cmake.BuildConfig) (cmake.Message, error) {
	configureMsg := cmake.NewMessageConfigure(config.Configure)
	return writeAndReadReply(rw, configureMsg)
}

func compute(rw io.ReadWriter) (cmake.Message, error) {
	computeMsg := cmake.NewMessageCompute()
	return writeAndReadReply(rw, computeMsg)
}

func codemodel(rw io.ReadWriter) (cmake.Message, error) {
	codemodelMsg := cmake.NewMessageCodeModel()
	return writeAndReadReply(rw, codemodelMsg)
}

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		done <- true
	}()

	build := flag.String("build", "o2", "which package to generate rules for")

	flag.Parse()

	var config cmake.BuildConfig

	switch *build {
	case "o2":
		config = cmake.O2Config
	case "fairlogger":
		config = cmake.FairLoggerConfig
	default:
		log.Fatalf("Unknown build config %s", *build)
	}

	go runserver(done, config)

	time.Sleep(1 * time.Second)

	c, err := net.Dial("unix", config.SocketName())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()

	hello(c)
	handshake(c, config)
	setGlobalSettings(c, config, false)

	_, err = configure(c, config)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = compute(c)
	if err != nil {
		log.Fatalln(err)
	}

	m, err := codemodel(c)

	cm := m.(*cmake.MessageReply)

	treatReply(cm, config)
}

func treatReply(cm *cmake.MessageReply, config cmake.BuildConfig) {

	fmt.Println("# of configurations:", len(cm.Configurations))
	conf := cm.Configurations[0]
	fmt.Println("# of projects:", len(conf.Projects))
	proj := conf.Projects[0]
	fmt.Println("# of targets:", len(proj.Targets))
	ttypes := make(map[string]int)
	for _, t := range proj.Targets {
		if len(t.Artifacts) > 0 {
			fmt.Println("------------")
			fmt.Println("FULLNAME:", t.FullName)
			fmt.Println("BUILDIR:", t.BuildDirectory)
			fmt.Println(len(t.Artifacts), " ARTIFACTS:", t.Artifacts)
			fmt.Println(t.Type)
			fmt.Println(len(t.FileGroups), " FILEGROUPS:")
			for i, fg := range t.FileGroups {
				fmt.Println("FileGroup", i, ":", fg)
			}
		}
		ttypes[t.Type]++
	}
	fmt.Println()
	fmt.Println()
	for k, t := range ttypes {
		fmt.Println(k, t)
	}
}
