package osspecific

import (
	"fmt"
	"os"
	"os/exec"
)

func (o OSSpecific) macInstall() error {
	err := o.SSH()
	if err != nil {
		return fmt.Errorf("OSX SSH err: %w", err)
	}

	err = o.MDS()
	if err != nil {
		return fmt.Errorf("OSX MDS err: %w", err)
	}

	err = o.DNS()
	if err != nil {
		return fmt.Errorf("OSX DNS err: %w", err)
	}

	return nil
}

// SSH ...
func (o OSSpecific) SSH() error {
	cmd := exec.Command("systemsetup", "-setremotelogin", "on")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// MDS ...
func (o OSSpecific) MDS() error {
	cmd := exec.Command("killall", "mds", ">", "/dev/null", "2>&1")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("MDS killall: %w", err)
	}

	cmd = exec.Command("sudo", "mdutil", "-i", "on", ">", "/dev/null")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("mdutil -i: %w", err)
	}

	cmd = exec.Command("sudo", "mdutil", "-E", "/", ">", "/dev/null")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("mdutil -E: %w", err)
	}

	return nil
}

// DNS ...
func (o OSSpecific) DNS() error {
	err := o.createFile()
	if err != nil {
		return fmt.Errorf("DNS createfile: %w", err)
	}

	cmd := exec.Command("sudo", "cp", fmt.Sprintf("%s/com.docker_alias.plist", o.Prefix), "/Library/LaunchDaemons/com.docker_alias.plist")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("DNS copy alias: %w", err)
	}

	cmd = exec.Command("sudo", "launchctl", "load", "/Library/LaunchDaemons/com.docker_alias.plist")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("DNS launchctl: %w", err)
	}

	cmd = exec.Command("echo")

	return nil
}

func (o OSSpecific) createFile() error {
	f, err := os.Create(fmt.Sprintf("%s/com.docker_alias.plist", o.Prefix))
	if err != nil {
		return fmt.Errorf("create file create: %w", err)
	}

	_, err = f.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
		"<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\"\n" +
		"<plist version=\"1.0\">\n" +
		"  <dict>\n" +
		"    <key>Label</key>\n" +
		"    <string>com.docker_alias</string>\n" +
		"    <key>ProgramArguments>\n" +
		"    <array>\n" +
		"      <string>ifconfig</string>\n" +
		"      <string>lo0</string>\n" +
		"      <string>alias</string>\n" +
		"      <string>10.254.254.254</string>\n" +
		"    </array>\n" +
		"    <key>RunAtLoad</key>\n" +
		"    <true />\n" +
		"  </dict>\n" +
		"</plist>\n")
	if err != nil {
		return fmt.Errorf("create file writestring: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("create file close: %w", err)
	}

	return nil
}
