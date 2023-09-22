package main

type MediaPlayer interface {
	Play(audioType string, fileName string) error
}
