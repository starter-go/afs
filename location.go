package afs

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// Location 是解析后的 URI 信息
type Location struct {
	Protocol string
	User     string
	Host     string
	Port     int
	Path     string
	Query    map[string]string
}

// URI 函数把这个 Location 转换为 URI
func (inst *Location) URI() URI {
	url := inst.URL()
	str := url.String()
	return URI(str)
}

// URL 函数把这个 Location 转换为 URL
func (inst *Location) URL() *url.URL {

	host := inst.Host
	port := inst.Port
	if port > 0 {
		host = host + ":" + strconv.Itoa(port)
	}

	query := inst.getQueryString()

	dst := &url.URL{
		Scheme:   inst.Protocol,
		User:     nil,
		Host:     host,
		Path:     inst.Path,
		RawQuery: query,
	}

	user := inst.User
	if user != "" {
		dst.User = url.User(user)
	}

	return dst
}

func (inst *Location) getQueryString() string {

	q := inst.Query
	if q == nil {
		return ""
	}
	if len(q) == 0 {
		return ""
	}

	keys := make([]string, 0)
	for key := range q {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	b := &strings.Builder{}
	for _, key := range keys {
		val := q[key]
		if val == "" {
			continue
		}
		if b.Len() > 0 {
			b.WriteRune('&')
		}
		b.WriteString(key)
		b.WriteRune('=')
		b.WriteString(val)
	}
	return b.String()
}

func (inst *Location) String() string {
	url := inst.URL()
	return url.String()
}

func (inst *Location) json() string {
	prefix := ""
	indent := "\t"
	data, err := json.MarshalIndent(inst, prefix, indent)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

// ParseLocation 解析 Location
func ParseLocation(uri URI) (*Location, error) {

	u2, err := uri.URL()
	if err != nil {
		return nil, err
	}

	dst := &Location{
		Protocol: u2.Scheme,
		Host:     u2.Hostname(),
		Path:     u2.Path,
	}

	user := u2.User
	if user != nil {
		dst.User = user.Username()
	}

	portStr := u2.Port()
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("bad URI port: %s", err.Error())
		}
		dst.Port = port
	}

	dst.Query = make(map[string]string)
	q := u2.Query()
	for key, vlist := range q {
		for _, val := range vlist {
			dst.Query[key] = val
		}
	}

	return dst, nil
}
