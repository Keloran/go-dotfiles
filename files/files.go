package files

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// Files ...
type Files struct {
	Username string
	Prefix   string
}

// GetFile ...
func (f Files) GetFile(path string, names ...string) error {
	c := http.Client{
		Timeout: 10 * time.Minute,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     2 * time.Minute,
		},
	}

	var filename, outname string

	if len(names) == 1 {
		filename = names[0]
		outname = filename
	} else {
		filename = names[0]
		outname = names[1]
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", path, filename), nil)
	if err != nil {
		return fmt.Errorf("%s get err: %w", filename, err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("%s do err: %w", filename, err)
	}

	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s content err: %w", filename, err)
	}

	err = resp.Body.Close()
	if err != nil {
		return fmt.Errorf("getfile close body: %w", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", f.Prefix, outname), cont, 0644)
	if err != nil {
		return fmt.Errorf("%s write err: %w", outname, err)
	}

	err = f.SetUserPerm(outname)
	if err != nil {
		return fmt.Errorf("%s perm: %w", outname, err)
	}

	return nil
}

// GetGithubFile ...
func (f Files) GetGithubFile(filename string) error {
	return f.GetFile("https://raw.githubusercontent.com/Keloran/dotfiles/master", filename)
}

// FileExists ...
func (f Files) FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// UserFileExists ...
func (f Files) UserFileExists(filename string) bool {
	return f.FileExists(fmt.Sprintf("%s/%s", f.Prefix, filename))
}

// MakeFolder ...
func (f Files) MakeFolder(folder string) error {
	err := os.Mkdir(folder, 755)
	if err != nil {
		return fmt.Errorf("makefolder: %w", err)
	}

	return nil
}

// MakeUserFolder ...
func (f Files) MakeUserFolder(folder string) error {
	err := f.MakeFolder(fmt.Sprintf("%s/%s", f.Prefix, folder))
	if err != nil {
		return fmt.Errorf("make userfolder: %w", err)
	}

	err = f.SetPerms(folder)
	if err != nil {
		return fmt.Errorf("perms userfolder: %w", err)
	}

	return nil
}

// SetPerms ...
func (f Files) SetPerms(folder string) error {
	folder = fmt.Sprintf("%s/%s", f.Prefix, folder)

	err := exec.Command("chown", "-R", f.Username, folder).Run()
	if err != nil {
		return fmt.Errorf("setperms >> chown: %w", err)
	}

	err = exec.Command("chgrp", "-R", "adusers", folder).Run()
	if err != nil {
		return fmt.Errorf("setperms >> chgrp: %w", err)
	}

	err = exec.Command("chmod", "755", "-R", folder).Run()
	if err != nil {
		return fmt.Errorf("setperms >> chmod: %w", err)
	}

	return nil
}

// SetUserPerm ...
func (f Files) SetUserPerm(file string) error {
	fName := fmt.Sprintf("%s/%s", f.Prefix, file)

	err := exec.Command("chown", "-R", f.Username, fName).Run()
	if err != nil {
		return fmt.Errorf("setuserperm >> chown: %w", err)
	}

	err = exec.Command("chgrp", "-R", "adusers", fName).Run()
	if err != nil {
		return fmt.Errorf("setuserperms >> chgrp: %w", err)
	}

	err = exec.Command("chmod", "644", fName).Run()
	if err != nil {
		return fmt.Errorf("setuserperms >> chmod: %w", err)
	}

	return nil
}

// MoveFile ...
func (f Files) MoveFile(file string, newloc string) error {
	err := os.Rename(file, newloc)
	if err != nil {
		return fmt.Errorf("movefile: %w", err)
	}

	return nil
}
