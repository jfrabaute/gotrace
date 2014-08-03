package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jfrabaute/libtrace"
)

func main() {

	log.SetFlags(0)

	if cmd, err := getCommand(); err != nil {
		log.Fatal(err)
	} else {

		var tracer libtrace.Tracer
		tracer = libtrace.NewTracer(cmd)
		tracer.RegisterGlobalCbOnExit(genericCallback)

		err = tracer.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getCommand() (cmd *exec.Cmd, err error) {
	if len(os.Args) < 2 {
		err = fmt.Errorf("Usage: %s cmd [cmd] ...", os.Args[0])
	} else {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	}
	return
}

func genericCallback(trace *libtrace.Trace) {
	args := ""
	if len(trace.Args) > 0 {
		for i, arg := range trace.Args {
			if i > 0 {
				args += ", "
			}
			args += fmt.Sprintf("%v", arg)
		}
	}
	log.Printf("%s(%s) = %d %s", trace.Signature.Name, args, trace.Return.Code, trace.Return.Description)
}
