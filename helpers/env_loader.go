package helpers

import (
	"os"
	"strings"
)

type CrispyPodEnv struct {
	PBAdminEnabled     bool
	PBAdminAcceptHosts []string
}

var LoadedEnv *CrispyPodEnv

func init() {
	LoadedEnv = &CrispyPodEnv{
		PBAdminEnabled:     os.Getenv("PB_ADMIN_ENABLED") == "1",
		PBAdminAcceptHosts: strings.Split(os.Getenv("PB_ADMIN_ACCEPT_HOSTS"), ","),
	}
}
