package param

import "flag"

var TestDir *string

func init() {
	TestDir = flag.String("testdir", "", "with test the server returns mock results of given directory")
	flag.Parse()
}
