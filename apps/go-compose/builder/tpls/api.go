package tpls

func GetDockerComposeAPIServices(appsToRun []string) map[string]DockerComposeService {
	var services = make(map[string]DockerComposeService)

	apiDefaultDeployCfg := Deploy{
		Replicas: 1,
		Resources: Resources{
			Limits: Limits{
				Memory: "1200M",
				CPU:    "2",
			},
		},
	}

	for _, app := range appsToRun {
		switch app {
		case "node-express":
			services["node-express"] = DockerComposeService{
				Build: Build{
					DockerFile: "./apps/node-express/Dockerfile",
					Context:    ".",
				},
				Image:   "node-express:latest",
				Restart: "always",
				Ports:   []string{"8989:8989"},
				Environment: map[string]string{
					"API_PORT": "8989",
				},
				Deploy: apiDefaultDeployCfg,
			}

		case "node-std":
			services["node-std"] = DockerComposeService{
				Build: Build{
					DockerFile: "./apps/node-std/Dockerfile",
					Context:    ".",
				},
				Image:   "node-std:latest",
				Restart: "always",
				Ports:   []string{"8991:8991"},
				Environment: map[string]string{
					"API_PORT": "8991",
				},
				Deploy: apiDefaultDeployCfg,
			}

		case "go-echo":
			services["go-echo"] = DockerComposeService{
				Build: Build{
					DockerFile: "./apps/go-echo/Dockerfile",
					Context:    ".",
				},
				Image:   "go-echo:latest",
				Restart: "always",
				Ports:   []string{"8990:8990"},
				Environment: map[string]string{
					"API_PORT": "8990",
				},
				Deploy: apiDefaultDeployCfg,
			}

		case "go-std":
			services["go-std"] = DockerComposeService{
				Build: Build{
					DockerFile: "./apps/go-std/Dockerfile",
					Context:    ".",
				},
				Image:   "go-std:latest",
				Restart: "always",
				Ports:   []string{"8992:8992"},
				Environment: map[string]string{
					"API_PORT": "8992",
				},
				Deploy: apiDefaultDeployCfg,
			}
		}
	}

	return services
}
