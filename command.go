package buildInfo

import (
	"fmt"
	"github.com/mitchellh/cli"
)

// CommandT is a Command implementation that returns version information
type CommandT struct {
	bi *BuildInfo
	ui cli.Ui
}

// Command Builds and returns a CommandT struct
func (bi *BuildInfo) Command(ui cli.Ui) (*CommandT, error) {
	return &CommandT{
		bi: bi,
		ui: ui,
	}, nil
}

// Run is a function to run the command
func (c *CommandT) Run(args []string) int {
	// show output
	return c.show()
}

// Synopsis shows the command summary
func (c *CommandT) Synopsis() string {
	return "Display application version information"
}

// Help shows the detailed command options
func (c *CommandT) Help() string {
	return fmt.Sprintf(`Usage: %s version [options]

	Display application version and dependency information.
`, c.bi.Name)
}

// print version information
func (c *CommandT) show() int {
	// print standard version
	c.ui.Output(fmt.Sprintf("==>\t%s v%s\nBuild:\t%s\nBranch:\t%s\nDate:\t%s",
		c.bi.Name,
		c.bi.Version,
		c.bi.Build,
		c.bi.Branch,
		c.bi.Date))

	// all good
	return 0
}
