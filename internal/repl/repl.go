package repl

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

type Repl struct {
    keepRunning bool
    reader      *bufio.Reader
}

func NewRepl() Repl {
    return Repl{true, bufio.NewReader(os.Stdin)}
}

func (r *Repl) printRepl() {
    fmt.Print("> ")
}

var clear map[string]func()

func init() {
    clear = make(map[string]func())
    clear["linux"] = func() {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        err := cmd.Run()
        if err != nil {
            return
        }
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        err := cmd.Run()
        if err != nil {
            return
        }
    }
}

func callClear() {
    value, ok := clear[runtime.GOOS]
    if ok {
        value()
    } else {
        panic("Clear not supported on this platform")
    }
}

var commands = map[string]interface{}{
    "say": func(str string) {
        fmt.Println(str)
    },
    "cls": callClear,
}

func (r *Repl) getInput() {
    input, err := r.reader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    input = strings.TrimSuffix(input, "\r\n")
    if input == "exit" {
        r.keepRunning = false
        return
    }
    inputs := strings.SplitN(input, " ", 2)
    if function, ok := commands[inputs[0]]; ok {
        if len(inputs) == 1 {
            function.(func())()
        } else {
            arg := inputs[1]
            if arg[0] == '"' {
                arg = arg[1:]
            } else {
                fmt.Println("Opening quote missing on arg")
                return
            }
            if arg[len(arg)-1] == '"' {
                arg = arg[:len(arg)-1]
            } else {
                fmt.Println("Closing quote missing on arg")
                return
            }
            function.(func(string))(arg)
        }
    } else {
        fmt.Println("Unknown command")
    }
}

func (r *Repl) Run() {
    for r.keepRunning {
        r.printRepl()
        r.getInput()
        if !r.keepRunning {
            break
        }
    }
}
