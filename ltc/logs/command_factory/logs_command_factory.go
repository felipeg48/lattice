package command_factory

import (
	"github.com/cloudfoundry-incubator/lattice/ltc/exit_handler"
	"github.com/cloudfoundry-incubator/lattice/ltc/logs/console_tailed_logs_outputter"
	"github.com/cloudfoundry-incubator/lattice/ltc/output"
	"github.com/codegangsta/cli"
    "github.com/cloudfoundry-incubator/lattice/ltc/logs/reserved_app_ids"
)

type logsCommandFactory struct {
	cmd *logsCommand
}

func NewLogsCommandFactory(output *output.Output, tailedLogsOutputter console_tailed_logs_outputter.TailedLogsOutputter, exitHandler exit_handler.ExitHandler) *logsCommandFactory {
	return &logsCommandFactory{
		&logsCommand{
			output:              output,
			tailedLogsOutputter: tailedLogsOutputter,
			exitHandler:         exitHandler,
		},
	}
}

func (factory *logsCommandFactory) MakeLogsCommand() cli.Command {
	var logsCommand = cli.Command{
		Name:        "logs",
		ShortName:   "l",
		Description: "Stream logs from the specified application",
		Usage:       "ltc logs APP_NAME",
		Action:      factory.cmd.tailLogs,
		Flags:       []cli.Flag{},
	}

	return logsCommand
}

func (factory *logsCommandFactory) MakeDebugLogsCommand() cli.Command{
    return cli.Command{
        Name: "debug-logs",
        Description: "Stream logs from the executor, rep, and garden-linux lattice components",
        Usage: "ltc debug-logs",
        Action: factory.cmd.tailDebugLogs,
    }
}

type logsCommand struct {
	output              *output.Output
	tailedLogsOutputter console_tailed_logs_outputter.TailedLogsOutputter
	exitHandler         exit_handler.ExitHandler
}

func (cmd *logsCommand) tailLogs(context *cli.Context) {
	appGuid := context.Args().First()

	if appGuid == "" {
		cmd.output.IncorrectUsage("")
		return
	}

	cmd.tailedLogsOutputter.OutputTailedLogs(appGuid)
}

func (cmd *logsCommand) tailDebugLogs(context *cli.Context) {
	cmd.tailedLogsOutputter.OutputTailedLogs(reserved_app_ids.LatticeDebugLogStreamAppId)
}
