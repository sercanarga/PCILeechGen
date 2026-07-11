package services

type OutcomeKind uint8

const (
	Completed OutcomeKind = iota + 1
	Error
	Timeout
	Cancelled
)

type Outcome struct {
	Tag  uint8
	Kind OutcomeKind
}

type tagRequest struct {
	tag       uint8
	deadline  uint64
	cancelled bool
}

type TagAllocator struct {
	first       uint8
	count       int
	timeout     uint64
	outstanding map[uint8]tagRequest
	next        int
}

func NewTagAllocator(first uint8, count int, timeout uint64) *TagAllocator {
	if count < 0 {
		count = 0
	}
	if max := 256 - int(first); count > max {
		count = max
	}
	return &TagAllocator{
		first:       first,
		count:       count,
		timeout:     timeout,
		outstanding: make(map[uint8]tagRequest, count),
	}
}

func (a *TagAllocator) Allocate(now uint64) (uint8, bool) {
	if a == nil || a.count == 0 || len(a.outstanding) >= a.count {
		return 0, false
	}
	for offset := range a.count {
		index := (a.next + offset) % a.count
		tag := uint8(int(a.first) + index)
		if _, busy := a.outstanding[tag]; busy {
			continue
		}
		deadline := uint64(0)
		if a.timeout != 0 {
			deadline = now + a.timeout
			if deadline < now {
				deadline = ^uint64(0)
			}
		}
		a.outstanding[tag] = tagRequest{tag: tag, deadline: deadline}
		a.next = (index + 1) % a.count
		return tag, true
	}
	return 0, false
}

func (a *TagAllocator) Complete(tag uint8) (Outcome, bool) {
	return a.finish(tag, Completed)
}

func (a *TagAllocator) Fail(tag uint8) (Outcome, bool) {
	return a.finish(tag, Error)
}

func (a *TagAllocator) finish(tag uint8, kind OutcomeKind) (Outcome, bool) {
	if a == nil {
		return Outcome{}, false
	}
	req, ok := a.outstanding[tag]
	if !ok {
		return Outcome{}, false
	}
	delete(a.outstanding, tag)
	if req.cancelled {
		return Outcome{}, false
	}
	return Outcome{Tag: tag, Kind: kind}, true
}

func (a *TagAllocator) Tick(now uint64) []Outcome {
	if a == nil || a.timeout == 0 {
		return nil
	}
	out := make([]Outcome, 0)
	for index := range a.count {
		tag := uint8(int(a.first) + index)
		req, ok := a.outstanding[tag]
		if ok && req.deadline != 0 && now >= req.deadline {
			delete(a.outstanding, tag)
			if !req.cancelled {
				out = append(out, Outcome{Tag: tag, Kind: Timeout})
			}
		}
	}
	return out
}

func (a *TagAllocator) CancelAll() []Outcome {
	if a == nil {
		return nil
	}
	out := make([]Outcome, 0, len(a.outstanding))
	for index := range a.count {
		tag := uint8(int(a.first) + index)
		req, ok := a.outstanding[tag]
		if ok && !req.cancelled {
			req.cancelled = true
			a.outstanding[tag] = req
			out = append(out, Outcome{Tag: tag, Kind: Cancelled})
		}
	}
	return out
}

func (a *TagAllocator) Outstanding() int {
	if a == nil {
		return 0
	}
	return len(a.outstanding)
}
