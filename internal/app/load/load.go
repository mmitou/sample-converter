package load

import (
	"bufio"
	"encoding/json"
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
	ufct := lib.NewUserFacet(users)
	bs, err := json.Marshal(ufct)
	fmt.Println(string(bs))
}
