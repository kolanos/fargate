package docker

import (
	"os"
	"os/exec"
	"time"

	"github.com/jpignata/fargate/console"
)

const timestampFormat = "20060102150405"

func GenerateTag() string {
	return time.Now().UTC().Format(timestampFormat)
}

type Repository struct {
	Uri string
}

func (repository *Repository) Login(username, password string) {
	console.Debug("Logging into Docker repository [%s]", repository.Uri)
	console.Shell("docker login --username %s --password ******* %s", username, repository.Uri)

	cmd := exec.Command("docker", "login", "--username", username, "--password", password, repository.Uri)

	if console.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		console.ErrorExit(err, "Couldn't login to Docker repository [%s]", repository.Uri)
	}

	cmd.Wait()
}

func (repository *Repository) Build(tag string) {
	console.Debug("Building Docker image [%s]", repository.UriFor(tag))
	console.Shell("docker build --tag %s .", repository.UriFor(tag))

	cmd := exec.Command("docker", "build", "--tag", repository.Uri+":"+tag, ".")

	if console.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		console.ErrorExit(err, "Couldn't build Docker image [%s]", repository.UriFor(tag))
	}

	cmd.Wait()
}

func (repository *Repository) Push(tag string) {
	console.Debug("Pushing Docker image [%s]", repository.UriFor(tag))
	console.Shell("docker push %s .", repository.UriFor(tag))

	cmd := exec.Command("docker", "push", repository.UriFor(tag))

	if console.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		console.ErrorExit(err, "Couldn't push Docker image [%s]", repository.UriFor(tag))
	}

	cmd.Wait()
}

func (repository *Repository) UriFor(tag string) string {
	return repository.Uri + ":" + tag
}