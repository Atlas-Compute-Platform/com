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
	"strings"

	"github.com/Atlas-Compute-Platform/lib"
)

func apiExec(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		req     lib.Dict
		buf     []byte
		ok      bool
		cmd     *exec.Cmd
		cmdStr  string
		argStr  string
		argList []string
		stdStr  string
		err     error
	)

	if req, err = lib.ReceiveDict(r.Body); err != nil {
		lib.LogError(w, "main.apiExec", err)
		return
	}

	cmdStr = req[lib.KEY_CMD]
	argStr = req[lib.KEY_ARG]
	argList = strings.Fields(argStr)
	cmd = exec.Command(cmdStr, argList...)

	if stdStr, ok = req["std"]; ok {
		cmd.Stdin = strings.NewReader(stdStr)
	}

	if buf, err = cmd.Output(); err != nil {
		lib.LogError(w, "main.apiExec", err)
		lib.LogError(os.Stderr, "main.apiExec", err)
		return
	}
	fmt.Fprint(w, string(buf))
}

func main() {
	lib.SvcName = "Atlas Compute Service"
	lib.SvcVers = "1.0"

	var netAddr = flag.String("p", lib.PORT, "Specify port")
	flag.Usage = lib.Usage
	flag.Parse()

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/exec", apiExec)

	http.ListenAndServe(*netAddr, nil)
}
