package main

import (
    "image/gif"
    "os"
    "log"
)

func handleError (err error, message string) {
    if err != nil {
        log.Fatal(message)
    }
}

func main () {
    if len(os.Args) < 2 {
        log.Fatal("Infile/Outfile path not provided")
    }

    // Reading and parsing the input GIF
    infile, err := os.Open(os.Args[1])
    handleError(err, "Unable to open `"+os.Args[1]+"`.")

    inGif, err := gif.DecodeAll(infile)
    handleError(err, "Unable to parse GIF.")

    // Adding all of the image frames in reverse order
    outGif := &gif.GIF{}
    for i := len(inGif.Image)-1; i >= 0; i-- {
        outGif.Image = append(outGif.Image, inGif.Image[i])
        outGif.Delay = append(outGif.Delay, inGif.Delay[i])
    }
    outGif.LoopCount = inGif.LoopCount

    // Writing the reversed GIF to the given out file
    f, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0600)
    handleError(err, "Could not open outfile.")
    defer f.Close()
    gif.EncodeAll(f, outGif)
}
