package cmd

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVersion_Execute(t *testing.T) {
	var buf bytes.Buffer

	cmd := newVersionCmd()
	cmd.SetOut(&buf)
	err := cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	actual := buf.String()

	expected := `{"Version":"given by LDFLAGS","Commit":"given by LDFLAGS","Edited":"given by LDFLAGS","Date":"given by LDFLAGS"}` + "\n"
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Error(diff)
	}
}
