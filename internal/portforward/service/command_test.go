package service

import (
	"context"
	"testing"

	"github.com/qdm12/gluetun/internal/command"
	"github.com/qdm12/log"
	"github.com/stretchr/testify/require"
)

func Test_Service_runUpCommand(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cmder := command.New()
	const commandTemplate = `/bin/sh -c "echo {{PORTS}}"`
	ports := []uint16{1234, 5678}
	logger := log.New()

	err := runUpCommand(ctx, cmder, logger, commandTemplate, ports)

	require.NoError(t, err)
}
