package main

import (
	"fmt"
	"io"
	"os"
)

type ExitCode int

const (
	exitCodeOK                  ExitCode = iota
	exitCodeErrCouldNotOpenFile ExitCode = 10 + iota
	exitCodeErrInvalidTemplateParam
	exitCodeErrCouldNotRender
	exitCodeErrCouldNotCreateFile
)

func main() {
	tmplFile := os.Args[1]
	outFile := os.Args[2]
	os.Exit(int(Main(tmplFile, outFile)))
}

func Main(tmplFile string, outFile string) ExitCode {
	// control ファイルのテンプレートを読み込み
	tf, err := os.Open(tmplFile)
	if err != nil {
		return exitCodeErrCouldNotOpenFile
	}
	defer tf.Close()
	b, err := io.ReadAll(tf)
	tmpl := string(b)

	// テンプレートへ値埋め込み
	p, err := loadTemplateParam()
	if err != nil {
		return exitCodeErrInvalidTemplateParam
	}
	p.format()
	s, err := render(tmpl, p)
	if err != nil {
		return exitCodeErrCouldNotRender
	}

	// 結果をファイルに出力
	f, err := os.Create(outFile)
	if err != nil {
		return exitCodeErrCouldNotCreateFile
	}
	defer f.Close()
	fmt.Fprint(f, s)

	return exitCodeOK
}
