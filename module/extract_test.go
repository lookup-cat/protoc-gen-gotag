package module_test

import (
	"bytes"
	"os"
	"testing"

	"gitea.com/hkrd/protoc-gen-gotag/module"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/spf13/afero"
)

func TestExtract(t *testing.T) {
	req, err := os.Open("../debug/code_generator_request.pb.bin")
	if err != nil {
		t.Fatal(err)
	}

	fs := afero.NewMemMapFs()
	res := &bytes.Buffer{}

	pgs.Init(
		pgs.ProtocInput(req),
		pgs.ProtocOutput(res),
		pgs.FileSystem(fs),
	).RegisterModule(module.New()).Render()
}
