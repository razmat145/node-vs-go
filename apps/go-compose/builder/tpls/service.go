package tpls

type Limits struct {
	Memory string `yaml:"memory,omitempty"`
	CPU    string `yaml:"cpus,omitempty"`
}

type Resources struct {
	Limits Limits `yaml:"limits,omitempty"`
}

type Build struct {
	DockerFile string `yaml:"dockerfile,omitempty"`
	Context    string `yaml:"context,omitempty"`
}

type Deploy struct {
	Replicas  int       `yaml:"replicas,omitempty"`
	Resources Resources `yaml:"resources,omitempty"`
}

type DockerComposeService struct {
	Build       Build             `yaml:"build,omitempty"`
	Image       string            `yaml:"image,omitempty"`
	Restart     string            `yaml:"restart,omitempty"`
	Ports       []string          `yaml:"ports,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Volumes     []string          `yaml:"volumes,omitempty"`
	Command     []string          `yaml:"command,omitempty"`
	Deploy      Deploy            `yaml:"deploy,omitempty"`
	DependsOn   []string          `yaml:"depends_on,omitempty"`
}

type ServiceConfig struct {
	dockerFile string
	context    string
	replicas   int
	resources  struct {
		limits struct {
			memory string
			cpu    string
		}
	}
	image       string
	restart     string
	ports       []string
	environment map[string]string
	volumes     []string
	command     []string
	dependsOn   []string
}

func NewDockerComposeService(serviceConfig ServiceConfig) DockerComposeService {
	return DockerComposeService{
		Build: Build{
			DockerFile: serviceConfig.dockerFile,
			Context:    serviceConfig.context,
		},
		Image:       serviceConfig.image,
		Restart:     serviceConfig.restart,
		Ports:       serviceConfig.ports,
		Environment: serviceConfig.environment,
		Volumes:     serviceConfig.volumes,
		Command:     serviceConfig.command,
		Deploy: Deploy{
			Replicas: serviceConfig.replicas,
			Resources: Resources{
				Limits: Limits{
					Memory: serviceConfig.resources.limits.memory,
					CPU:    serviceConfig.resources.limits.cpu,
				},
			},
		},
		DependsOn: serviceConfig.dependsOn,
	}
}
