package ffmpeg

import (
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg/templates"
)

type FFmpegTranscodeVideoStreamParams struct {
	Map     string `json:"map"` // strams map, http://ffmpeg.org/ffmpeg-all.html#Advanced-options
	Codec   string `json:"codec"`
	Preset  string `json:"preset"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Bitrate string `json:"bitrate"`
	Fps     int    `json:"fps"`
	CRF     int    `json:"crf"`
}
type FFmpegTranscodeAudioStreamParams struct {
	Map        string `json:"map"` // strams map, http://ffmpeg.org/ffmpeg-all.html#Advanced-options
	Codec      string `json:"codec"`
	Channels   int    `json:"channles"`
	SampleRate int    `json:"sample_rate"`
	Bitrate    string `json:"bitrate"`
}

type FFmpegTranscodeStreamParams struct {
	Kind  string                           `json:"kind"` // video, audio, subtitle, data
	Video FFmpegTranscodeVideoStreamParams `json:"video"`
	Audio FFmpegTranscodeAudioStreamParams `json:"audio"`
}

type FFmpegTranscodeOutputParams struct {
	Output  string                        `json:"output"`
	Format  string                        `json:"format"`
	Streams []FFmpegTranscodeStreamParams `json:"streams"`
}

type FFmpegTranscodeRequest struct {
	Inputs  []string                      `json:"inputs"`
	Outputs []FFmpegTranscodeOutputParams `json:"outputs"`
	Globals FFmpegGlobalParams            `json:"globals"`
}

func FFmpegTranscode(request FFmpegTranscodeRequest) (string, error) {
	return templates.GenerateCommand("transcode", templates.TranscodeTemplate, request)
}
