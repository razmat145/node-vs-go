package tpls

func GetDockerComposeDefaultServices() map[string]DockerComposeService {
	return map[string]DockerComposeService{
		"prometheus": {
			Image:   "prom/prometheus:latest",
			Restart: "always",
			Volumes: []string{"./configs/prom/prometheus.yaml:/etc/prometheus/prometheus.yaml"},
			Ports:   []string{"9090:9090"},
			Command: []string{"--config.file=/etc/prometheus/prometheus.yaml"},
		},
		"grafana": {
			Image:   "grafana/grafana:latest",
			Restart: "always",
			Ports:   []string{"3000:3000"},
			Environment: map[string]string{
				"GF_SECURITY_ADMIN_USER":           "admin",
				"GF_SECURITY_ADMIN_PASSWORD":       "admin",
				"GF_SECURITY_ADMIN_PASSWORD_FORCE": "false",
			},
			Volumes:   []string{"./configs/prom/datasources.yaml:/etc/grafana/provisioning/datasources/datasources.yaml"},
			DependsOn: []string{"prometheus"},
		},
	}
}
