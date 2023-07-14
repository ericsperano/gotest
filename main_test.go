package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const input = `?   	github.com/ericsperano/yfh	[no test files]
?	github.com/ericsperano/yfh/apiserver	[no test files]
?   	github.com/ericsperano/yfh/cmd	[no test files]
?   	github.com/ericsperano/yfh/docs	[no test files]
ok  	github.com/ericsperano/yfh/core	0.327s	coverage: 28.6% of statements
ok  	github.com/ericsperano/yfh/core/model	(cached)	coverage: 93.3% of statements
ok  	github.com/ericsperano/yfh/core/xmlmodel	(cached)	coverage: 96.7% of statements
?   	github.com/ericsperano/yfh/graph/generated	[no test files]
?   	github.com/ericsperano/yfh/graph/model	[no test files]
?   	github.com/ericsperano/yfh/worker	[no test files]
ok  	github.com/ericsperano/yfh/graph	(cached)	coverage: 0.2% of statements [no tests to run]
`

const output = "?   github.com/ericsperano/yfh                                        \n" +
	"?   github.com/ericsperano/yfh/apiserver                              \n" +
	"?   github.com/ericsperano/yfh/cmd                                    \n" +
	"?   github.com/ericsperano/yfh/docs                                   \n" +
	"ok  github.com/ericsperano/yfh/core                              28.6%\n" +
	"ok  github.com/ericsperano/yfh/core/model                        93.3%\n" +
	"ok  github.com/ericsperano/yfh/core/xmlmodel                     96.7%\n" +
	"?   github.com/ericsperano/yfh/graph/generated                        \n" +
	"?   github.com/ericsperano/yfh/graph/model                            \n" +
	"?   github.com/ericsperano/yfh/worker                                 \n" +
	"ok  github.com/ericsperano/yfh/graph                              0.2%\n"

func getInputLines() []string {
	return strings.Split(input, "\n")
}

func getOutputLines() []string {
	return strings.Split(output, "\n")
}

func longest(strs []string) int {
	max := 0
	for _, str := range strs {
		if len(str) > max {
			max = len(str)
		}
	}
	return max
}

func TestGetOutput(t *testing.T) {
	ilines := getInputLines()
	assert.Equal(t, 12, len(ilines))
	olines := getOutputLines()
	assert.Equal(t, len(ilines), len(olines))
	for i, line := range ilines {
		line = convertLine(line)
		assert.Equal(t, olines[i], line)
	}
}
