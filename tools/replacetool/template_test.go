package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		desc    string
		p       *TemplateParam
		wantErr bool
	}{
		{
			desc: "ok: validation ok",
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
			},
			wantErr: false,
		},
		{
			desc: "ng: Package is empty",
			p: &TemplateParam{
				Package:       "",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: Version is empty",
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: InstalledSize is empty",
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: Architecture is empty",
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "",
				Maintainer:    "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: Maintainer is empty",
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			err := tt.p.validate()
			if tt.wantErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)
		})
	}
}

func TestRender(t *testing.T) {
	f, err := os.Open(filepath.Join("template", "control"))
	assert.NoError(t, err)
	defer f.Close()
	b, err := io.ReadAll(f)
	assert.NoError(t, err)
	tmpl := string(b)

	tests := []struct {
		desc    string
		tmpl    string
		p       *TemplateParam
		want    string
		wantErr bool
	}{
		{
			desc: "ok: minimal pattern",
			tmpl: tmpl,
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
`,
			wantErr: false,
		},
		{
			desc: "ok: maximum pattern",
			tmpl: tmpl,
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
				Depends:       "libc6 (>= 2.2.1), git",
				Homepage:      "https://github.com/jiro4989/nimjson",
				Section:       "unknown",
				Priority:      "required",
				Description:   "sample description.\nsample description2.\n\nsample description3.",
				Conflicts:     "hello-traditional",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
Depends: libc6 (>= 2.2.1), git
Homepage: https://github.com/jiro4989/nimjson
Section: unknown
Priority: required
Description: sample description.
 sample description2.
 .
 sample description3.
Conflicts: hello-traditional
`,
			wantErr: false,
		},
		{
			desc: "ok: homepage and description",
			tmpl: tmpl,
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
				Homepage:      "https://github.com/jiro4989/nimjson",
				Description:   "sample description.",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
Homepage: https://github.com/jiro4989/nimjson
Description: sample description.
`,
			wantErr: false,
		},
		{
			desc: "ok: depends and section",
			tmpl: tmpl,
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
				Depends:       "libc6 (>= 2.2.1), git",
				Section:       "unknown",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
Depends: libc6 (>= 2.2.1), git
Section: unknown
`,
			wantErr: false,
		},
		{
			desc: "ok: depends and pre-depends",
			tmpl: tmpl,
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
				Depends:       "libc6 (>= 2.2.1), git",
				PreDepends:    "zstd, zsh",
				Section:       "unknown",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
Depends: libc6 (>= 2.2.1), git
Pre-Depends: zstd, zsh
Section: unknown
`,
			wantErr: false,
		},
		{
			// text/template を使っているので HTML 文字列がエスケープされたりしないはず
			desc: "ok: multiline description",
			tmpl: tmpl,
			p: &TemplateParam{
				Package:       "nimjson",
				Version:       "v1.0.0",
				InstalledSize: "999",
				Architecture:  "amd64",
				Maintainer:    "jiro4989",
				Description:   "sample1.\nsample2.\n\nsample3.\n<div>test</div><br>\n<a href=\"https://github.com/jiro4989\">URL</a>",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
Description: sample1.
 sample2.
 .
 sample3.
 <div>test</div><br>
 <a href="https://github.com/jiro4989">URL</a>
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			tt.p.format()
			got, err := render(tt.tmpl, tt.p)
			if tt.wantErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestFormatDescription(t *testing.T) {
	tests := []struct {
		desc string
		s    string
		want string
	}{
		{
			desc: "ok: 1 line",
			s:    "1 line",
			want: "1 line",
		},
		{
			desc: "ok: ignore trailing newline",
			s:    "1 line\n",
			want: "1 line",
		},
		{
			desc: "ok: ignore prefix newlines and whitespaces",
			s:    " \n  \n  1 line",
			want: "1 line",
		},
		{
			desc: "ok: multi lines",
			s:    "1 line\n2 line\n3 line",
			want: `1 line
 2 line
 3 line`,
		},
		{
			desc: "ok: multi lines and ignore trailing whitespaces",
			s:    "1 line\n2 line\n3 line     ",
			want: `1 line
 2 line
 3 line`,
		},
		{
			desc: "ok: multi lines and ignore newlines and trailing whitespaces",
			s:    "1 line\n2 line\n3 line     \n \n  \n\n ",
			want: `1 line
 2 line
 3 line`,
		},
		{
			desc: "ok: multi lines and blank line",
			s:    "1 line\n2 line\n3 line\n\n4 line\n",
			want: `1 line
 2 line
 3 line
 .
 4 line`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := formatDescription(tt.s)
			assert.Equal(tt.want, got)
		})
	}
}
