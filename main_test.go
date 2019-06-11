package main

import (
	"os"
	"testing"

	"github.com/ypapax/glog3"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	glog3.AlsoToStderr(true)
	glog3.SetVerbosity(4)
	ret := m.Run()
	os.Exit(ret)

}

func TestConvertToTestRun(t *testing.T) {
	input := `--- FAIL: TestOne (11.45s)
--- FAIL: TestTwo (22.89s)
--- FAIL: TestThree (11.15s)
`
	as := assert.New(t)
	result, err := ConvertToTestRun(input)
	if !as.NoError(err) {
		return
	}

	expected := `go run -v -run TestOne|TestTwo|TestThree`
	as.Equal(expected, result)
}
