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
				Description:   "sample description.",
			},
			want: `Package: nimjson
Version: 1.0.0
Installed-Size: 999
Architecture: amd64
Maintainer: jiro4989
Depends: libc6 (>= 2.2.1), git
Homepage: https://github.com/jiro4989/nimjson
Section: unknown
Description: sample description.
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
