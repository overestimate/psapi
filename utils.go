package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func getEnv(v string) *string {
	var ret *string
	ret = nil
	for _, env := range os.Environ() {
		if strings.Contains(env, v) {
			t := strings.ReplaceAll(env, v+"=", "")
			ret = &t
		}
	}
	return ret
}

func readFile(p string) (*[]byte, *error) {
	t, e := ioutil.ReadFile(p)
	if e != nil {
		return nil, &e
	}
	return &t, nil
}
