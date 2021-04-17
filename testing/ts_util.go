package testing

import (
	"fmt"
	"io"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// CallDetail print a function call shortly.
func CallDetail(msg []byte, fn interface{}, args ...interface{}) []byte {
	f := runtime.FuncForPC(reflect.ValueOf(fn).Pointer())
	if f != nil {
		msg = append(msg, f.Name()...)
		msg = append(msg, '(')
		msg = argsDetail(msg, args)
		msg = append(msg, ')')
	}
	return msg
}

func argsDetail(b []byte, args []interface{}) []byte {
	nlast := len(args) - 1
	for i, arg := range args {
		b = appendValue(b, arg)
		if i != nlast {
			b = append(b, ',', ' ')
		}
	}
	return b
}

func appendValue(b []byte, arg interface{}) []byte {
	if arg == nil {
		return append(b, "nil"...)
	}
	v := reflect.ValueOf(arg)
	kind := v.Kind()
	if kind >= reflect.Bool && kind <= reflect.Complex128 {
		return append(b, fmt.Sprint(arg)...)
	}
	if kind == reflect.String {
		val := arg.(string)
		if len(val) > 16 {
			val = val[:16] + "..."
		}
		return strconv.AppendQuote(b, val)
	}
	if kind == reflect.Array {
		return append(b, "Array"...)
	}
	if kind == reflect.Struct {
		return append(b, "Struct"...)
	}
	val := v.Pointer()
	b = append(b, '0', 'x')
	return strconv.AppendInt(b, int64(val), 16)
}

// --------------------------------------------------------------------

// Frame represents a program counter inside a stack frame.
// For historical reasons if Frame is interpreted as a uintptr
// its value represents the program counter + 1.
type Frame uintptr

func (f Frame) pc() uintptr { return uintptr(f) - 1 }

func (f Frame) fileLine() (string, int) {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown", 0
	}
	return fn.FileLine(f.pc())
}

// name returns the name of this function, if known.
func (f Frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// Format formats the frame according to the fmt.Formatter interface.
//
//    %n    <funcname>
//    %v    <file>:<line>
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//    %+v   equivalent to <funcname>\n\t<file>:<line>
//
func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		file, line := f.fileLine()
		if s.Flag('+') {
			io.WriteString(s, f.name())
			io.WriteString(s, "\n\t")
			io.WriteString(s, file)
		} else {
			io.WriteString(s, path.Base(file))
		}
		io.WriteString(s, ":")
		io.WriteString(s, strconv.Itoa(line))
	case 'n':
		io.WriteString(s, funcname(f.name()))
	default:
		panic("Frame.Format: unsupport verb - " + string(verb))
	}
}

// --------------------------------------------------------------------

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			fmt.Fprintf(st, "\ngoroutine %d [running]:", len(*s))
			fallthrough
		default:
			for _, pc := range *s {
				f := Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func callers(skip int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
