package demo

import (
	cipher "github.com/skycoin/skywire/pkg/cipher"
)

func publicKey(hex string) cipher.PubKey {
	s := cipher.PubKey{}
	s.Set(hex)
	return s
}

var publicKeys = map[string]cipher.PubKey{
	"host1": publicKey("0327f6951690898c10a02f793e367b6349492bfa266b696a81005cb62bf0f32e52"),
	"host2": publicKey("03477919f99e6d4ab2d0f002fd2313562cb0678bbd2e980c32a2c7e4a2483dab5e"),
}
