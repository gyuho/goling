package morphemekor

import (
	"bytes"
	stdlog "log"
	"os"
	"strings"
)

// func main() {
// 	// bs := []byte{}
// 	// b := New().Add("morphemekor")
// 	// for {
// 	// 	n, err := b.ReadByte()
// 	// 	if err != nil {
// 	// 		log.Println(err)
// 	// 		break
// 	// 	}
// 	// 	bs = append(bs, n)
// 	// }
// 	// fmt.Println(bs)
// 	// fmt.Println(string(bs))
// 	fmt.Println(Segment("나는친구한테듣는다안녕"))
// }

// Stream is a stream of bytes.
type Stream struct {
	*bytes.Buffer
}

// New returns *morphemekor.Stream.
func New() *Stream {
	out := new(Stream)
	out.Buffer = new(bytes.Buffer)
	return out
}

// Init initializes Stream.
func (s *Stream) Init() *Stream {
	// *s = *New()
	s.Buffer.Reset()
	return s
}

// Add adds a string to the Stream.
func (s *Stream) Add(str string) *Stream {
	s.Buffer.WriteString(str)
	return s
}

// Get returns the string from the String.
func (s *Stream) Get() string {
	return s.Buffer.String()
}

var log *stdlog.Logger

func init() {
	// import stdlog "log"
	log = stdlog.New(
		os.Stdout,
		"[Log] ",
		stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile,
	)
}

// Segment segments Korean with morphemic approach.
func Segment(str string) string {
	if len(str)/3 < 4 {
		return str
	}
	isInflection1 := make(map[string]bool)
	for _, elem := range strings.Split("이히리기었렀", "") {
		isInflection1[elem] = true
	}
	isTargetEnding1 := make(map[string]bool)
	for _, elem := range strings.Split("다,고,지,마,지만", ",") {
		isTargetEnding1[elem] = true
	}
	isTargetEnding2 := make(map[string]bool)
	for _, elem := range strings.Split("지만", ",") {
		isTargetEnding2[elem] = true
	}
	isPostposition1 := make(map[string]bool)
	for _, elem := range strings.Split("은,이,가,도,을,를,께,로,의,와,과", ",") { // 는,
		isPostposition1[elem] = true
	}
	isPostposition2 := make(map[string]bool)
	for _, elem := range strings.Split("에,한,으", ",") {
		isPostposition2[elem] = true
	}
	isPostposition3 := make(map[string]bool)
	for _, elem := range strings.Split("에게,한테,으로,에서", ",") {
		isPostposition3[elem] = true
	}

	buffer := New()
	wasInflection := false
	doSkip := false
	characters := strings.Split(str, "")

	for idx, elem := range characters {
		if doSkip {
			doSkip = false
			continue
		}
		if idx < 2 {
			buffer.Add(elem)
			doSkip = false
			continue
		}
		if elem == " " {
			buffer.Add(elem)
			doSkip = false
			continue
		}
		if _, ok := isInflection1[elem]; ok {
			wasInflection = true
			doSkip = false
			buffer.Add(elem)
			continue
		}
		if wasInflection {
			if _, ok := isTargetEnding1[elem]; ok {
				buffer.Add(elem)
				log.Println(elem)
				buffer.Add(" ")
				wasInflection = false
				doSkip = false
				continue
			} else {
				if len(characters) > idx+1 {
					if _, ok := isTargetEnding2[elem]; ok {
						buffer.Add(elem)
						log.Println(elem)
						buffer.Add(" ")
						wasInflection = false
						doSkip = false
						continue
					}
				}
			}
		}
		if _, ok := isPostposition1[elem]; ok {
			buffer.Add(elem)
			buffer.Add(" ")
			wasInflection = false
			continue
		} else if _, ok := isPostposition2[elem]; ok {
			if len(characters) > idx+1 {
				char := characters[idx] + characters[idx+1]
				if _, ok := isPostposition3[char]; ok {
					buffer.Add(char)
					log.Println(char)
					buffer.Add(" ")
					wasInflection = false
					doSkip = true
					continue
				}
			}
		}
		wasInflection = false
		doSkip = false
		buffer.Add(elem)
	}
	return buffer.Get()
}
