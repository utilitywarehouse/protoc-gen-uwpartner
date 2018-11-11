package service

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type ServiceModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func Module() *ServiceModule {
	return &ServiceModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *ServiceModule) Name() string {
	return "service"
}

func (m *ServiceModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())

	// Required params
	if m.ctx.Params().Str("namespace") == "" {
		m.AddError("`namespace` param is required")
	}

	tpl := template.New("service").Funcs(map[string]interface{}{
		"package": m.ctx.PackageName,
		"option":  m.ctx.Params().Str,
	})
	m.tpl = template.Must(tpl.Parse(partnerServiceTpl))
}

func (m *ServiceModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		m.Debugf("Generating for target: %s", t.Name())
		m.generate(t)
	}

	return m.Artifacts()
}

func (m *ServiceModule) generate(f pgs.File) {
	if len(f.Services()) == 0 {
		m.Debugf("No services contained in `%s`, skipping", f.Name())
		return
	}

	name := m.ctx.OutputPath(f).SetExt(".uwpartner.go")
	m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
}
