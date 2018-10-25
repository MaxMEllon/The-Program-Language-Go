package ftpd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/maxmellon/The-Program-Language-Go/ch08/ex02/dir"
)

type Mode int

const (
	ASCII Mode = iota
	BINARY
)

type Ftp struct {
	dataConn   net.Conn
	Conn       net.Conn
	currentDir string
	user       string
	mode       Mode
}

func New(conn net.Conn) *Ftp {
	dir := filepath.Dir(os.Getenv("HOME"))
	return &Ftp{
		Conn:       conn,
		currentDir: dir,
	}
}

func (f *Ftp) setUser(str string) {
	f.user = str
}

func (f *Ftp) setCurrentDir(dir string) {
	f.currentDir = dir
}

func (f *Ftp) setDataConn(dataConn net.Conn) {
	f.dataConn = dataConn
}

func (f *Ftp) write(str string) {
	_, err := fmt.Fprintf(f.Conn, "%s\r\n", str)
	if err != nil {
		log.Print(err)
	}
}

func (f *Ftp) feat() {
	f.write("211-Features:")
	f.write(" EPRT")
	f.write(" EPSV")
	f.write("211 End")
}

func (f *Ftp) pwd() {
	cmd := fmt.Sprintf("257 \"%s\"", f.currentDir)
	f.write(cmd)
}

func (f *Ftp) typ(i string) {
	switch i {
	case "I":
		f.mode = BINARY
		f.write("200 type set binary")
	case "A":
		f.mode = ASCII
		f.write("200 type set ASCII")
	}
}

func (f *Ftp) cwd(relativeDir string) {
	// refs: http://www.itsenka.com/contents/development/ftp/cd.html
	var dir string
	if !strings.HasSuffix(relativeDir, "/") {
		relativeDir += "/"
	}
	if strings.HasPrefix(relativeDir, "/") {
		dir = filepath.Dir(relativeDir)
	} else {
		dir = filepath.Dir(f.currentDir + "/" + relativeDir)
	}
	log.Print(dir)
	f.setCurrentDir(dir)
	f.write("250 CWD command successful")
}

func (f *Ftp) port(args string) {
	ports := strings.Split(args, ",")
	if len(ports) != 6 {
		f.write("550 parse error")
		return
	}

	p1, err := strconv.Atoi(ports[4])
	if err != nil {
		f.write("550 parse error")
		return
	}

	p2, err := strconv.Atoi(ports[5])
	if err != nil {
		f.write("550 parse error")
		return
	}
	port := p1*256 + p2
	address := fmt.Sprintf("%s.%s.%s.%s:%d", ports[0], ports[1], ports[2], ports[3], port)
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		f.write("550 connection error")
		return
	}
	f.setDataConn(dataConn)
	f.write("227 Entering Passive Mode")
}

func (f *Ftp) list() {
	files, err := dir.List(f.currentDir)
	if err != nil {
		f.write("550 no such file or directory")
	}
	log.Print(files)
	if f.dataConn != nil {
		f.write("150 ASCII mode data connection for file list")
		filesTxt := strings.Join(files, "\r\n")
		fmt.Fprintf(f.dataConn, "%s\r\n", filesTxt)
		f.dataConn.Close()
		f.write("226 Closing data connection")
	}
}

func (f *Ftp) retr(name string) {
	log.Print(f.currentDir + "/" + name)
	file, err := os.Open(f.currentDir + "/" + name)
	if err != nil {
		f.write("550 no such file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.dataConn.Write([]byte(scanner.Text()))
	}
	f.dataConn.Close()
	f.write("226 Closing data connection")
}

func HandleConnect(conn net.Conn) {
	// refs: https://www.itbook.info/network/ftp01.html
	f := New(conn)
	f.write("200 Welcome")
	input := bufio.NewScanner(f.Conn)
	for input.Scan() {
		cmds := strings.Fields(input.Text())
		log.Print(cmds)
		switch cmds[0] {
		case "USER":
			f.write("230 Login successful.")
			f.setUser(cmds[1])
		case "SYST":
			f.write("215")
		case "FEAT":
			f.feat()
		case "PWD":
			f.pwd()
		case "TYPE":
			f.typ(cmds[1])
		case "PORT":
			f.port(cmds[1])
		case "RETR":
			f.retr(cmds[1])
		case "CWD":
			f.cwd(cmds[1])
		case "LIST":
			f.list()
		case "QUIT":
		case "EPSV", "PASV", "NLST":
			f.write("502 Command not implemented")
		}
	}
}
