package repl

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"monkey/eval"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"strings"
)

const PROMPT = ">> "

func Run(path string, out io.Writer) {
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error occured while opening file: %s\n", err)
	}

	content_str := string(content[:])

	evaluated, err := exec(content_str, env, macroEnv)
	if err == nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	} else {
		io.WriteString(out, err.Error())
	}

}

func exec(src string, env *object.Environment, macroEnv *object.Environment) (object.Object, error) {
	l := lexer.New(src)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, formatParserErrors(p.Errors())
	}
	eval.DefineMacros(program, macroEnv)
	expanded := eval.ExpandMacros(program, macroEnv)
	evaluated := eval.Eval(expanded, env)
	return evaluated, nil

}
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "quit" {
			break
		}
		evaluated, err := exec(line, env, macroEnv)
		if err == nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, err.Error())
		}

	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func formatParserErrors(errs []string) error {
	var out bytes.Buffer
	out.WriteString(MONKEY_FACE)
	out.WriteString("Woops! We ran into some monkey business here!\n")
	out.WriteString(" parser errors:\n")
	for _, msg := range errs {
		out.WriteString("\t" + msg + "\n")
	}

	return errors.New(out.String())
}
