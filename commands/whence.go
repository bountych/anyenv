package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
)

func whenceCmd(args cli.Args) {
	exeName := args.List[0]
	var exeFile utils.Pathname

	versionsDir := config.VersionsDir()
	versionPaths := versionsDir.Entries()

	for _, versionPath := range versionPaths {
		exeFile = versionPath.Join("bin", exeName)
		if exeFile.IsExecutable() {
			cli.Println(versionPath.Base())
		}
	}
}

func init() {
	cli.Register("whence", whenceCmd)
}
