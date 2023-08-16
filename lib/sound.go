package lib

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

type SoundError struct {
	message  string
	filename string
	error    error
}

func (err SoundError) Error() string {
	var tail = ""
	if err.error != nil {
		tail = fmt.Sprintf(" : %s", err.error)
	}
	return fmt.Sprintf("%s: %s%s", err.message, err.filename, tail)
}

func PlaySafe(filename string) {
	if filename != "" {
		Log.ErrWrap(Play(filename))
	}
}

func Play(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return SoundError{"Couldn't open file", filename, err}
	}
	defer file.Close()

	var streamer beep.StreamSeekCloser
	var format beep.Format

	extension := path.Ext(filename)
	switch extension {
	case ".mp3":
		streamer, format, err = mp3.Decode(file)
	case ".ogg":
		fallthrough
	case ".oga":
		streamer, format, err = vorbis.Decode(file)
	case ".wav":
		streamer, format, err = wav.Decode(file)
	default:
		return SoundError{
			"No suitable player found",
			filename, nil,
		}
	}
	if err != nil {
		return SoundError{
			fmt.Sprintf("Failed decoding %s file", extension),
			filename, err,
		}
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate,
		format.SampleRate.N(time.Second/10))
	if err != nil {
		return SoundError{
			"Failed initializing the speaker",
			filename, err,
		}
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	// Wait for the audio to finish playing
	<-done

	return nil
}
