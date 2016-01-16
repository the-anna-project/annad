package crypto

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func DecryptGPGBytesWithPass(raw, pass []byte) ([]byte, error) {
	decbuf := bytes.NewBuffer(raw)
	result, err := armor.Decode(decbuf)
	if err != nil {
		return nil, maskAny(err)
	}

	md, err := openpgp.ReadMessage(result.Body, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		return pass, nil
	}, nil)
	if err != nil {
		return nil, maskAny(err)
	}

	b, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, maskAny(err)
	}

	return b, nil
}

func EncryptGPGFileWithPass(raw, pass []byte) ([]byte, error) {
	encbuf := bytes.NewBuffer(nil)

	w, err := armor.Encode(encbuf, "PGP SIGNATURE", nil)
	if err != nil {
		return nil, maskAny(err)
	}

	plaintext, err := openpgp.SymmetricallyEncrypt(w, pass, nil, nil)
	if err != nil {
		return nil, maskAny(err)
	}

	_, err = plaintext.Write(raw)
	if err != nil {
		return nil, maskAny(err)
	}

	plaintext.Close()
	w.Close()

	return encbuf.Bytes(), nil
}
