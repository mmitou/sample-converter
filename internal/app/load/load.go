package load

import (
	"bufio"
	"fmt"

	"os"

	"github.com/mmitou/sample/internal/pkg/lib"
)

func Run() {
	fp, err := os.Open("./items.json")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var users []*lib.User
	var rctns []*lib.Reaction
	var fships []*lib.Friendship
	var photos []*lib.Photo
	for scanner.Scan() {
		str := scanner.Text()
		bs := []byte(str)
		user, rctn, fship, photo, err := lib.Unmarshal(bs)
		switch {
		case user != nil:
			users = append(users, user)
		case rctn != nil:
			rctns = append(rctns, rctn)
		case fship != nil:
			fships = append(fships, fship)
		case photo != nil:
			photos = append(photos, photo)
		default:
			fmt.Printf("lib.Unmarshal(%s): %+v\n", str, err)
		}
	}
	fmt.Printf("%+v\n", users[0])
	fmt.Printf("%+v\n", rctns[0])
	fmt.Printf("%+v\n", fships[0])
	fmt.Printf("%+v\n", photos[0])
}
