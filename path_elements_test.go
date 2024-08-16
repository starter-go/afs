package afs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestElements(t *testing.T) {

	list := make([]string, 0)
	list = append(list, "")
	list = append(list, "/a//b/c")
	list = append(list, "/a/./b/c")
	list = append(list, "/a/../b/c")
	list = append(list, "/a/~/b/c")

	pe := &PathElements{}

	for _, item := range list {
		//////////////////////////////////
		fmt.Println("test elements, path = ", item)
		elist1 := pe.Split(item)
		elist2, err := pe.Resolve(elist1)
		if err != nil {
			t.Error(err.Error())
			break
		}
		path := pe.Stringify(elist2, "", '/')
		fmt.Println("         final path = ", path)
		//////////////////////////////////

		js, _ := json.Marshal(elist2)
		fmt.Println("           elements = ", string(js))
		fmt.Println("")
	}
}
