package dns

type DnsOption struct {
	Name    string
	IPs     []string
	Address string
}

var DnsOptions = []DnsOption{
	{
		Name:    "shecan",
		IPs:     []string{"178.22.122.100", "185.51.200.2"},
		Address: "https://shecan.ir/",
	},
	{
		Name:    "403",
		IPs:     []string{"10.202.10.202", "10.202.10.102"},
		Address: "https://403.online/",
	},
	{
		Name:    "server",
		IPs:     []string{"194.104.158.48", "194.104.158.78"},
		Address: "https://server.ir",
	},
	{
		Name:    "hostiran",
		IPs:     []string{"172.29.0.100", "172.29.2.100"},
		Address: "https://hostiran.net/landing/proxy",
	},
	{
		Name:    "begzar",
		IPs:     []string{"185.55.226.26", "185.55.225.25"},
		Address: "https://begzar.ir/",
	},
	{
		Name:    "electro",
		IPs:     []string{"78.157.42.100", "78.157.42.101"},
		Address: "https://electrotm.org/",
	},
	{
		Name:    "radar",
		IPs:     []string{"10.202.10.10", "10.202.10.11"},
		Address: "https://radar.game/#/dns",
	},
	{
		Name:    "asiatech",
		IPs:     []string{"194.36.174.161", "178.22.122.100"},
		Address: "https://asiatech.cloud/",
	},
	{
		Name:    "asrenovin",
		IPs:     []string{"46.224.1.42", "178.22.122.100"},
		Address: "",
	},
	{
		Name:    "tums",
		IPs:     []string{"194.225.62.80", "178.22.122.100"},
		Address: "",
	},
	{
		Name:    "pishgaman",
		IPs:     []string{"5.202.100.101", "178.22.122.100"},
		Address: "",
	},
	{
		Name:    "parsonline",
		IPs:     []string{"91.99.101.12", "178.22.122.100"},
		Address: "",
	},
	{
		Name:    "cloud",
		IPs:     []string{"1.1.1.1", "1.0.0.1"},
		Address: "https://www.cloudflare.com/",
	},
	{
		Name:    "google",
		IPs:     []string{"8.8.8.8", "8.8.4.4"},
		Address: "https://google.com",
	},
}
