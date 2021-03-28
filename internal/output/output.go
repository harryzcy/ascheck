package output

import (
	"os"

	"github.com/harryzcy/ascheck/internal/macapp"
	"github.com/olekukonko/tablewriter"
)

// Table prints application information in table format
func Table(apps []macapp.Application) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Architectures"})

	for _, app := range apps {
		table.Append([]string{app.Name, app.Architectures.String()})
	}

	table.Render()
}
