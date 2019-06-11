package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/ypapax/glog3"
	"github.com/ypapax/helper3"
)

func main() {
	glog3.AlsoToStderr(true)
	glog3.SetVerbosity(4)
	if err := func() error {
		cl, err := clipboard.ReadAll()
		if err != nil {
			glog3.Error(err)
			return err
		}
		result, err := ConvertToTestRun(cl)
		if err != nil {
			glog3.Error(err)
			return err
		}
		clipboard.WriteAll(result)
		return nil
	}(); err != nil {
		glog3.Error(err)
		glog3.Flush()
		os.Exit(1)
	}

}

var testNameRegex = regexp.MustCompile(`--- FAIL: (\S+) \(`)

func ConvertToTestRun(inp string) (string, error) {
	lines := strings.Split(inp, "\n")
	var testNames []string
	for _, l := range lines {
		subs := testNameRegex.FindStringSubmatch(l)
		glog3.V(4).Infof("subs %+v", helper3.ToJsonB(subs))
		if len(subs) < 2 {
			continue
		}
		testNames = append(testNames, subs[1])
	}

	return fmt.Sprintf("go test -v -run '%s'", strings.Join(testNames, "|")), nil
}
