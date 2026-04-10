package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/lwmacct/251128-workspace/internal/command/client"
	"github.com/lwmacct/251128-workspace/internal/command/server"
	"github.com/lwmacct/251207-go-pkg-version/pkg/version"
	"github.com/lwmacct/251219-go-pkg-logm/pkg/logm"
	"github.com/urfave/cli/v3"
)

func main() {
	logm.MustInit(logm.PresetAuto())

	cmd := &cli.Command{
		Version: version.AppVersion,
		Flags:   []cli.Flag{},
		Commands: []*cli.Command{
			client.Command,
			server.Command,
			version.Command,
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return cli.ShowSubcommandHelp(c)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("app run", "error", err)
		os.Exit(1)
	}

}
