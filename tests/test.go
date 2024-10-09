package tests

import "testing"

type TestCase struct {
	TestName    string
	Request     Request
	RequestBody interface{}
	Expected    ExpectedResponse
}
type ExpectedResponse struct {
	StatusCode       int
	BodyPart         string
	BodyParts        []string
	BodyPartMissing  string
	BodyPartsMissing []string
	Callback         func(t *testing.T)
}
type Request struct {
	Method string
	Url    string
}
