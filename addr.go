package tunnel

import (
	"errors"
	"fmt"
	"regexp"
)

type AddrType byte

const (
	IP AddrType = iota
	PORT
	URL
	DIR
)

type Address struct {
	addr     string
	addrType AddrType
}

func isPort(a string) bool {
	m, _ := regexp.MatchString("^[0-9]+$", a)
	return m
}

func isIP(a string) bool {
	m, _ := regexp.MatchString("(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}", a)
	return m
}

func isUrl(a string) bool {
	m, _ := regexp.MatchString("^(?:https?://)([^:\\/\\n]+)", a)
	return m
}

func isDir(a string) bool {
	m, _ := regexp.MatchString("^(?i)(dir://)(.*)", a)
	return m
}

func NewAddress(a string) (*Address, error) {
	addr := &Address{
		addr: a,
	}

	if isIP(a) {
		addr.addrType = IP
	} else if isPort(a) {
		addr.addrType = PORT
	} else if isUrl(a) {
		addr.addrType = URL
	} else if isDir(a) {
		addr.addrType = DIR
	} else {
		return nil, errors.New(fmt.Sprintf("unknown address type: %s", a))
	}

	return addr, nil
}

func (a Address) Type() AddrType {
	return a.addrType
}

func (a Address) ToProtoString() string {
	addr := a.ToString()
	if a.addrType == DIR || a.addrType == URL {
		return addr
	}

	m, _ := regexp.MatchString("^(?:https?://)", addr)
	if !m {
		return "http://" + addr
	}

	return a.addr
}

func (a Address) ToString() string {
	if a.addrType == PORT {
		return "127.0.0.1:" + a.addr
	}
	if a.addrType == DIR {
		return a.addr[5:]
	}

	return a.addr
}
