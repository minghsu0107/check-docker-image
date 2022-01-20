package main

import (
	"errors"
	"strings"

	"github.com/minghsu0107/check-docker-image/registry"
	log "github.com/sirupsen/logrus"
)

func parseImage(image string) (string, string, error) {
	splitTag := strings.Split(image, ":")
	if len(splitTag) != 2 {
		return "", "", errors.New("invalid image format")
	}
	imageName := splitTag[0]
	imageTag := splitTag[1]

	return imageName, imageTag, nil
}

func checkImages(hub *registry.Registry, images []string) bool {
	existTagCnt := 0
	for _, image := range images {
		imageName, imageTag, err := parseImage(image)
		if err != nil {
			log.Errorf("image name parsing error: %v", err)
			return false
		}
		exist, err := hub.Tags(imageName, imageTag)
		if err != nil {
			log.Errorf("listing image tags error: %v", err)
			return false
		}
		if exist {
			log.Infof("image %s found", image)
			existTagCnt++
		} else {
			log.Errorf("image %s not found", image)
		}
	}
	return existTagCnt == len(images)
}
