package ffmpeg

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os/exec"
)

// GetSnapShotByPath 获取视频截图
func GetSnapShotByPath(path string) (*bytes.Buffer, error) {
	buff := bytes.NewBuffer(nil)
	err := ffmpeg.Input(path).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buff).Run()
	return buff, err
}

func GetSnapShot(video []byte) ([]byte, error) {
	inputBuffer := bytes.NewBuffer(video)
	outputBuffer := bytes.NewBuffer(nil)

	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vf", `select=gte(n\,0)`, "-vframes", "1", "-f", "image2", "pipe:1")
	cmd.Stdin, cmd.Stdout = inputBuffer, outputBuffer

	err := cmd.Run()
	return outputBuffer.Bytes(), err
}
