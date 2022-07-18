package tests

import "github.com/revel/revel/testing"

type BasicHttpTest struct {
	testing.TestSuite
}

func (t *BasicHttpTest) TestNotFoundEndPoint() {
	t.Get("/it-will-not-be-found")
	t.AssertNotFound()
	t.AssertContentType("application/json; charset=utf-8")
	t.AssertContainsRegex("\"code\".*404")
	t.AssertContainsRegex("\"error\".*\"EP_NOT_FOUND\"")
}
