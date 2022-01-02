package ffmpeg

import (
	"testing"
)

func Test_FFmpegTranscode(t *testing.T)  {
	var request = FFmpegTranscodeRequest{}
	request.Inputs = []string{"abc	     def g.mp4"}
	request.Outputs = []FFmpegTranscodeOutputParams{
		{
			Output: "output.mp4",
			Format: "mp4",
			Streams: []FFmpegTranscodeStreamParams{
				{
					Kind: "video",
					Video: FFmpegTranscodeVideoStreamParams{
						Codec: "h264",
						Preset: "slow",
						Height: 1080,
						Fps: 25,
					},
				},
				{
					Kind: "audio",
					Audio: FFmpegTranscodeAudioStreamParams{
						Channels: 2,
						Codec: "aac",
						SampleRate: 44100,
					},
				},
			},
		},
	}

	cmdString, err := FFmpegTranscode(request)
	if err != nil {
		t.Error("err:", err.Error())
	}
	t.Log("cmdString:", cmdString)
}