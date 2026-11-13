package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gogorch/gorch/gorchc/tpls"
	"github.com/gogorch/gorch/internal/lang/ast"
)

type codeGenerator struct {
	p     *ast.Primary
	todir string
	goCmd string
	// noTidy 为 true 时跳过 `go mod tidy`，避免生成代码时隐式改写调用方的 go.mod/go.sum
	noTidy bool
}

func (cg *codeGenerator) importOperator() error {
	if cg.goCmd == "" {
		cg.goCmd = "go"
	}

	if err := cg.ensureToDir(); err != nil {
		return err
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
	if cg.goCmd == "" {
		cg.goCmd = "go"
	}

	if err := cg.ensureToDir(); err != nil {
		return err
	}
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

// formatGo 用 go/format 规范化生成的源码，避免生成结果与 gofmt 不一致
func formatGo(name string, src []byte) []byte {
	formatted, err := format.Source(src)
	if err != nil {
		log.Printf("format %s failed: %v (keep raw content)", name, err)
		return src
	}
	return formatted
}

func (cg *codeGenerator) ensureToDir() error {
	if cg.todir == "" {
		return fmt.Errorf("output directory is empty")
	}
	if err := os.MkdirAll(cg.todir, 0755); err != nil {
		return fmt.Errorf("create output directory %q failed: %w", cg.todir, err)
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
	err = os.WriteFile(f, formatGo(f, buf.Bytes()), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (cg *codeGenerator) runModTidy() error {
	if cg.noTidy {
		log.Println("skip: go mod tidy (-no-tidy)")
		return nil
	}

	// 在目标代码所在模块根目录执行，而不是进程当前目录
	dir := cg.moduleRoot()
	log.Printf("start: go mod tidy (dir: %s)", dir)
	cmd := exec.Command(cg.goCmd, "mod", "tidy")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Printf("run `go mod tidy` out:\n%s\n", string(out))
	}
	if err != nil {
		return fmt.Errorf("run `go mod tidy` in %q failed: %w", dir, err)
	}
	log.Println("go mod tidy, done")
	return nil
}

// moduleRoot 从输出目录向上查找 go.mod，找不到则回退到进程当前目录
func (cg *codeGenerator) moduleRoot() string {
	dir, err := filepath.Abs(cg.todir)
	if err != nil {
		return "."
	}
	for {
		if _, statErr := os.Stat(filepath.Join(dir, "go.mod")); statErr == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "."
		}
		dir = parent
	}
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
	// 在模块根目录执行，否则临时程序的模块解析依赖调用方当前目录，容易失败
	cmd.Dir = cg.moduleRoot()
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Printf("generate all code out:\n%s", string(out))
	}
	if err != nil {
		return fmt.Errorf("run middle code failed: %w", err)
	}
	_ = os.RemoveAll(tmpPath)

	// 生成出来的算子代码同样做一次 gofmt，保证仓库里 `gofmt -l` 干净
	for _, name := range []string{"gen_operators.go"} {
		f := filepath.Join(cg.todir, name)
		src, readErr := os.ReadFile(f)
		if readErr != nil {
			return fmt.Errorf("read %q failed: %w", f, readErr)
		}
		if err := os.WriteFile(f, formatGo(f, src), 0644); err != nil {
			return fmt.Errorf("write %q failed: %w", f, err)
		}
	}

	return nil
}
