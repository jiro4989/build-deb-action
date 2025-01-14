package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type TemplateParam struct {
	Package       string
	Version       string
	InstalledSize string
	Architecture  string
	Maintainer    string
	Depends       string
	Homepage      string
	Section       string
	Description   string
}

func loadTemplateParam() (*TemplateParam, error) {
	p := &TemplateParam{
		Package:       os.Getenv("INPUT_PACKAGE"),
		Version:       os.Getenv("INPUT_VERSION"),
		InstalledSize: os.Getenv("INPUT_INSTALLED_SIZE"),
		Architecture:  os.Getenv("INPUT_ARCH"),
		Maintainer:    os.Getenv("INPUT_MAINTAINER"),
		Depends:       os.Getenv("INPUT_DEPENDS"),
		Homepage:      os.Getenv("INPUT_HOMEPAGE"),
		Section:       os.Getenv("INPUT_SECTION"),
		Description:   os.Getenv("INPUT_DESC"),
	}
	if err := p.validate(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *TemplateParam) validate() error {
	const errFmt = "'%s' must not be empty"
	type validateCases struct {
		param     string
		paramName string
	}

	// 必須パラメータの空チェック
	vc := []validateCases{
		{param: p.Package, paramName: "Package"},
		{param: p.Version, paramName: "Version"},
		{param: p.InstalledSize, paramName: "InstalledSize"},
		{param: p.Architecture, paramName: "Architecture"},
		{param: p.Maintainer, paramName: "Maintainer"},
	}
	for _, v := range vc {
		if v.param == "" {
			return fmt.Errorf(errFmt, v.paramName)
		}
	}

	return nil
}

func (p *TemplateParam) format() {
	p.Description = formatDescription(p.Description)
}

func render(tmpl string, p *TemplateParam) (string, error) {
	t, err := template.New("debian_control").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, p); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// formatDescription は control ファイルの Description 部用に複数行書式に変換する。
//
// 複数行の Description は以下のようになる。
//
// Description: 1 行目.
//
//	2 行目.
//	.
//	4 行目.
//
// 2 行目以降は、先頭が半角スペース 1 文字から始まる。
// 空行の場合は . のみ書く。
// といった感じ。
func formatDescription(s string) string {
	if s == "" {
		return ""
	}

	lines := make([]string, 0)
	for i, line := range strings.Split(s, "\n") {
		var l strings.Builder
		if 0 < i {
			l.WriteString(" ")
		}
		if line == "" {
			l.WriteString(".")
		}
		l.WriteString(line)
		lines = append(lines, l.String())
	}
	return strings.Join(lines, "\n")
}
