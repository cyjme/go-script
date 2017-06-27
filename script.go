package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/build",build)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func build(w http.ResponseWriter, r *http.Request){
	fmt.Println("i am build")
	pwdCmd := exec.Command("pwd")
	pwdOutput, _ := pwdCmd.Output()
	fmt.Println(string(pwdOutput))
}

