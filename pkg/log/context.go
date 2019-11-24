package log

import "context"

type contextKey string

var tagsContextKey = contextKey("logger")

func ContextWithTags(ctx context.Context, tags Tags) context.Context {
	return context.WithValue(ctx, tagsContextKey, tags)
}

func tagsFromContext(ctx context.Context) Tags {
	if ctx == nil {
		return Tags{}
	}
	if tags, _ := ctx.Value(tagsContextKey).(Tags); tags != nil {
		return tags
	}
	return Tags{}
}

func DebugC(ctx context.Context, message string, tags ...Tags) {
	Debug(message, joinTagsFromContext(ctx, tags)...)
}

func InfoC(ctx context.Context, message string, tags ...Tags) {
	Info(message, joinTagsFromContext(ctx, tags)...)
}

func WarnC(ctx context.Context, message string, tags ...Tags) {
	Warn(message, joinTagsFromContext(ctx, tags)...)
}

func ErrorC(ctx context.Context, message string, err error, tags ...Tags) {
	Error(message, err, joinTagsFromContext(ctx, tags)...)
}

func PanicC(ctx context.Context, message string, tags ...Tags) {
	Panic(message, joinTagsFromContext(ctx, tags)...)
}

func PanicWithErrorC(ctx context.Context, message string, err error, tags ...Tags) {
	PanicWithError(message, err, joinTagsFromContext(ctx, tags)...)
}

func ThrowPanicC(ctx context.Context, message string, tags ...Tags) {
	ThrowPanic(message, joinTagsFromContext(ctx, tags)...)
}

func ThrowPanicWithErrorC(ctx context.Context, message string, err error, tags ...Tags) {
	ThrowPanicWithError(message, err, joinTagsFromContext(ctx, tags)...)
}

func copyTags(tags Tags) Tags {
	cp := Tags{}
	for k, v := range tags {
		cp[k] = v
	}
	return cp
}

func joinTagsFromContext(ctx context.Context, tags []Tags) []Tags {
	ctxTags := tagsFromContext(ctx)
	if len(tags) == 0 {
		// Tags added by the logger (like DATA_ERROR or DATA_STACK_TRACE) are written to the first Tags map in the Tags slice
		// This avoids adding tags to the context mutable Tags map
		tags = []Tags{copyTags(ctxTags)}
	} else {
		tags = append(tags, ctxTags)
	}
	return tags
}
