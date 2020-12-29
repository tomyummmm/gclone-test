package main

import (
	_ "github.com/tomyummmm/gclone-test/backend/all" // import all backends
	"github.com/rclone/rclone/cmd"
	_ "github.com/tomyummmm/gclone-test/cmd/copy"
	_ "github.com/rclone/rclone/cmd/all"    // import all commands
	"github.com/rclone/rclone/fs"
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	fs.Version = fs.Version+"-mod1.3.1"
	cmd.Main()
}
