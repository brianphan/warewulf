package list

import (
	"bytes"
	"path"
	"testing"

	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/stretchr/testify/assert"

	"github.com/hpcng/warewulf/internal/pkg/testenv"
	"github.com/hpcng/warewulf/internal/pkg/warewulfd"
)

func Test_Overlay_List(t *testing.T) {
	env := testenv.New(t)
	env.WriteFile(t, path.Join(testenv.WWOverlaydir, "testoverlay/email.ww"), `
{{ if .Tags.email }}eMail: {{ .Tags.email }}{{else}} noMail{{- end }}
`)
	defer env.RemoveAll(t)
	warewulfd.SetNoDaemon()
	t.Run("overlay list", func(t *testing.T) {
		baseCmd := GetCommand()
		buf := new(bytes.Buffer)
		baseCmd.SetOut(buf)
		baseCmd.SetErr(buf)
		wwlog.SetLogWriter(buf)
		err := baseCmd.Execute()
		assert.NoError(t, err)
		assert.Contains(t, buf.String(), "testoverlay")
	})
	t.Run("overlay list all", func(t *testing.T) {
		baseCmd.SetArgs([]string{"-a"})
		baseCmd := GetCommand()
		buf := new(bytes.Buffer)
		baseCmd.SetOut(buf)
		baseCmd.SetErr(buf)
		wwlog.SetLogWriter(buf)
		err := baseCmd.Execute()
		assert.NoError(t, err)
		assert.Contains(t, buf.String(), "email.ww")
	})
}
