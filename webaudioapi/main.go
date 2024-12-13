package main

import (
	"syscall/js"
)

func main() {
	// Access the audio context from JavaScript
	audioCtx := js.Global().Get("AudioContext").New()

	// Create an oscillator node
	oscillator := audioCtx.Call("createOscillator")
	oscillator.Set("type", "sine")
	oscillator.Set("frequency", 440) // A4 note

	// Connect the oscillator to the audio context's destination
	oscillator.Call("connect", audioCtx.Get("destination"))

	// Start playing the sound
	oscillator.Call("start")
}
