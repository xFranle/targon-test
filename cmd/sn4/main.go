package main

import (
	_ "github.com/manifold-inc/targon/internal/cli/attest"
	_ "github.com/manifold-inc/targon/internal/cli/config"
	_ "github.com/manifold-inc/targon/internal/cli/get"
	"github.com/manifold-inc/targon/internal/cli/root"
	_ "github.com/manifold-inc/targon/internal/cli/vali"
)

func main() {
	root.Execute()
}
