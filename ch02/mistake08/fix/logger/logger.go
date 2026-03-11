package logger

import "io"

//Package logger demonstrates a case where embedding is appropriate
//Write and Close are intentionally public, embedding removes forwarding boilerplate

//Logger wraps any io.WriteCloser
//Embedding is correct here: Write and Close are meant to be visible
//The promoted methods also make logger satisfy io.WriteCloser automatically
type Logger struct {
	io.WriteCloser
}

func New(wc io.WriteCloser) *Logger {
	return &Logger{WriteCloser: wc}
}

//No need to manually write forwarding methods like:
//Write() and Close() - Embedding handles this automatically
