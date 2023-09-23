package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// VLCPlayer реализует AdvancedMediaPlayer
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVideo(fileName string) error {
	fmt.Printf("Playing video file. Name: %s\n", fileName)
	return nil
}

func (v *VLCPlayer) PlayAudio(fileName string) error {
	fmt.Printf("Playing audio file. Name: %s\n", fileName)
	return nil
}

// VLCPlayer реализует MediaPlayer
type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) Play(audioType string, fileName string) error {
	if audioType == "vlc" {
		return m.advancedMusicPlayer.PlayVideo(fileName)
	} else if audioType == "mp4" {
		return m.advancedMusicPlayer.PlayAudio(fileName)
	}
	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

type AudioPlayer struct {
	mediaAdapter MediaPlayer
}

func (a *AudioPlayer) Play(audioType string, fileName string) error {
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file. Name: %s\n", fileName)
		return nil
	} else if audioType == "vlc" || audioType == "mp4" {
		a.mediaAdapter = &MediaAdapter{&VLCPlayer{}}
		return a.mediaAdapter.Play(audioType, fileName)
	}

	return fmt.Errorf("invalid media. %s format not supported", audioType)
}

func main() {
	player := AudioPlayer{}

	player.Play("vlc", "test")
}
