package init_process

type Config struct {
	BackDir string `json:"backup_dir"`
	LogDir  string `json:"logdir"`
	MaxFile int    `json:"max_file"`
	Debug   bool
}
