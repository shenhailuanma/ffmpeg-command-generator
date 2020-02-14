package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenhailuanma/ffmpeg-command-generator/apiserver/models"
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateSnapshotController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	var request = ffmpeg.FFmpegTranscodeRequest{}
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("CreateSnapshotController, bind data, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	cmd, err := ffmpeg.FFmpegSnapshot(request)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Msg = err.Error()
		logrus.Error("CreateSnapshotController, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	response.Data = cmd

	c.JSON(response.Status, &response)
}