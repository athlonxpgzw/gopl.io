package main

import (
    "github.com/athlonxpgzw/gopl.io/ch7/ex7_15/eval"
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    input := bufio.NewScanner(os.Stdin)
    fmt.Println("Welcome to Calculator")
outer:
    for {
        var expr eval.Expr
        var err error
        fmt.Println("Please input the a single expression")
        if err := input.Scan(); err != true {
            fmt.Println("Scann err with "+input.Err().Error())
        }
        if expr, err = eval.Parse(input.Text()); err != nil {
            fmt.Println("Bad expression "+err.Error())
            continue outer
        }
        vars := make(map[eval.Var]bool)
        if err = expr.Check(vars); err != nil {
            fmt.Println("Expression check failed "+err.Error())
            continue outer
        }
        env := make(eval.Env)
        fmt.Println("Please provide values for any variables")
        for var_name, _ := range vars {
            var var_value float64
            fmt.Printf("Please enter %s values:\n", var_name)
            if err := input.Scan(); err != true {
                fmt.Println("Scann err with "+input.Err().Error())
            }
            if var_value, err = strconv.ParseFloat(input.Text(), 64); err != nil {
                fmt.Println("Please enter the right value")
                continue outer
            }
            env[var_name] = var_value
        }
        fmt.Printf("The result is : %f\n", expr.Eval(env))
        continue outer
    }
}
