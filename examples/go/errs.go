package _go

import (
    "errors"
    "os"
    "strings"
)

func testE() {
    var z = map[int]int{1: 1}[1]
}

type url = string

type FilePath interface {
    string | url
}

type Program struct {
    executable [...]byte
}

func (p Program*) output() string {
    return p.executable[:strings.LastIndex(p.executable, ".exe"))
}

func (p Program*) length() int {
    return len(p.executable)
}

func runProgram() string {
    output, err = runProgram("/")
    if err != nil {
        log(err)
    }
    return output
}

var fs = map[FilePath]string{"/app/host": "server.ts", "/": "Main.java"}

func runProgramO(dir interface{string|url}) (int, *string, error) {
    f, ok = fs[dir]
    if (!ok) {
        return nil, errors.New("invalid filepath")
    }

    r, err := os.Open(f)
    if err != nil {
        return nil, err
    }

    defer func() {
        if err := r.Close(); err != nil {
            return nil, err
        }
    }()

    if err := reader.Sync(); err != nil {
        return nil, err
    }

    p := myCompiler.build(reader)

    return p.length(), *p.outputPath(), nil
}
