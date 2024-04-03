package tpls

func GetDockerComposeSpammerServices(target string, appsToRun []string) map[string]DockerComposeService {
	var services = make(map[string]DockerComposeService)

	defaultSpammerBuild := Build{
		DockerFile: "./apps/go-spammer/Dockerfile",
		Context:    ".",
	}
	defaultSpammerDeployCfg := Deploy{
		Replicas: 1,
		Resources: Resources{
			Limits: Limits{
				Memory: "600M",
				CPU:    "2",
			},
		},
	}
	spamUri := getSpamUri(target)

	dependsOn := append(appsToRun, "prometheus")

	for _, app := range appsToRun {
		switch app {
		case "node-express":
			services["go-spamm-node-express"] = DockerComposeService{
				Build: defaultSpammerBuild,
				Image: "go-std:latest",
				Environment: map[string]string{
					"SPAM_URL": "http://node-express:8989/" + spamUri,
				},
				Deploy:    defaultSpammerDeployCfg,
				DependsOn: dependsOn,
			}

		case "node-std":
			services["go-spamm-node-std"] = DockerComposeService{
				Build: defaultSpammerBuild,
				Image: "go-std:latest",
				Environment: map[string]string{
					"SPAM_URL": "http://node-std:8991/" + spamUri,
				},
				Deploy:    defaultSpammerDeployCfg,
				DependsOn: dependsOn,
			}

		case "go-echo":
			services["go-spamm-go-echo"] = DockerComposeService{
				Build: defaultSpammerBuild,
				Image: "go-std:latest",
				Environment: map[string]string{
					"SPAM_URL": "http://go-echo:8990/" + spamUri,
				},
				Deploy:    defaultSpammerDeployCfg,
				DependsOn: dependsOn,
			}

		case "go-std":
			services["go-spamm-go-std"] = DockerComposeService{
				Build: defaultSpammerBuild,
				Image: "go-std:latest",
				Environment: map[string]string{
					"SPAM_URL": "http://go-std:8992/" + spamUri,
				},
				Deploy:    defaultSpammerDeployCfg,
				DependsOn: dependsOn,
			}
		}
	}

	return services
}

func getSpamUri(target string) string {
	switch target {
	case "hello-world":
		return "hello"
	case "factorial":
		return "factorial"
	case "make-garbage":
		return "garbage"
	}

	return ""
}
