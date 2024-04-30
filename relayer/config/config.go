package config

type RelayerConfig struct {
	Production  bool   `yaml:"production"`
	RpcUrl      string `yaml:"rpc_url"`
	DaAccountId string `yaml:"da_account_id"`
	KeyPath     string `yaml:"key_path"`
	Network     string `yaml:"network"`
	MetricsAddr string `yaml:"metrics_addr"`
}

func (c RelayerConfig) CompileCMD() []string {
	var cmd []string
	cmd = append(cmd, "args")
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
