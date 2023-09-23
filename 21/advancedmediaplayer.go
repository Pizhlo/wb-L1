package main

type AdvancedMediaPlayer interface {
	PlayVideo(fileName string) error
	PlayAudio(fileName string) error
}
