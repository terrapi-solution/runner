package common

type ManagerConfig struct {
	Address string `json:"address"`
}

type StateConfig struct {
	Address string `json:"address"`
}

type RunnerConfig struct {
	Manager ManagerConfig `json:"manager"`
	State   StateConfig   `json:"state"`
}
