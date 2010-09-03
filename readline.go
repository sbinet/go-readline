// Wrapper around the GNU readline(3) library

package readline

// TODO:
//  implement a go-oriented command completion

// #include <stdio.h>
// #include <stdlib.h>
// #include <readline/readline.h>
// #include <readline/history.h>
import "C"
import "unsafe"
import "os"

func ReadLine(prompt *string) *string {
	var p *C.char;

	//readline allows an empty prompt(NULL)
	if prompt != nil { p = C.CString(*prompt) }

	ret := C.readline(p);

	if p != nil { C.free(unsafe.Pointer(p)) }

	if ret == nil { return nil } //EOF

	s := C.GoString(ret);
	C.free(unsafe.Pointer(ret));
	return &s
}

func AddHistory(s string) {
	p := C.CString(s);
	defer C.free(unsafe.Pointer(p))
	C.add_history(p);	
}

// Parse and execute single line of a readline init file.
func ParseAndBind(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.rl_parse_and_bind(p)
}

// Parse a readline initialization file.
// The default filename is the last filename used.
func ReadInitFile(s string) os.Errno {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.rl_read_init_file(p)
	return os.Errno(errno)
}

// Load a readline history file.
// The default filename is ~/.history.
func ReadHistoryFile(s string) os.Errno {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.read_history(p)
	return os.Errno(errno)
}

var (
	HistoryLength = -1
)

// Save a readline history file.
// The default filename is ~/.history.
func WriteHistoryFile(s string) os.Errno {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.write_history(p)
	if errno==0 && HistoryLength >= 0 {
		errno = C.history_truncate_file(p, C.int(HistoryLength))
	}
	return os.Errno(errno)
}

// Set the readline word delimiters for tab-completion
func SetCompleterDelims(break_chars string) {
	p := C.CString(break_chars)
	//defer C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(C.rl_completer_word_break_characters))
	C.rl_completer_word_break_characters = p
}

// Get the readline word delimiters for tab-completion
func GetCompleterDelims() string {
	cstr := C.rl_completer_word_break_characters
	delims := C.GoString(cstr)
	return delims
}
/* EOF */
