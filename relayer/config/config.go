package config

type RelayerConfig struct {
	Production        bool   `yaml:"production"`
	RpcUrl            string `yaml:"rpc_url"`
	DaAccountId       string `yaml:"da_account_id"`
	KeyPath           string `yaml:"key_path"`
	Network           string `yaml:"network"`
	MetricsIpPortAddr string `yaml:"metrics_ip_port_address"`
}

func (c RelayerConfig) CompileCMD() []string {
	var cmd []string
	cmd = append(cmd, "run-args")
	if c.Production {
		cmd = append(cmd, "--production")
	}

	cmd = append(cmd, "--key-path", c.KeyPath)
	cmd = append(cmd, "--rpc-url", c.RpcUrl)
	cmd = append(cmd, "--da-account-id", c.DaAccountId)
	cmd = append(cmd, "--network", c.Network)
	if c.MetricsIpPortAddr != "" {
		cmd = append(cmd, "--metrics-ip-port-address", c.MetricsIpPortAddr)
	}

	return cmd
}
