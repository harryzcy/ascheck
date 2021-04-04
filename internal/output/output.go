package output

import (
	"encoding/json"
	"io"
	"os"

	"github.com/harryzcy/ascheck/internal/macapp"
	"github.com/olekukonko/tablewriter"
)

var out io.Writer = os.Stdout

// Table prints application information in table format.
func Table(apps []macapp.Application) {
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Name", "Current Architectures", "Arm Support"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetBorder(false)

	for _, app := range apps {
		table.Append([]string{app.Name, app.Architectures.String(), app.ArmSupport.String()})
	}

	table.Render()
}

// JSON prints application information in json format.
func JSON(apps []macapp.Application) {
	items := make([]map[string]string, len(apps))

	for idx, app := range apps {
		row := map[string]string{
			"name":                 app.Name,
			"currentArchitectures": app.Architectures.String(),
			"armSupport":           app.ArmSupport.String(),
		}
		items[idx] = row
	}

	output, _ := json.Marshal(map[string]interface{}{
		"items": items,
	})
	output = append(output, byte('\n'))
	out.Write(output)
}
