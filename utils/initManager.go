package utils

import (
	"fmt"
	"os"
	"os/exec"
)

const FAILED_DISABLE_SERVICE = "Failed to disable "
const FAILED_ENABLE_SERVICE = "Failed to enable "
const DISABLE_SERVICE = "Disabled "
const ENABLED_SERVICE = "Enabled "
const SET_TARGET = "Set target to "
const FAILED_SET_TARGET = "Failed to set target to "

var service_name string
var msg = ENABLED_SERVICE + service_name
var failed_msg = FAILED_ENABLE_SERVICE + service_name
var msg_disabled = DISABLE_SERVICE + service_name
var failed_msg_disabled = FAILED_DISABLE_SERVICE + service_name
var openrc_path = fmt.Sprintf("/etc/init.d/" + user + "." + user)
var user = os.Getenv("USER")

// Enable a runit service by making a symlink from /etc/sv to /var/service
/*
	Example
	EnableRunitService("NetworkManager")
*/
func EnableRunitService(service_name string) {
	var service_path = "/var/service"

	var full_service = "/etc/sv/" + service_name

	err := os.Symlink(full_service, service_path)

	if err != nil {
		Error(failed_msg)
	} else {
		Done(msg)
	}
}

// Disable a runit service by removing a symlink from /var/service
/*
	Example
	DisableRunitService("NetworkManager")
*/
func DisableRunitService(service_name string) {
	var service_path = "/var/service"
	var full_service = service_path + service_name

	err := os.RemoveAll(full_service)

	if err != nil {
		Error(failed_msg_disabled)
	} else {
		Done(msg_disabled)
	}
}

// Disable a systemd service by calling systemctl disable or systemctl --user disable
/*
	Example (system-service)
	DisableSystemdService("NetworkManager", false)

	Example (user-service)
	DisableSystemdService("pipewire", true)
*/
func DisableSystemdService(service_name string, user_mode bool) {

	if user_mode {
		cmd := exec.Command("systemctl", "--user", "disable", service_name)
		err := cmd.Run()

		if err != nil {
			Error(failed_msg_disabled)
		} else {
			Done(msg_disabled)
		}
	} else {
		cmd := exec.Command("systemctl", "disable", service_name)
		err := cmd.Run()

		if err != nil {
			Error(failed_msg_disabled)
		} else {
			Done(msg_disabled)
		}
	}
}

// Enable a openRC service by calling rc-update add or rc-update --user add
/*
	Example (system-service)
	EnableOpenrcService("NetworkManager", false)

	Example (user-service)
	EnableOpenrcService("pipewire", true)
*/
func EnableSystemdService(service_name string, user_mode bool) {
	if user_mode {
		cmd := exec.Command("systemctl", "--user", "enable", service_name)
		err := cmd.Run()

		if err != nil {
			Error(failed_msg)
		} else {
			Done(msg)
		}
	} else {
		cmd := exec.Command("systemctl", "enable", service_name)
		err := cmd.Run()

		if err != nil {
			Error(failed_msg)
		} else {
			Done(msg)
		}
	}
}

func SetSystemdTarget(target_name string) {
	cmd := exec.Command("systemctl", "isolate", target_name)
	err := cmd.Run()

	if err != nil {
		Error(FAILED_SET_TARGET)
	} else {
		Done(SET_TARGET)
	}
}

func EnableOpenrcService(service_name string, user_mode bool) {
	if _, err := os.Stat(openrc_path); os.IsNotExist(err) && user_mode {
		enable_command := exec.Command("rc-update", "--user", "add", service_name)

		err := enable_command.Run()

		if err != nil {
			Error(failed_msg)
		} else {
			Done(msg)
		}
	} else {
		enable_command := exec.Command("rc-update", "add", service_name)

		err := enable_command.Run()

		if err != nil {
			Error(failed_msg)
		} else {
			Done(msg)
		}
	}
}

// Disable a openRC service by calling rc-update del or rc-update --user del
/*
	Example (system-service)
	DisableOpenrcService("NetworkManager", false)

	Example (user-service)
	DisableOpenrcService("pipewire", true)
*/

func DisableOpenrcService(service_name string, user_mode bool) {
	if _, err := os.Stat(openrc_path); os.IsNotExist(err) && user_mode {
		disable_command := exec.Command("rc-update", "del", service_name)

		err := disable_command.Run()

		if err != nil {
			Error(failed_msg_disabled)
		} else {
			Done(msg_disabled)
		}
	} else {
		disable_command := exec.Command("rc-update", "del", service_name)

		err := disable_command.Run()

		if err != nil {
			Error(failed_msg_disabled)
		} else {
			Done(msg_disabled)
		}
	}
}

// Enable a Dinit service by calling dinitctl enable
/*
	Example (system-service)
	EnableDinitcService("NetworkManager")

	Example (user-service)
	EnableDinitService("pipewire")
*/
func EnableDinitService(service_name string) {
	enable_command := exec.Command("dinitctl", "enable", service_name)

	err := enable_command.Run()

	if err != nil {
		Error(failed_msg)
	} else {
		Done(msg)
	}
}

// Disable a Dinit service by calling dinitctl disable
/*
	Example (system-service)
	DisableDinitcService("NetworkManager")

	Example (user-service)
	DisableDinitService("pipewire")
*/

func DisableDinitService(service_name string) {
	disable_command := exec.Command("dinitctl", "disable", service_name)

	err := disable_command.Run()

	if err != nil {
		Error(failed_msg_disabled)
	} else {
		Done(msg_disabled)
	}
}
