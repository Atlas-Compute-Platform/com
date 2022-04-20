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
		cmdStr string
		argStr string
		req    lib.Dict
		cmd    *exec.Cmd
		buf    []byte
		err    error
	)

	if req, err = lib.ReceiveDict(r.Body); err != nil {
		lib.LogError(os.Stderr, "main.apiExec", err)
		fmt.Fprint(w, err)
		return
	}

	cmdStr = req[lib.KEY_CMD]
	argStr = req[lib.KEY_ARG]
	cmd = exec.Command(cmdStr, argStr)

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
