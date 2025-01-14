package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		desc string
		p *TemplateParam
		wantErr bool
	} {
		{
			desc: "ok: validation ok",
			p: &TemplateParam{
				Package: "nimjson",
				Version: "v1.0.0",
				InstalledSize: "999",
				Architecture: "amd64",
				Maintainer: "jiro4989",
			},
			wantErr: false,
		},
		{
			desc: "ng: Package is empty",
			p: &TemplateParam{
				Package: "",
				Version: "v1.0.0",
				InstalledSize: "999",
				Architecture: "amd64",
				Maintainer: "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: Version is empty",
			p: &TemplateParam{
				Package: "nimjson",
				Version: "",
				InstalledSize: "999",
				Architecture: "amd64",
				Maintainer: "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: InstalledSize is empty",
			p: &TemplateParam{
				Package: "nimjson",
				Version: "v1.0.0",
				InstalledSize: "",
				Architecture: "amd64",
				Maintainer: "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: Architecture is empty",
			p: &TemplateParam{
				Package: "nimjson",
				Version: "v1.0.0",
				InstalledSize: "999",
				Architecture: "",
				Maintainer: "jiro4989",
			},
			wantErr: true,
		},
		{
			desc: "ng: Maintainer is empty",
			p: &TemplateParam{
				Package: "nimjson",
				Version: "v1.0.0",
				InstalledSize: "999",
				Architecture: "amd64",
				Maintainer: "",
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
