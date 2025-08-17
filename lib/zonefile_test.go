package lib

import (
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

var tmplA *template.Template
var tmplCNAME *template.Template

func TestWriteARecord (t *testing.T) {
	// ionos-vps0.dylt.dev.    1       IN      A       74.208.205.87 ; cf_tags=cf-proxied:false
	data := map[string]string{
		"Name": "ionos-vps-.dylt.dev",
		"Ip": "74.208.205.87",
	}
	tmplA.Execute(os.Stdout, data)
}

func TestReadWrite (t *testing.T) {
	path := "/tmp/hello.txt"
	f, err := os.Create(path)
	require.NoError(t, err)
	f.WriteString("hello")
	err = f.Sync()
	require.NoError(t, err)
	f.Seek(0, 0)
	buf := make([]byte, 0, 32)
	n, err := f.Read(buf)
	require.NoError(t, err)
	t.Logf("n=%d", n)
	t.Logf("buf=%s", string(buf))
}

func init () {
	// ionos-vps0.dylt.dev.    1       IN      A       74.208.205.87 ; cf_tags=cf-proxied:false
	tmplA = template.New("A")
	tmplA.Parse("{{.Name}}\t1\tIN\tA\t{{.Ip}}")

	// etcd0.dylt.dev. 1       IN      CNAME   ionos-vps0.dylt.dev. ; cf_tags=cf-proxied:false
	tmplCNAME = template.New("CNAME")
	tmplCNAME.Parse("");
}
