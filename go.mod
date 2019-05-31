module github.com/nanohard/scribble

go 1.12

require (
	github.com/jcelliott/lumber v0.0.0-20160324203708-dd349441af25
	github.com/kelindar/binary v1.0.4
	github.com/nanohard/scribble/codec v0.0.0
	github.com/nanohard/scribble/codec/json v0.0.0
	github.com/nanohard/scribble/codec/binary v0.0.0
)

//replace github.com/nanohard/scribble/codec => ./codec

//replace github.com/nanohard/scribble/codec/json => ./codec/json

//replace github.com/nanohard/scribble/codec/binary => ./codec/binary
