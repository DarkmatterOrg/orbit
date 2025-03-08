package utils

import (
	"os/exec"
)

func Disable_systemd_service(service_name string) {
	services := []struct {
		path        string
		serviceName string
	}{
		{"/usr/share/horizon", service_name},
		{"/usr/share/nova", service_name},
		{"/usr/share/umbra", service_name},
	}

	for _, service := range services {
		if PathExists(service.path) {
			cmd := exec.Command("systemctl", "disable", service.serviceName)
			err := cmd.Run()

			if err != nil {
				Error("Failed to disable the systemd serivce.")
			}

			return
		}
	}
}

func Disable_user_systemd_service(service_name string) {
	services := []struct {
		path        string
		serviceName string
	}{
		{"/usr/share/horizon", service_name},
		{"/usr/share/nova", service_name},
		{"/usr/share/umbra", service_name},
	}

	for _, service := range services {
		if PathExists(service.path) {
			cmd := exec.Command("systemctl", "disable", "--global", service.serviceName)
			err := cmd.Run()

			if err != nil {
				Error("Failed to disable the user systemd serivce.")
			}

			return
		}
	}
}

func Enable_systemd_service(service_name string) {
	services := []struct {
		path        string
		serviceName string
	}{
		{"/usr/share/horizon", service_name},
		{"/usr/share/nova", service_name},
		{"/usr/share/umbra", service_name},
	}

	for _, service := range services {
		if PathExists(service.path) {
			cmd := exec.Command("systemctl", "enable", service.serviceName)
			err := cmd.Run()

			if err != nil {
				Error("Failed to enable the systemd serivce.")
			}

			return
		}
	}
}

func Enable_user_systemd_service(service_name string) {
	services := []struct {
		path        string
		serviceName string
	}{
		{"/usr/share/horizon", service_name},
		{"/usr/share/nova", service_name},
		{"/usr/share/umbra", service_name},
	}

	for _, service := range services {
		if PathExists(service.path) {
			cmd := exec.Command("systemctl", "enable", "--global", service.serviceName)
			err := cmd.Run()

			if err != nil {
				Error("Failed to enable the user systemd serivce.")
			}

			return
		}
	}
}
