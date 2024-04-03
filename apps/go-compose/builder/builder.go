package builder

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/razmat145/learning/go-compose/builder/tpls"
)

type BuildProgress struct {
	Progress string
}

type RunEnvironment struct {
	BenchmarkTarget string

	AppsToRun []string
}

func (rEnv RunEnvironment) Build() {
	services := tpls.GetDockerComposeDefaultServices()

	services = attachToMap(services, tpls.GetDockerComposeAPIServices(rEnv.AppsToRun))
	services = attachToMap(services, tpls.GetDockerComposeSpammerServices(rEnv.BenchmarkTarget, rEnv.AppsToRun))

	baseDockerTpl := tpls.NewDockerComposeBase(services)

	data, err := yaml.Marshal(&baseDockerTpl)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = os.WriteFile(getBaseMonorepoPath()+"/docker-compose.yaml", data, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func attachToMap(target map[string]tpls.DockerComposeService, source map[string]tpls.DockerComposeService) map[string]tpls.DockerComposeService {
	for k, v := range source {
		target[k] = v
	}

	return target
}

func getBaseMonorepoPath() string {
	cwd, _ := os.Getwd()

	for filepath.Base(cwd) != "node-vs-go" {
		cwd = filepath.Dir(cwd)

	}

	return cwd
}
