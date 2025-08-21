package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = WithFunction{}
)

func NewWithFunction() function.Function {
	return WithFunction{}
}

type WithFunction struct{}

func (r WithFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "with"
}

func (r WithFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "with function",
		MarkdownDescription: "Return string with sleep(n).",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "str",
				MarkdownDescription: "String to return.",
			},
			function.StringParameter{
				Name:                "dur",
				MarkdownDescription: "Duration string. see https://pkg.go.dev/time#ParseDuration",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r WithFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var str string
	var durStr string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &durStr))

	if resp.Error != nil {
		return
	}

	dur, err := time.ParseDuration(durStr)

	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	time.Sleep(dur)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, str))
}
