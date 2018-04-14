package loomplugin

// Command interface represents a CLI command.
type Command interface {
	// Use is the one-line usage message.
	GetUse() string
	// Short is the short description shown in the 'help' output.
	GetShortDesc() string
	// Expected arguments
	CheckArgs(args []string) error
	// RunE: Run but returns an error.
	RunE(args []string) error
	Flags() FlagSet
}

// FlagSet interface represents flags accepted by a CLI command.
type FlagSet interface {
	GetString(name string) (string, error)
	StringP(name, shorthand string, value string, usage string) *string
}

// CmdPluginSystem interface is used by command plugins to hook into the Loom admin CLI.
type CmdPluginSystem interface {
	// GetClient returns a DAppChainClient that can be used to commit txs to a Loom DAppChain.
	GetClient(nodeURI string) (DAppChainClient, error)
}

type CmdPlugin interface {
	Init(sys CmdPluginSystem) error
	GetCmds() []Command
}
