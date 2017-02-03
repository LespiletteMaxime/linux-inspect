package psn

import (
	"fmt"
	"testing"
	"time"
)

func TestRunTopDefault(t *testing.T) {
	o, err := RunTopDefault(0)
	if err != nil {
		t.Skip(err)
	}
	fmt.Println(o)
}

func Test_parseTopCommandKiB(t *testing.T) {
	bts, bs, err := parseTopCommandKiB("50.883g")
	if err != nil {
		t.Fatal(err)
	}
	if bts != 53687091200 {
		t.Fatalf("bytes expected %d, got %d", 53687091200, bts)
	}
	if bs != "54 GB" {
		t.Fatalf("humanized bytes expected '54 GB', got %q", bs)
	}
}

func TestGetTopDefault(t *testing.T) {
	now := time.Now()
	rows, err := GetTopDefault(0)
	if err != nil {
		t.Skip(err)
	}
	fmt.Printf("found %d entrines in %v", len(rows), time.Since(now))
	fmt.Println(rows[0])
}
