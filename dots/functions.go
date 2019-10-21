package dots

import (
	"fmt"
	"os"

	"github.com/keloran/go-dotfiles/console"
)

func (d Dots) functions() error {
	if !d.Files.UserFileExists(".functions") || d.Force {
		if d.Force {
			console.Warning("Forced .functions")
		}

		if d.Github {
			return d.Files.GetGithubFile(".functions")
		}

		console.Nice("Creating .functions")
		return d.createFunctions()
	}

	console.Info("Skipped .functions")
	return nil
}

func (d Dots) createFunctions() error {
	f, err := os.Create(fmt.Sprintf("%s/.functions", d.Prefix))
	if err != nil {
		return fmt.Errorf("function create err: %w", err)
	}

	err = mkd(f)
	if err != nil {
		return fmt.Errorf("mkd err: %w", err)
	}
	err = dockerPsClean(f)
	if err != nil {
		return fmt.Errorf("dockerPsClean err: %w", err)
	}
	err = dockerUI(f)
	if err != nil {
		return fmt.Errorf("dockerUI err: %w", err)
	}
	err = dockerStart(f)
	if err != nil {
		return fmt.Errorf("dockerStart err: %w", err)
	}
	err = dockerStop(f)
	if err != nil {
		return fmt.Errorf("dockerStop err: %w", err)
	}
	err = dockerExec(f)
	if err != nil {
		return fmt.Errorf("dockerExec err: %w", err)
	}
	err = dockerLogs(f)
	if err != nil {
		return fmt.Errorf("dockerLogs err: %w", err)
	}
	err = dockerRestart(f)
	if err != nil {
		return fmt.Errorf("dockerRestart err: %w", err)
	}
	err = cdiff(f)
	if err != nil {
		return fmt.Errorf("cdiff err: %w", err)
	}
	err = getdockerpath(f)
	if err != nil {
		return fmt.Errorf("getDockerPath err: %w", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("functions file close: %w", err)
	}

	err = d.Files.SetUserPerm(".functions")
	if err != nil {
		return fmt.Errorf("functions: %w", err)
	}

	return nil
}

func mkd(f *os.File) error {
	return writeFunc("mkd", "mkdir -p \"$@\" && cd \"$@\"\n", f)
}

func dockerPsClean(f *os.File) error {
	return writeFunc("dockerPsClean", "docker -ps -a --format '{{.Names}} {{.Status}}' | grep 'Exited' | awk '{print $1}' | xargs docker rm\n", f)
}

func dockerUI(f *os.File) error {
	return writeFunc("dockerUI", "docker run -dp 9001:9000 --privileged -v /var/run/docker.sock:/var/run/docker.sock uifd/ui-for-docker\n", f)
}

func dockerStart(f *os.File) error {
	return writeFunc("dockerStart", "dockerStop\n"+
		"  if [[ -e $(pwd)/docker.sh ]]; then\n"+
		"    sh $(pwd)/docker.sh\n"+
		"  else\n"+
		"    if [[ -e $(pwd)/docker-compose.yml ]]; then\n"+
		"      docker-compose build\n"+
		"      docker-compose up -d\n"+
		"      docker-compose -ps\n"+
		"    else\n"+
		"      docker build -t $(getDockerPath) .\n"+
		"      docker run -P --rm -d -it --name $(getDockerPath)_build $(getDockerPath)\n"+
		"      docker ps -a\n"+
		"    fi\n"+
		"  fi", f)
}

func dockerStop(f *os.File) error {
	return writeFunc("dockerStop", "if [[ -e $(pwd)/docker-compose.yml ]]; then\n"+
		"    docker-compose stop\n"+
		"    dockerPsClean\n"+
		"  else\n"+
		"    docker stop $(getDockerPath)_build\n"+
		"    dockerPsClean\n"+
		"  fi", f)
}

func dockerExec(f *os.File) error {
	return writeFunc("dockerExec", "if [[ -e $(pwd)/docker-compose.yml ]]; then\n"+
		"    if [[ -z $1 ]]; then\n"+
		"      echo \"you need to provide a server\"\n"+
		"      docker-compose ps\n"+
		"    else\n"+
		"      if [[ -z \"$2\" ]]; then\n"+
		"        docker-compose exec $1 sh\n"+
		"      else\n"+
		"        docker-compose exec $1 $2\n"+
		"      fi\n"+
		"    fi\n"+
		"  else\n"+
		"    if [[ -z \"$1\" ]]; then\n"+
		"      docker exec $(getDockerPath)_build sh\n"+
		"    else\n"+
		"      docker exec $(getDockerPath)_build $1\n"+
		"    fi\n"+
		"  fi\n", f)
}

func dockerLogs(f *os.File) error {
	return writeFunc("dockerLogs", "if [[ -e $(pwd)/docker-compose.yml ]]; then\n"+
		"    if [[ -z $1 ]]; then\n"+
		"      docker-compose logs -f\n"+
		"    else\n"+
		"      docker-compose logs -f $1\n"+
		"    fi\n"+
		"  else\n"+
		"    docker logs -f $(getDockerPath)_build\n"+
		"  fi", f)
}

func dockerRestart(f *os.File) error {
	return writeFunc("dockerRestart", "if [[ -e $(pwd)/docker-compose.yml ]]; then\n"+
		"    if [[ -z $1 ]]; then\n"+
		"      dockerStop\n"+
		"      dockerStart\n"+
		"    else\n"+
		"      docker-compose stop $1\n"+
		"      docker-compose up -d $1\n"+
		"    fi\n"+
		"  else\n"+
		"    dockerStop\n"+
		"    dockerStart\n"+
		"  fi", f)
}

func cdiff(f *os.File) error {
	return writeFunc("cdiff", "colordiff -u \"$@\"", f)
}

func writeFunc(funcName string, cont string, f *os.File) error {
	_, err := f.WriteString(fmt.Sprintf("\n\nfunction %s()\n{\n  %s\n}", funcName, cont))
	if err != nil {
		return fmt.Errorf("%s: %w", funcName, err)
	}

	return nil
}

func getdockerpath(f *os.File) error {
	return writeFunc("getDockerPath", "echo $(basename ${PWD##*/} | tr 'A-Z' 'a-z')\n", f)
}
