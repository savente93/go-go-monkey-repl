package repl

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"monkey/compiler"
	"monkey/eval"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/vm"
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
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")

		}
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
	// env := object.NewEnvironment()
	// macroEnv := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "quit" || strings.TrimSpace(line) == "q" {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			formatParserErrors(p.Errors())
		}

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "whoops! Compilations failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "whoops exectuing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.LastPopedStackElm()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")

		// evaluated, err := exec(line, env, macroEnv)
		// if err == nil {
		// 	if evaluated != nil {
		// 		io.WriteString(out, evaluated.Inspect())
		// 		io.WriteString(out, "\n")

		// 	}
		// } else {
		// 	io.WriteString(out, err.Error())
		// }

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
