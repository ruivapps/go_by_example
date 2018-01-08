/*
install ssh lib: go get golang.org/x/crypto/ssh
ssh documentation: https://godoc.org/golang.org/x/crypto/ssh
*/
package main

import (
	"bytes" // assign ssh stdout to bytes.Buffer locally
	"fmt"
	"golang.org/x/crypto/ssh" // ssh library
	"io/ioutil"
	"net" // need net.Addr for ssh HostKeyCallback
	"os"
	"strings"
)

func PrivateKey(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func MySSH(ip_addr string) (*string, error) {
	user := os.Getenv("USER")
	ssh_home := os.Getenv("HOME") + "/.ssh/id_rsa"
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// to use password auth, use ssh.Password("password") instead PrivateKey function
			PrivateKey(ssh_home)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	connection, err := ssh.Dial("tcp", ip_addr+":22", sshConfig)
	defer connection.Close()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect: %s", err)
	}
	session, err := connection.NewSession()
	defer session.Close()
	if err != nil {
		return nil, fmt.Errorf("Failed to create new session: %s", err)
	}

	// create pseudo terminal (xterm 80 columns by 40 rows)
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed
		ssh.TTY_OP_OSPEED: 14400, // output speed
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, fmt.Errorf("request for pseudo terminal failed: %s", err)
	}
	var b bytes.Buffer
	session.Stdout = &b
	command := "/usr/bin/uname -a && date"
	if len(os.Args) > 2 {
		command = strings.Join(os.Args[2:], " ")
	}
	if err := session.Run(command); err != nil {
		return nil, fmt.Errorf("error execute command from remote: %s", err)
	}
	fmt.Println(b.String())
	return nil, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ssh target_ip\n")
		os.Exit(1)
	}
	_, err := MySSH(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
}
