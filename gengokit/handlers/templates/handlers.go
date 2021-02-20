package templates

// TODO: fix bug, pb.Empty when change the response to empty
const HandlerMethods = `
{{ with $te := .}}
		{{range $i := .Methods}}
			{{if ne .ResponseType.Name "Empty" }}
				func (s {{ToLower $te.ServiceName}}Service) {{.Name}}(ctx context.Context, in *pb.{{GoName .RequestType.Name}})(*pb.{{GoName $i.ResponseType.Name}}, error){
					var resp pb.{{GoName .ResponseType.Name}}
					return &resp, nil
				}
			{{else}}
				func (s {{ToLower $te.ServiceName}}Service) {{.Name}}(ctx context.Context, in *pb.{{GoName .RequestType.Name}})(*types.{{GoName $i.ResponseType.Name}}, error){
					return new(types.Empty), nil
				}
			{{end}}		
		{{end}}
{{- end}}
`

const Handlers = `
package handlers

import (
	"context"

	"github.com/gogo/protobuf/types"

	pb "{{.PBImportPath -}}"
)

// NewService returns a stateless implementation of Service.
func NewService() pb.{{GoName .Service.Name}}Server {
	return {{ToLower .Service.Name}}Service{}
}

type {{ToLower .Service.Name}}Service struct{}

{{with $te := . }}
	{{range $i := $te.Service.Methods}}
		{{if eq  $i.ResponseType.Name "Empty" }}
			func (s {{ToLower $te.Service.Name}}Service) {{$i.Name}}(ctx context.Context, in *pb.{{GoName $i.RequestType.Name}}) (*types.{{GoName $i.ResponseType.Name}}, error){
				return new(types.Empty), nil
			}
		{{else}}
			func (s {{ToLower $te.Service.Name}}Service) {{$i.Name}}(ctx context.Context, in *pb.{{GoName $i.RequestType.Name}}) (*pb.{{GoName $i.ResponseType.Name}}, error){
				var resp pb.{{GoName $i.ResponseType.Name}}
				return &resp, nil
			}
		{{end}}
	{{end}}
{{- end}}
`
