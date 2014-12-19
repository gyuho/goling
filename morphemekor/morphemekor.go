package morphemekor // import "github.com/gyuho/goling/morphemekor"

import (
	"bytes"
	"strings"
)

// bs := []byte{}
// b := New().Add("morphemekor")
// for {
// 	n, err := b.ReadByte()
// 	if err != nil {
// 		log.Println(err)
// 		break
// 	}
// 	bs = append(bs, n)
// }
// fmt.Println(bs)
// fmt.Println(string(bs))

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

// Segment segments Korean with morphemic approach.
// (Reference: https://github.com/gyuho/learn_other/blob/master/ling105/gyuho.pdf)
func Segment(str string) string {
	if len(str)/3 < 4 {
		return str
	}
	isInflection1 := make(map[string]bool)
	for _, elem := range strings.Split("이히리기었렀", "") {
		isInflection1[elem] = true
	}
	isTargetEnding1 := make(map[string]bool)
	for _, elem := range strings.Split("다,고,지,마,지", ",") {
		isTargetEnding1[elem] = true
	}
	isTargetEnding2 := make(map[string]bool)
	for _, elem := range strings.Split("지만", ",") {
		isTargetEnding2[elem] = true
	}
	isPostposition1 := make(map[string]bool)
	for _, elem := range strings.Split("은,는,이,가,도,을,를,께,로,의,와,과", ",") {
		isPostposition1[elem] = true
	}
	isPostposition2 := make(map[string]bool)
	for _, elem := range strings.Split("에,한,으", ",") {
		isPostposition2[elem] = true
	}
	isPostposition3 := make(map[string]bool)
	for _, elem := range strings.Split("에게,에서,한테,으로", ",") {
		isPostposition3[elem] = true
	}
	buffer := New()
	characters := strings.Split(str, "")
	marker := 0 // another index in word phrase
	for idx := 0; idx < len(characters); idx++ {
		elem := characters[idx]
		if idx == 0 || marker == 0 {
			buffer.Add(elem)
			marker++
			continue
		}
		if elem == " " {
			buffer.Add(elem)
			marker++
			continue
		}
		// Check if the character is inflection
		if _, ok := isInflection1[elem]; ok {
			buffer.Add(elem)
			// Check the next chracter
			if len(characters) > idx+1 {
				nelem1 := characters[idx+1]
				if _, ok := isTargetEnding1[nelem1]; ok {
					// Check the second next chracter
					if len(characters) > idx+2 {
						nelem2 := characters[idx+2]
						if _, ok := isTargetEnding2[nelem1+nelem2]; ok {
							buffer.Add(nelem1 + nelem2)
							buffer.Add(" ")
							idx = idx + 2
							marker = 0
							continue
						}
					}
					buffer.Add(nelem1)
					buffer.Add(" ")
					idx++
					marker = 0
					continue
				}
			}
			marker++
			continue
		}
		// Check if the character is postposition
		if _, ok := isPostposition1[elem]; ok {
			buffer.Add(elem)
			buffer.Add(" ")
			marker = 0
			continue
		} else if _, ok := isPostposition2[elem]; ok {
			if len(characters) > idx+1 {
				nelem2 := characters[idx+1]
				if _, ok := isPostposition3[elem+nelem2]; ok {
					buffer.Add(elem + nelem2)
					buffer.Add(" ")
					idx++
					marker = 0
					continue
				}
			}
		}
		buffer.Add(elem)
		marker++
	}
	return strings.TrimSpace(buffer.Get())
}
