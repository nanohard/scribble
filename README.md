# Scribble
[![GoDoc](https://godoc.org/github.com/boltdb/bolt?status.svg)](http://godoc.org/github.com/nanohard/scribble) [![Go Report Card](https://goreportcard.com/badge/github.com/nanohard/scribble)](https://goreportcard.com/report/github.com/nanohard/scribble)
--------

Flat-file database in Golang using choosable encoders.

Currently supports:
- JSON
- Binary

### Installation

Install using `go get github.com/nanohard/scribble`.

### Usage
Unless you have a need to use JSON, it is recommended to use the default Binary for reduced file 
size.

#### Initialize Database
```go
import(
    "github.com/nanohard/scribble"

    // Needed for Unmarshaling when using db.ReadAll()
    "github.com/nanohard/scribble/codec/json"
    // Or
    "github.com/nanohard/scribble/codec/binary"
)

// A new scribble driver, providing the directory where it will be writing to,
// and a qualified logger if desired.
// Defaults to using encoding/binary (well, really kelindar/binary)
db, err := scribble.New(dir, nil) // `binary.Codec` is used by default
if err != nil {
    fmt.Println("Error", err)
}

// Or

options := scribble.Options{
    Codec: json.Codec,  // `binary.Codec` is used by default
}
db, err := scribble.New(dir, options)
if err != nil {
    fmt.Println("Error", err)
}
```

#### DB Functions
```go
// Write a fish to the database
fish := Fish{}
if err := db.Write("fish", "onefish", fish); err != nil {
    fmt.Println("Error", err)
}

// Read a fish from the database (passing fish by reference)
onefish := Fish{}
if err := db.Read("fish", "onefish", &onefish); err != nil {
    fmt.Println("Error", err)
}

// Read all fish from the database, unmarshaling the response.
records, err := db.ReadAll("fish")
if err != nil {
    fmt.Println("Error", err)
}

fishies := []Fish{}
for _, f := range records {
    fishFound := Fish{}
    if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
        fmt.Println("Error", err)
    }
    fishies = append(fishies, fishFound)
}

// Delete a fish from the database
if err := db.Delete("fish", "onefish"); err != nil {
    fmt.Println("Error", err)
}

// Delete all fish from the database
if err := db.Delete("fish", ""); err != nil {
    fmt.Println("Error", err)
}
```

## Documentation
- Complete documentation is available on [godoc](http://godoc.org/github.com/nanohard/scribble).

## Todo/Doing
- More encoders?

## Contributing
- Rebase into one commit.
- Have liberal code comments.
