package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func run() error {

	resp, err := http.Get("http://localhost:8080/test.mp3")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	d, err := mp3.NewDecoder(resp.Body)
	if err != nil {
		return err
	}
	defer d.Close()

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	log.Println("end...")
}
