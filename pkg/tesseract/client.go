package tesseract

import (
	"context"
	"io"
	"os/exec"
)

type Client struct {
	binaryPath string
}

func NewClient() *Client {
	return &Client{binaryPath: "tesseract"}
}

func (c *Client) Version() ([]byte, error) {
	cmd := exec.Command(c.binaryPath, "--version")

	return cmd.Output()
}

func (c *Client) Help() ([]byte, error) {
	cmd := exec.Command(c.binaryPath, "--help-extra")

	return cmd.Output()
}

func (c *Client) Languages() ([]byte, error) {
	cmd := exec.Command(c.binaryPath, "--list-langs")

	return cmd.Output()
}

func (c *Client) Parameters() ([]byte, error) {
	cmd := exec.Command(c.binaryPath, "--print-parameters")

	return cmd.Output()
}

func (c *Client) Convert(ctx context.Context, image io.Reader) ([]byte, error) {
	options := parseContext(ctx)

	arguments := append([]string{"stdin", "stdout"}, options...)

	cmd := exec.Command(c.binaryPath, arguments...)
	cmd.Stdin = image

	return cmd.Output()
}
