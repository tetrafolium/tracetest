package pkg1

import (
	"gopkg.in/errgo.v1"
)

func pkg() string {
	var name string
	if err := json.Unmarshal([]byte("pkg1"), &name); err != nil {
		panic(errgo.Notef(err, "failed to unmarshal"))
	}
	return name
}

func hello() {
	name := pkg()
	fmt.Println(strings.Join([]string{"hello", name}, " "))
}
