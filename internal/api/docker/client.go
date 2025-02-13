package docker

import (
	"context"
	"fmt"

	"github.com/bluewave-labs/checkmate-cli/internal/fs"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	Client *client.Client
}

// BackupVolume creates a backup of the Docker volume identified by volumeName.
// It does so by launching a temporary container using the "alpine" image that mounts
// the specified volume at "/data". Within the container, it runs a tar command to compress
// the contents of the volume into a gzip archive.
// The function collects the container's output (which is the tarball data) from the container logs,
// writes this output to a file specified by backupPath, and then cleans up by removing the temporary container.
// Any errors encountered during container creation, starting, log retrieval, file operations, or data copying
// are returned accordingly.
func (d *DockerClient) BackupVolume(ctx context.Context, volumeName string, backupPath string) error {
	vol, err := d.Client.VolumeInspect(ctx, volumeName)
	if err != nil {
		return fmt.Errorf("error inspecting volume: %v", err)
	}

	return fs.CreateTar(vol.Mountpoint, backupPath)
}

// RestoreVolume restores the contents of a Docker volume from a backup file.
// It creates a temporary container with a mounted volume and uses the tar utility
// to extract the backup data into the volume.
// Parameters:
//   - ctx: The context for managing the lifecycle of the restore operation.
//   - volumeName: The name of the Docker volume to restore.
//   - backupPath: The file system path to the backup file (a tar.gz archive) containing the volume data.
//
// Returns:
//
//	An error if the restore operation fails at any step, otherwise nil.
func (d *DockerClient) RestoreVolume(ctx context.Context, volumeName string, backupPath string) error {
	vol, err := d.Client.VolumeInspect(ctx, volumeName)
	if err != nil {
		return fmt.Errorf("error inspecting volume: %v", err)
	}

	return fs.ExtractTarToVolume(backupPath, vol.Mountpoint)
}

func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %v", err)
	}

	return &DockerClient{Client: cli}, nil
}
