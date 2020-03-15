package segment

import "strings"

// Segment ...
type Segment struct {
	path string
	size int
	pos  int
}

// NewSegment ...
func NewSegment(path, toIgnore string) Segment {
	if path == "" {
		path = "/"
	}
	path = strings.Replace(path, toIgnore, "", -1)
	return Segment{
		path: path,
		size: len(path),
	}
}

// Current ...
func (s Segment) Current() string {
	return s.path[s.pos:s.size]
}

// Previous ...
func (s Segment) Previous() string {
	return s.path[0:s.pos]
}

// Root ...
func (s Segment) Root() bool {
	return s.size <= s.pos+1
}

// Init ...
func (s Segment) Init() bool {
	return 0 == s.pos
}

// Extract ...
func (s *Segment) Extract() string {
	if s.Root() {
		return ""
	}
	offs := s.pos + 1
	s.pos = s.index(offs)
	return s.path[offs:s.pos]
}

// Retract ...
func (s *Segment) Retract() string {
	if s.Init() {
		return ""
	}
	offs := s.pos
	s.pos = s.lastIndex()
	return s.path[s.pos+1 : offs]
}

// Consume ...
func (s *Segment) Consume(value string) bool {
	origin := s.pos
	if value == s.Extract() {
		return true
	}
	s.pos = origin
	return false
}

// Restore ...
func (s *Segment) Restore(value string) bool {
	origin := s.pos
	if value == s.Retract() {
		return true
	}
	s.pos = origin
	return false
}

// Capture ...
func (s *Segment) Capture(key string, store map[string]string) {
	value := s.Extract()
	if value != "" {
		store[key] = value
	}
	return
}

func (s Segment) index(start int) (index int) {
	for index = start; index <= s.size; index++ {
		if index == s.size {
			break
		}
		if string(s.path[index]) == "/" {
			break
		}
	}
	return
}

func (s Segment) lastIndex() (index int) {
	for index = s.pos - 1; index >= 0; index-- {
		if index == 0 {
			break
		}
		if string(s.path[index]) == "/" {
			break
		}
	}
	return
}
