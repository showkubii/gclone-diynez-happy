// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
package main

import (
	_ "github.com/gclone-diynez-happy/gclone/backend/all" // import all backends
	"github.com/rclone/rclone/cmd"
	_ "github.com/gclone-diynez-happy/gclone/cmd/all"    // import all commands
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
