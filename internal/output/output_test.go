package output

import (
	"strings"
	"testing"

	"github.com/harryzcy/ascheck/internal/localcheck"
	"github.com/harryzcy/ascheck/internal/macapp"
	"github.com/harryzcy/ascheck/internal/remotecheck"
	"github.com/stretchr/testify/assert"
)

var str = new(strings.Builder)

func init() {
	out = str
}

func TestTable(t *testing.T) {
	apps := []macapp.Application{
		{Name: "a", Architectures: localcheck.Architectures{Intel: 0b10}, ArmSupport: remotecheck.SupportNative},
		{Name: "b", Architectures: localcheck.Architectures{Intel: 0b10}, ArmSupport: remotecheck.SupportNative},
	}

	str.Reset()

	Table(apps)

	assert.Equal(t, ""+
		"  NAME  CURRENT ARCHITECTURES  ARM SUPPORT  \n"+
		"--------------------------------------------\n"+
		"  a     Intel 64               Supported    \n"+
		"  b     Intel 64               Supported    \n", str.String())

}
