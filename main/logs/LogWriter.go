package logs

type logWriter struct{}

func (p logWriter) Write(bs []byte) {}
