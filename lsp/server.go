package lsp

import (
	"protolsd/util"

	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

const (
	languageServerName    = "ProtoLSD Language Server"
	languageServerVersion = "0.0.1"
)

var (
	srv       *server.Server
	handler   protocol.Handler
	workspace string
)

type lspServer struct {
	srv       *server.Server
	handler   protocol.Handler
	workspace string
}

func Start(workspacePath string, debug bool) {
	if debug {
		commonlog.Configure(2, nil)
	} else {
		commonlog.Configure(1, nil)
	}

	workspace = workspacePath

	handler = protocol.Handler{
		Initialize: initialize,
		Shutdown:   shutdown,
	}

	srv = server.NewServer(&handler, languageServerName, debug)
	srv.RunStdio()
}

func initialize(ctx *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    languageServerName,
			Version: util.Ptr(languageServerVersion),
		},
	}, nil
}

func shutdown(ctx *glsp.Context) error {
	return nil
}
