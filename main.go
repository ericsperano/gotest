package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("go", "test", "./...", "-cover")
	out, err := cmd.Output()
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(-1)
	}
	data := string(out)
	for _, line := range strings.Split(data, "\n") {
		fmt.Println(convertLine(line))
	}
}

func convertLine(line string) string {
	if len(line) == 0 {
		return ""
	}
	fields := strings.Fields(line)
	percent := ""
	if fields[0] == "ok" {
		percent = fields[4]
	}
	return fmt.Sprintf("%-2s  %-60s %5s", fields[0], fields[1], percent)
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
