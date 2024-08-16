package afs

import (
	"fmt"
	"testing"
)

func TestLocation(t *testing.T) {

	list := make([]string, 0)
	list = append(list, "https://user@example.host:6666/a/path/to/file?z=1&y=2&x=3")
	list = append(list, "https://user@中文.example.host:6666/a/也/可/path/to/file?z=壹&y=2&x=3")
	list = append(list, "file:/P:/user中文.example.host:6666/a/也/可/path/to/file?z=壹&y=2&x=3")
	list = append(list, "/绝对路径/a/也/可/path/to/file?z=壹&y=2&x=3")
	list = append(list, "相对路径/a/也/可/path/to/file?z=壹&y=2&x=3")
	list = append(list, "dav://user@host1:999/P:/user中文.example.host/a/也/可/path/to/file?z=壹&y=2&x=3")

	for idx, str := range list {

		fmt.Printf("urls.list[%d]: %s \n", idx, str)

		u1 := URI(str)
		l1, err := ParseLocation(u1)
		if err != nil {
			t.Error(err.Error())
			continue
		}

		u2 := l1.URI()
		l2, err := ParseLocation(u2)
		if err != nil {
			t.Error(err.Error())
			continue
		}

		u3 := l2.URI()

		fmt.Println("  u1 = ", u1)
		fmt.Println("  u2 = ", u2)
		fmt.Println("  u3 = ", u3)
		fmt.Println("  js = ", l2.json())
	}
}
