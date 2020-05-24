package cmd

import (
	_ "github.com/chirino/graphql-gw/internal/cmd/completion"
	_ "github.com/chirino/graphql-gw/internal/cmd/config/add/upstream"
	_ "github.com/chirino/graphql-gw/internal/cmd/config/mount"
	_ "github.com/chirino/graphql-gw/internal/cmd/new"
	"github.com/chirino/graphql-gw/internal/cmd/root"
	_ "github.com/chirino/graphql-gw/internal/cmd/serve"
	"github.com/chirino/graphql-gw/internal/cmd/version"
)

type VersionConfig = version.VersionConfig

func Main(versionConfig VersionConfig) {
	version.Config = versionConfig
	root.Main()
}
