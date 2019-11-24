package log

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
)

type Tags logrus.Fields

var log *logrus.Logger

func init() {
	log = &logrus.Logger{
		Out:   os.Stdout,
		Hooks: make(logrus.LevelHooks),
		Formatter: &MercuryFormatter{
			DisableTimestamp: false,
			TimestampFormat:  "",
			FieldMap:         FieldMap{},
		},
		ReportCaller: false,
		Level:        logrus.DebugLevel,
		ExitFunc:     nil,
	}
}

func SetLevel(level string) {
	if lvl, err := logrus.ParseLevel(level); err != nil {
		panic(err)
	} else {
		log.SetLevel(lvl)
	}
}

func Debug(message string, tags ...Tags) {
	if log.Level >= logrus.DebugLevel {
		log.WithFields(unifyTagsToFields(tags)).Debug(message)
	}
}

func Info(message string, tags ...Tags) {
	if log.Level >= logrus.InfoLevel {
		log.WithFields(unifyTagsToFields(tags)).Info(message)
	}
}

func Warn(message string, tags ...Tags) {
	if log.Level >= logrus.WarnLevel {
		log.WithFields(unifyTagsToFields(tags)).Warn(message)
	}
}

func Error(message string, err error, tags ...Tags) {
	if log.Level >= logrus.ErrorLevel {
		tags = appendTaggableErrorTags(err, tags)
		log.WithFields(addErrorField(unifyTagsToFields(tags), err)).Error(message)
	}
}

func Panic(message string, tags ...Tags) {
	log.WithFields(addStackField(unifyTagsToFields(tags))).Error(message)
}

func PanicWithError(message string, err error, tags ...Tags) {
	tags = appendTaggableErrorTags(err, tags)
	log.WithFields(addStackField(addErrorField(unifyTagsToFields(tags), err))).Error(message)
}

func ThrowPanic(message string, tags ...Tags) {
	log.WithFields(addStackField(unifyTagsToFields(tags))).Panic(message)
}

func ThrowPanicWithError(message string, err error, tags ...Tags) {
	tags = appendTaggableErrorTags(err, tags)
	log.WithFields(addStackField(addErrorField(unifyTagsToFields(tags), err))).Panic(message)
}

func unifyTagsToFields(tagsSlice []Tags) logrus.Fields {
	if len(tagsSlice) == 0 {
		return logrus.Fields{}
	}
	for i := 1; i < len(tagsSlice); i++ {
		for k, v := range tagsSlice[i] {
			tagsSlice[0][k] = v
		}
	}
	return logrus.Fields(tagsSlice[0])
}

func addErrorField(fields logrus.Fields, err error) logrus.Fields {
	if err != nil {
		fields[dataKeyPrefix+"ERROR"] = err
	}
	return fields
}

func addStackField(fields logrus.Fields) logrus.Fields {
	fields[dataKeyPrefix+"STACK_TRACE"] = string(debug.Stack())
	return fields
}

func appendTaggableErrorTags(err error, tags []Tags) []Tags {
	if taggable, ok := errors.Cause(err).(Taggable); ok {
		tags = append(tags, taggable.Tags())
	}
	return tags
}
