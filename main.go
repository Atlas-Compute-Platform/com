package main

/*
	Atlas Compute Service
	Thijs Haker
*/

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/Atlas-Compute-Platform/lib"
)

func apiExec(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		cmdStr = r.URL.Query().Get(lib.KEY_CMD)
		argStr = r.URL.Query().Get(lib.KEY_ARG)
		cmd    *exec.Cmd
		buf    []byte
		err    error
	)

	cmd = exec.Command(cmdStr, argStr)
	cmd.Stdin = r.Body

	if buf, err = cmd.Output(); err != nil {
		lib.LogError(os.Stderr, "main.apiExec", err)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(buf))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Atlas Compute Service %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var netAddr = flag.String("p", lib.PORT, "Specify port")
	flag.Usage = usage
	flag.Parse()

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/exec", apiExec)

	http.ListenAndServe(*netAddr, nil)
}
