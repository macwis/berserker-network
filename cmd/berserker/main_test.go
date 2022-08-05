package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := txAddCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--from", "odin", "--to", "odin", "--value=1"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "TX successfully added to the ledger."
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}
