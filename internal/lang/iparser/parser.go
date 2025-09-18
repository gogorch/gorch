package iparser

import (
	"errors"
	"os"
	"path"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gogorch/gorch/internal/lang/ast"
	alr "github.com/gogorch/gorch/internal/lang/iantlr"
	"github.com/gogorch/gorch/internal/lang/utils"
)

var (
	ErrSelfcheckgorchGramer = errors.New("selfcheck gorch gramer error")
	ErrParsegorchFile       = errors.New("parse gorch file error")
)

func ParseDirectory(dir string) (*ast.Primary, error) {
	matches, err := utils.GlobGorchFiles(dir)
	if err != nil {
		return nil, err
	}
	p := ast.NewPrimary()

	parserListener := ast.NewGorchlangParserListener()
	parserListener.Primary = p

	for _, f := range matches {
		parserListener.Stack.Push(p)

		in, err := NewIoStream(path.Join(dir, f))
		if err != nil {
			return nil, err
		}

		parserListener.File = f

		lexer := alr.NewgorchlangLexer(in)
		el := NewErrorListener(os.Stderr, f, parserListener)
		lexer.AddErrorListener(el)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		psr := alr.NewgorchlangParser(stream)
		psr.RemoveErrorListeners()
		psr.AddErrorListener(el)
		psr.BuildParseTrees = true
		antlr.ParseTreeWalkerDefault.Walk(parserListener, psr.Primary())

		if len(parserListener.ParseErrors) > 0 {
			errs := []error{ErrParsegorchFile}
			errs = append(errs, parserListener.ParseErrors...)
			return nil, errors.Join(errs...)
		}
	}

	if err := p.SelfCheck(); err != nil {
		return nil, errors.Join(ErrSelfcheckgorchGramer, err)
	}

	return p, nil
}

func Parse(s string) (*ast.Primary, error) {
	p := ast.NewPrimary()
	parserListener := ast.NewGorchlangParserListener()
	parserListener.Stack.Push(p)
	parserListener.Primary = p

	in := antlr.NewInputStream(s)

	lexer := alr.NewgorchlangLexer(in)
	el := NewErrorListener(os.Stderr, "", parserListener)
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	psr := alr.NewgorchlangParser(stream)
	psr.RemoveErrorListeners()
	psr.AddErrorListener(el)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(parserListener, psr.Primary())

	if len(parserListener.ParseErrors) > 0 {
		errs := []error{ErrParsegorchFile}
		errs = append(errs, parserListener.ParseErrors...)
		return nil, errors.Join(errs...)
	}

	if err := p.SelfCheck(); err != nil {
		return nil, errors.Join(ErrSelfcheckgorchGramer, err)
	}

	return p, nil
}
