package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/ahmetb/sorucevap/genpb/sorucevap"
)

func TestProtoToMap(t *testing.T) {
	v := &sorucevap.UserSummary{
		Id:                "foo",
		FullName:          "bar",
		ProfilePictureURL: "quux"}

	expected := map[string]interface{}{
		"id":                "foo",
		"fullName":          "bar",
		"profilePictureURL": "quux"}

	got := ProtoToMap(v)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Fatalf("got wrong map:\n%s", diff)
	}
}

func TestMapToProto(t *testing.T) {
	v := map[string]interface{}{
		"id":                "foo",
		"fullName":          "bar",
		"profilePictureURL": "quux"}
	expected := &sorucevap.UserSummary{
		Id:                "foo",
		FullName:          "bar",
		ProfilePictureURL: "quux"}

	out := &sorucevap.UserSummary{}
	MapToProto(v, out)

	if diff := cmp.Diff(expected, out); diff != "" {
		t.Fatalf("got wrong proto object:\n%s", diff)
	}
}
