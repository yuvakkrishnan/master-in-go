package main

import "fmt"

type song struct {
	Title, Artist string
}

type playlist struct {
	Genre string
	//

	songs []song
}

func main() {

	songs := []song{
		{Title: "Oh soldraya Mama", Artist: "Sanda"},
		{Title: "Oh soldraya Mama", Artist: "Manda"},
	}
	rock := playlist{
		Genre: "Munda", songs: songs,
	}
	fmt.Printf("%-20s %20s\n", "Title", "Artist")

	for _, s := range rock.songs {
		s.Title = "Mother Fucker!!!!"
		fmt.Printf("%-20s %20s\n", s.Title, s.Artist)

	}

	// fmt.Printf("%+v <<==", so)
}
