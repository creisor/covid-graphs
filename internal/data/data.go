package data

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const Repo = "https://github.com/CSSEGISandData/COVID-19.git"

type Data struct {
	Directory string
}

func (d *Data) Clone() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "git", "clone", "--depth", "1", Repo, d.Directory)

	log.Info(fmt.Sprintf("Cloning %s", Repo))
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "Failed to clone %s", Repo)
	}
	return nil
}

func (d *Data) Pull() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "git", "pull", Repo)

	log.Info(fmt.Sprintf("Pulling %s", Repo))
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "Failed to pull %s", Repo)
	}
	return nil
}
