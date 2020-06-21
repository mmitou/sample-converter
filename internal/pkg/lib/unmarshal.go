package lib

import (
	"encoding/json"
	"strings"

	"golang.org/x/xerrors"
)

type BaseObject struct {
	PK string
	SK string
}

func Unmarshal(bs []byte) (*User, *Reaction, *Friendship, *Photo, error) {
	base := &BaseObject{}
	if err := json.Unmarshal(bs, base); err != nil {
		return nil, nil, nil, nil, err
	}
	if strings.HasPrefix(base.SK, "#METADATA#") {
		u, err := ToUser(bs)
		return u, nil, nil, nil, err
	}
	if strings.HasPrefix(base.SK, "#FRIEND#") {
		fship, err := ToFriendship(bs)
		return nil, nil, fship, nil, err
	}
	if strings.HasPrefix(base.PK, "REACTION#") {
		rctn, err := ToReaction(bs)
		return nil, rctn, nil, nil, err
	}
	if strings.HasPrefix(base.PK, "USER") && strings.HasPrefix(base.SK, "PHOTO") {
		photo, err := ToPhoto(bs)
		return nil, nil, nil, photo, err
	}

	return nil, nil, nil, nil, xerrors.New("Unknown Type")
}

func ToUser(bs []byte) (*User, error) {
	u := &User{}
	err := json.Unmarshal(bs, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ToReaction(bs []byte) (*Reaction, error) {
	r := &Reaction{}
	err := json.Unmarshal(bs, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func ToFriendship(bs []byte) (*Friendship, error) {
	f := &Friendship{}
	err := json.Unmarshal(bs, f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ToPhoto(bs []byte) (*Photo, error) {
	p := &Photo{}
	err := json.Unmarshal(bs, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
