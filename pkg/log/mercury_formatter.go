package log

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
)

const defaultTimestampFormat = time.RFC3339Nano

const dataKeyPrefix = "__DATA_"

const (
	FieldKeyLevel = "level"
	FieldKeyTime  = "time"
)

type FieldMap map[string]string

func (f FieldMap) resolve(key string) string {
	if k, ok := f[key]; ok {
		return k
	}
	return key
}

type MercuryFormatter struct {
	DisableTimestamp bool
	TimestampFormat  string
	FieldMap         FieldMap
}

func (f *MercuryFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	tags := make(logrus.Fields, len(entry.Data)+1)
	data := map[string]interface{}{}
	for k, v := range entry.Data {
		if strings.HasPrefix(k, dataKeyPrefix) {
			data[k[len(dataKeyPrefix):]] = v
		} else {
			tags[k] = v
		}
	}
	prefixFieldClashes(tags, f.FieldMap)
	if !f.DisableTimestamp {
		appendTimestampTag(f, tags, entry)
	}
	tags[f.FieldMap.resolve(FieldKeyLevel)] = entry.Level.String()
	var buf *bytes.Buffer
	if entry.Buffer != nil {
		buf = entry.Buffer
	} else {
		buf = &bytes.Buffer{}
	}
	message := strings.TrimSpace(entry.Message)
	if message != "" {
		appendString(buf, message)
		buf.WriteByte(' ')
	}
	for k, v := range tags {
		appendTag(buf, k, v)
	}
	for k, v := range data {
		appendData(buf, k, v)
	}
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

func appendTimestampTag(f *MercuryFormatter, tags logrus.Fields, entry *logrus.Entry) {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	tags[f.FieldMap.resolve(FieldKeyTime)] = entry.Time.Format(timestampFormat)
}

func prefixFieldClashes(data logrus.Fields, fieldMap FieldMap) {
	timeKey := fieldMap.resolve(FieldKeyTime)
	if t, ok := data[timeKey]; ok {
		data["tags."+timeKey] = t
		delete(data, timeKey)
	}
	levelKey := fieldMap.resolve(FieldKeyLevel)
	if l, ok := data[levelKey]; ok {
		data["tags."+levelKey] = l
		delete(data, levelKey)
	}
}

func appendTag(buf *bytes.Buffer, key string, value interface{}) {
	buf.WriteByte('[')
	appendStringWithoutControls(buf, key)
	buf.WriteByte(':')
	switch value := value.(type) {
	case error:
		appendStringWithoutControls(buf, value.Error())
	default:
		appendStringWithoutControls(buf, fmt.Sprintf("%+v", value))
	}
	buf.WriteByte(']')
}

func appendData(buf *bytes.Buffer, key string, value interface{}) {
	buf.WriteByte(' ')
	appendStringWithoutControls(buf, key)
	buf.WriteByte(':')
	buf.WriteByte(' ')
	appendString(buf, fmt.Sprintf("%+v", value))
}

func appendString(buf *bytes.Buffer, s string) {
	for i := 0; i < len(s); i++ {
		appendChar(buf, s[i])
	}
}

func appendStringWithoutControls(buf *bytes.Buffer, s string) {
	for i := 0; i < len(s); i++ {
		appendCharWithoutControls(buf, s[i])
	}
}

func appendChar(buf *bytes.Buffer, c byte) {
	if c >= utf8.RuneSelf {
		return
	}
	switch c {
	case '\\', '"':
		buf.WriteByte('\\')
		buf.WriteByte(c)
	case '\r':
	case '\n':
		buf.WriteByte('\n')
		buf.WriteByte('\t')
	default:
		buf.WriteByte(c)
	}
}

func appendCharWithoutControls(buf *bytes.Buffer, c byte) {
	if c >= utf8.RuneSelf {
		return
	}
	switch c {
	case '\\', '"':
		buf.WriteByte('\\')
		buf.WriteByte(c)
	case '\t':
		buf.WriteByte('\\')
		buf.WriteByte('\t')
	case '\r':
		buf.WriteByte('\\')
		buf.WriteByte('\r')
	case '\n':
		buf.WriteByte('\\')
		buf.WriteByte('\n')
	default:
		buf.WriteByte(c)
	}
}

func DataKey(k string) string {
	return dataKeyPrefix + strings.ToUpper(k)
}
