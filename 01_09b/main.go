package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
	AlbumCount, SongCount int
}

type playlist []Song

func (h playlist) Len() int{
	return len(h)
}

func (h playlist) Less(i,j int) bool{
	return h[i].PlayCount>h[j].PlayCount
}

func (h playlist) Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}

func (h *playlist) Push(i any) {
	*h = append(*h, i.(Song))
}

func (h *playlist) Pop() any{
	n:=len(*h)
	o:=*h
	t:= o[n-1]
	*h=o[:n-1]
	return t
}


// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	var list []Song
	ph:= &playlist{}
	if len(albums)==0{
		return list
	}

	heap.Init(ph)

	for i,album := range albums{
		for j,song := range album{
			song.AlbumCount,song.SongCount = i,j
			heap.Push(ph,song)
		}
	}

	for ph.Len()!=0{
		i:=heap.Pop(ph)
		t:=i.(Song)
		list = append(list, t)
		fmt.Println(list)
		// if t.SongCount<len(albums[t.AlbumCount])-1{
		// 	ns:=albums[t.AlbumCount][t.SongCount+1]
		// 	ns.AlbumCount,ns.SongCount = t.AlbumCount,t.SongCount+1
		// 	heap.Push(ph,ns)
		// }
	}
	return list
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
