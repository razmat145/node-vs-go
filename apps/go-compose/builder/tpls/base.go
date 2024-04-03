package tpls

type Network struct {
	Name string `yaml:"name"`
}

type DockerComposeBase struct {
	Version  string                          `yaml:"version"`
	Services map[string]DockerComposeService `yaml:"services"`
	Networks map[string]Network              `yaml:"networks"`
}

func NewDockerComposeBase(services map[string]DockerComposeService) DockerComposeBase {
	return DockerComposeBase{
		Version:  "3.0",
		Services: services,

		Networks: map[string]Network{
			"default": {
				Name: "my-playground-network",
			},
		},
	}
}
