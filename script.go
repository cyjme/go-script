package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/build", build)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func build(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start build")
	run("pwd")
	run("git", "pull")
	run("git", "status")
	run("npm", "install")
	run("composer", "install")
	run("pwd")
}

func run(command string, param ...string) bool {
	cmd := exec.Command(command, param...)

	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
