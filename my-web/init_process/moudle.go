package init_process

type Config struct {
	BackDir string `json:"backup_dir"`
	LogDir  string `json:"logdir"`
	MaxFile int    `json:"max_file"`
	Debug   bool
	// MachineAuth []MachineItem `json:"machine_auth"`
	MachineAuth map[string]string `json:"machine_auth"`
}

type MachineItem struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

const (
	XRAYSCRIPTPATH = "/etc/code/conf/xray_doing/ui.sh"
)
