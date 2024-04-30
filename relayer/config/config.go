package config

type RelayerConfig struct {
	Production  bool
	RpcUrl      string
	DaAccountId string
	KeyPath     string
	Network     string
	MetricsAddr string
}

func (c RelayerConfig) CompileCMD() []string {
	var cmd []string
	if c.Production {
		cmd = append(cmd, "--production")
	}

	cmd = append(cmd, "--key-path", c.KeyPath)
	cmd = append(cmd, "--rpc-url", c.RpcUrl)
	cmd = append(cmd, "--da-account-id", c.DaAccountId)
	cmd = append(cmd, "--network", c.Network)
	if c.MetricsAddr != "" {
		cmd = append(cmd, "--metrics", c.MetricsAddr)
	}

	return cmd
}
