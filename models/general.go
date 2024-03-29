package models

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type GeneralModel struct {
	ParentSpan opentracing.Span
	OttoZaplog *zap.Logger
	SpanId     string
	Context    context.Context
}
