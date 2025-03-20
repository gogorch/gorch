package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorch/gorch/gorchc/tpls"
	"github.com/gorch/gorch/internal/lang/ast"
)

type codeGenerator struct {
	p     *ast.Primary
	todir string
	goCmd string
}

func (cg *codeGenerator) importOperator() error {
	if cg.goCmd == "" {
		cg.goCmd = "go"
	}

	if err := cg.importStruct(); err != nil {
		return err
	}
	if err := cg.runModTidy(); err != nil {
		return err
	}

	return nil
}

func (cg *codeGenerator) releaseOperator() error {
	if err := cg.importStruct(); err != nil {
		return err
	}
	if err := cg.runModTidy(); err != nil {
		return err
	}
	if err := cg.generateMiddleCode(); err != nil {
		return err
	}

	return nil
}

type PackageOperators struct {
	PkgAlias  string
	Operators [][3]string
}

func (cg *codeGenerator) importStruct() error {
	t := template.New("import_operator")
	t, err := t.Parse(tpls.OperatorImportTpl)
	if err != nil {
		return err
	}
	buf := bytes.NewBufferString("")

	pkgs := map[string]*PackageOperators{}

	for i := range cg.p.Registers.Sort {
		r := cg.p.Registers.Sort[i]
		ol := r.Operators.Sort
		if len(ol) == 0 {
			continue
		}

		or := ol[0]
		pkg := fmt.Sprintf(`%s/%s`,
			strings.Trim(r.Pkg, "/"),
			strings.Trim(or.PkgPath, "/"))

		if _, has := pkgs[pkg]; !has {
			pkgs[pkg] = &PackageOperators{
				Operators: [][3]string{},
			}
		}

		pkgs[pkg].Operators = append(pkgs[pkg].Operators, [3]string{
			or.StructName,
			or.Name,
			strconv.FormatInt(or.Seq, 10),
		})
	}
	err = t.Execute(buf, map[string]any{"Pkgs": pkgs})
	if err != nil {
		return err
	}
	f := filepath.Join(cg.todir, "gen_import.go")
	_ = os.Remove(f)
	err = os.WriteFile(f, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (cg *codeGenerator) runModTidy() error {
	done := make(chan error)
	go func() {
		cmds := []string{"mod", "tidy"}
		cmd := exec.Command(cg.goCmd, cmds...)
		out, err := cmd.CombinedOutput()
		if len(out) > 0 {
			fmt.Printf("run `go mod tidy` out:\n%s\n", string(out))
		}
		if err != nil {
			log.Printf("run `go mod tidy` failed, error: %s\n", err)
			done <- err
		}
		done <- nil
	}()

	log.Println("start: go mod tidy")
	err := <-done
	log.Println("go mod tidy, done")
	return err
}

func (cg *codeGenerator) generateMiddleCode() error {
	t, err := template.New("middle_code").Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
	}).Parse(tpls.OperatorMiddleCodeTpl)
	if err != nil {
		return fmt.Errorf("parse middle code failed, error: %s", err)
	}
	buf := bytes.NewBuffer(nil)

	pkgs := map[string]*PackageOperators{}

	for i := range cg.p.Registers.Sort {
		r := cg.p.Registers.Sort[i]
		ol := r.Operators.Sort
		if len(ol) == 0 {
			continue
		}

		for j := range ol {
			or := ol[j]
			pkg := fmt.Sprintf(`%s/%s`,
				strings.Trim(r.Pkg, "/"),
				strings.Trim(or.PkgPath, "/"))

			if _, has := pkgs[pkg]; !has {
				pkgs[pkg] = &PackageOperators{
					PkgAlias:  "package" + strconv.Itoa(i) + strconv.Itoa(j),
					Operators: [][3]string{},
				}
			}

			pkgs[pkg].Operators = append(pkgs[pkg].Operators, [3]string{
				or.StructName,
				or.Name,
				strconv.FormatInt(or.Seq, 10),
			})
		}
	}
	err = t.Execute(buf, map[string]any{
		"Pkgs":  pkgs,
		"ToDir": cg.todir,
	})
	if err != nil {
		return fmt.Errorf("generate middle code failed, error: %s", err)
	}
	tmpPath := filepath.Join(cg.todir, "tmp")
	if err := os.RemoveAll(tmpPath); err != nil {
		return fmt.Errorf("remove tmp dir failed, error: %s", err)
	}
	if err := os.Mkdir(tmpPath, 0755); err != nil {
		return fmt.Errorf("create tmp dir failed, error: %s", err)
	}
	tmpMainPath := filepath.Join(tmpPath, "main.go")
	err = os.WriteFile(tmpMainPath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("write tmp file failed, error: %s", err)
	}

	cmd := exec.Command(cg.goCmd, "run", tmpMainPath)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Printf("generate all code out:\n%s", string(out))
	}
	if err != nil {
		log.Fatalf("failed with %s\n", err)
	}
	_ = os.RemoveAll(tmpPath)

	return nil
}
