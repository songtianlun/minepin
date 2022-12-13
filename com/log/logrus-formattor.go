package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"sort"
	"strings"
)

// Modify by "github.com/Lyrics-you/sail-logrus-formatter/sailor"

// Formatter - logrus formatter, implements logrus.Formatter
type Formatter struct {
	// FieldsOrder - default: fields sorted alphabetically
	FieldsOrder []string

	// TimeStampFormat - default: StampNormal = "2006-01-02 15:04:05.000 MST"
	TimeStampFormat string

	// CharStampFormat - default: "" , like "yyyy-MM-dd hh:mm:ss.SSS zzz"
	CharStampFormat string

	// HideKeys - show [fieldValue] instead of [fieldKey:fieldValue]
	HideKeys bool

	// Position - Enable position [file:line @name()]
	Position     bool
	PositionSkip int

	// Colors - Enable colors
	Colors bool

	// FieldsColors - apply colors only to the level, default is level + fields
	FieldsColors bool

	// FieldsSpace - pace between fields
	FieldsSpace bool

	// ShowFullLevel - show a full level [WARNING] instead of [WARN]
	ShowFullLevel bool

	// LowerCaseLevel - no upper case for level value
	LowerCaseLevel bool

	// TrimMessages - trim whitespaces on messages
	TrimMessages bool

	// CallerFirst - print caller info first
	CallerFirst bool

	// CustomCallerFormatter - set custom formatter for caller info
	CustomCallerFormatter func(*runtime.Frame) string
}

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Stamp Handy time stamps.
	Stamp       = "Jan _2 15:04:05"
	StampMilli  = "Jan _2 15:04:05.000"
	StampMicro  = "Jan _2 15:04:05.000000"
	StampNano   = "Jan _2 15:04:05.000000000"
	StampNormal = "2006-01-02 15:04:05.000 MST"
)

type goDate struct {
	year         string
	month        string
	day          string
	hour         string
	minute       string
	second       string
	microseconds string
	zone         string
}

var (
	date = &goDate{
		year:         "2006",
		month:        "01",
		day:          "02",
		hour:         "15",
		minute:       "04",
		second:       "05",
		microseconds: "0000",
		zone:         "MST",
	}
)

// Transform a CharStamp to GoStamp
func (f *Formatter) transformToStamp(timeStampFormat string) string {
	timeStampFormat = strings.Replace(timeStampFormat, "yyyy", date.year, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "YYYY", date.year, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "yy", date.year[2:4], 1)
	timeStampFormat = strings.Replace(timeStampFormat, "YY", date.year[2:4], 1)
	timeStampFormat = strings.Replace(timeStampFormat, "MM", date.month, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "dd", date.day, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "DD", date.day, 1)

	timeStampFormat = strings.Replace(timeStampFormat, "HH", date.hour, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "mm", date.minute, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "ss", date.second, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "S", "0", -1)

	timeStampFormat = strings.Replace(timeStampFormat, "zzz", date.zone, 1)
	timeStampFormat = strings.Replace(timeStampFormat, "ZZZ", date.zone, 1)
	return timeStampFormat
}

// Format an log entry
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)

	if f.TimeStampFormat == "" {
		if f.CharStampFormat != "" {
			f.TimeStampFormat = f.transformToStamp(f.CharStampFormat)
		} else {
			// f.TimeStampFormat = time.StampMilli
			f.TimeStampFormat = StampNormal
		}
	}

	timeStampFormat := f.TimeStampFormat

	// output buffer
	b := &bytes.Buffer{}

	// write time
	b.WriteString("[")
	b.WriteString(entry.Time.Format(timeStampFormat))
	b.WriteString("]")

	// write position
	if f.Position {
		if f.FieldsSpace {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprintf("[%s]", FindCaller(f.PositionSkip)))
	}

	// write level
	var level string
	if f.LowerCaseLevel {
		level = entry.Level.String()
	} else {
		level = strings.ToUpper(entry.Level.String())
	}

	if f.FieldsSpace {
		b.WriteString(" ")
	}

	if f.CallerFirst {
		f.writeCaller(b, entry)
		if f.FieldsSpace {
			b.WriteString(" ")
		}
	}

	if f.Colors {
		fmt.Fprintf(b, "\x1b[%dm", levelColor)
	}

	b.WriteString("[")
	if f.ShowFullLevel {
		b.WriteString(level)
	} else {
		b.WriteString(level[:4])
	}
	b.WriteString("]")

	if f.FieldsSpace {
		b.WriteString(" ")
	}

	if f.Colors && !f.FieldsColors {
		b.WriteString("\x1b[0m")
	}

	// write fields
	if f.FieldsOrder == nil {
		f.writeFields(b, entry)
	} else {
		f.writeOrderedFields(b, entry)
	}

	if f.Colors && f.FieldsColors {
		b.WriteString("\x1b[0m")
	}

	// write message
	if f.TrimMessages {
		b.WriteString(strings.TrimSpace(entry.Message))
	} else {
		b.WriteString(entry.Message)
	}

	if !f.CallerFirst {
		f.writeCaller(b, entry)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *Formatter) writeCaller(b *bytes.Buffer, entry *logrus.Entry) {
	if entry.HasCaller() {
		if f.CustomCallerFormatter != nil {
			fmt.Fprint(b, f.CustomCallerFormatter(entry.Caller))
		} else {
			fmt.Fprintf(
				b,
				"(%s:%d %s)",
				entry.Caller.File,
				entry.Caller.Line,
				entry.Caller.Function,
			)
		}
	}
}

func (f *Formatter) writeFields(b *bytes.Buffer, entry *logrus.Entry) {
	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)

		for _, field := range fields {
			f.writeField(b, entry, field)
		}
	}
}

func (f *Formatter) writeOrderedFields(b *bytes.Buffer, entry *logrus.Entry) {
	length := len(entry.Data)
	foundFieldsMap := map[string]bool{}
	for _, field := range f.FieldsOrder {
		if _, ok := entry.Data[field]; ok {
			foundFieldsMap[field] = true
			length--
			f.writeField(b, entry, field)
		}
	}

	if length > 0 {
		notFoundFields := make([]string, 0, length)
		for field := range entry.Data {
			if !foundFieldsMap[field] {
				notFoundFields = append(notFoundFields, field)
			}
		}

		sort.Strings(notFoundFields)

		for _, field := range notFoundFields {
			f.writeField(b, entry, field)
		}
	}
}

func (f *Formatter) writeField(b *bytes.Buffer, entry *logrus.Entry, field string) {
	if f.HideKeys {
		fmt.Fprintf(b, "[%v]", entry.Data[field])
	} else {
		fmt.Fprintf(b, "[%s:%v]", field, entry.Data[field])
	}

	if f.FieldsSpace {
		b.WriteString(" ")
	}
}

type color uint8

const (
	Black   color = 30
	Red     color = 31
	Green   color = 32
	Yellow  color = 33
	Blue    color = 34
	Magenta color = 35
	Cyan    color = 36
	Gray    color = 37
)

func getColorByLevel(level logrus.Level) color {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return Gray
	case logrus.WarnLevel:
		return Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return Red
	default:
		return Cyan
	}
}
