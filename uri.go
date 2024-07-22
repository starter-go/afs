package afs

import "net/url"

// URI 以 string(URI) 的形式表示一个 node
type URI string

func (u URI) String() string {
	return string(u)
}

// URL 把字符串转换为 URL 结构
func (u URI) URL() (*url.URL, error) {
	str := u.String()
	return url.Parse(str)
}
