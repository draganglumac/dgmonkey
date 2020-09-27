// compiler/compiler.go

package compiler

import (
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

type Compiler struct {
	instructions code.Instructions
	constants []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants: []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

type Bytecode struct {
	Instructions code.Instructions
	Constants []object.Object
}

func (c *Compiler) Bytecode() *Bytecode  {
	return &Bytecode{
		Instructions: code.Instructions{},
		Constants: []object.Object{},
	}
}
