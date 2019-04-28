package tesseract

import (
	"context"
	"fmt"
	"strings"
)

type ContextKey string
type Configs map[string]string

const (
	languagesKey ContextKey = "langs"
	configsKey   ContextKey = "configs"
	psmKey       ContextKey = "psm"
	oemKey       ContextKey = "oem"
)

func contextOrBackground(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}

func WithLanguages(ctx context.Context, languages []string) context.Context {
	return context.WithValue(contextOrBackground(ctx), languagesKey, languages)
}

func WithConfigs(ctx context.Context, configs Configs) context.Context {
	return context.WithValue(contextOrBackground(ctx), configsKey, configs)
}

func WithPsm(ctx context.Context, psm string) context.Context {
	return context.WithValue(contextOrBackground(ctx), psmKey, psm)
}

func WithOem(ctx context.Context, oem string) context.Context {
	return context.WithValue(contextOrBackground(ctx), oemKey, oem)
}

func parseContext(ctx context.Context) []string {
	var params []string

	psm := ctx.Value(psmKey)
	if psm != nil {
		params = append(params, "--psm", psm.(string))
	}

	oem := ctx.Value(oemKey)
	if oem != nil {
		params = append(params, "--oem", oem.(string))
	}

	languages := ctx.Value(languagesKey)
	if languages != nil {
		params = append(params, "-l", strings.Join(languages.([]string), "+"))
	}

	configs := ctx.Value(configsKey)
	if configs != nil {
		for key, value := range configs.(Configs) {
			params = append(params, "-c", fmt.Sprintf("%s=%s", key, value))
		}
	}

	return params
}
