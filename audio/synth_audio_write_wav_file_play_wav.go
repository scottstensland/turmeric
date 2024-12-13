package main

//  go-wave/example/writing.go
//  https://github.com/zenwerk/go-wave/blob/20cf3f50c3377abfe9771ffac92190dafd34af38/example/writing.go#L17

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sync"
	"syscall/js"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/zenwerk/go-wave"
)

func synthesize_output_audio_file(output_wav_file string) {

	// Create the file for writing
	output_filehandle, err := os.Create(output_wav_file)
	if err != nil {
		panic(err)
	}
	defer output_filehandle.Close()

	// Set up the writer parameters
	param := wave.WriterParam{
		Out:            output_filehandle,
		WaveFormatType: 1, // PCM
		Channel:        1,
		SampleRate:     44100,
		BitsPerSample:  16,
	}

	// Create wave writer
	w, err := wave.NewWriter(param)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	amplitude := 0.1
	hz := 440.0
	length := param.SampleRate

	for i := 0; i < length; i++ {
		_data := amplitude * math.Sin(2.0*math.Pi*hz*float64(i)/float64(param.SampleRate))
		// Scale _data to the range of int16
		_data = (_data + 1.0) / 2.0 * 65536.0
		if _data > 65535.0 {
			_data = 65535.0
		} else if _data < 0.0 {
			_data = 0.0
		}
		// Convert to int16, as WriteSample16 expects []int16
		data := int16(_data + 0.5 - 32768) // Rounding and offset adjustment

		// Write the sample, note the type conversion from int16 to []int16
		_, err = w.WriteSample16([]int16{data})
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf(" open %s\n", output_wav_file)
}

func synth_audio_write_wav_file_play_wav() {

	output_wav_file := filepath.Join(os.TempDir(), "synthesize_audio_wav_file_then_play_audio_file.wav")

	synthesize_output_audio_file(output_wav_file)

	//  above synthesizes audio file   whereas   below renders this new wav file

	input_wav_file := output_wav_file

	// Open the WAV file
	input_audio_filehandle, err := os.Open(input_wav_file)
	if err != nil {
		log.Fatal(err)
	}
	defer input_audio_filehandle.Close()

	// Decode WAV file
	d, err := wav.DecodeWithSampleRate(44100, input_audio_filehandle)
	if err != nil {
		log.Fatal(err)
	}

	// Create an audio context
	ctx := audio.NewContext(44100)

	// Create a player
	p, err := ctx.NewPlayer(d)
	if err != nil {
		log.Fatal(err)
	}

	// Use a WaitGroup to wait for the audio to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Play audio in a goroutine
	go func() {
		defer wg.Done()
		p.Play()
		// Wait for audio to finish
		for p.IsPlaying() {
			// Sleep or do something else, but this simple loop checks if audio is still playing
		}
		fmt.Println("Audio finished playing")
	}()

	// Wait for the goroutine to signal that audio playback has finished
	wg.Wait()
}

func main() {

	js.Global().Set("synth_audio_write_wav_file_play_wav", js.FuncOf(synth_audio_write_wav_file_play_wav))

	select {} // block forever
}
