package tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestBuild(t *testing.T) {
	r := Record{ID: 0}
	tree := &Tree{}
	tree.Build(r)
	fmt.Println("ROOT: ", tree.Root)

}

func TestBuildRandom64(t *testing.T) {
	rcds := makeRandomRecords(64)
	tree := &Tree{}
	tree.Build(rcds...)
	// spew.Dump(tree)
	tree.Print()
}

func makeRandomRecords(length int) []Record {
	records := make([]Record, 0)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < length; i++ {
		records = append(records, Record{ID: r1.Int63n(int64(length))})
	}
	spew.Dump(records)
	fmt.Println("LEN OF RECORDS====", len(records))
	return records
}
