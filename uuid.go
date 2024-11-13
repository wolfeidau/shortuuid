package shortuuid

import (
	"bytes"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

const (
	uuidLen = 16
)

// MustShorten take one or more uuids and encode them using base58
// with the encoded values being appended in order.
func MustShorten(uuidv4s ...string) string {

	short, err := Shorten(uuidv4s...)
	if err != nil {
		panic(err)
	}

	return short
}

// Shorten take one or more uuids and encode them using base58
// with the encoded values being appended in order.
func Shorten(uuidv4s ...string) (string, error) {
	buf := new(bytes.Buffer)

	for _, u := range uuidv4s {
		puid, err := uuid.Parse(u)
		if err != nil {
			return "", fmt.Errorf("failed to parse %s as a uuid: %w", u, err)
		}

		data, err := puid.MarshalBinary()
		if err != nil {
			return "", fmt.Errorf("failed to marshall %s as a uuid: %w", u, err)
		}

		_, err = buf.Write(data)
		if err != nil {
			return "", fmt.Errorf("failed to write %s as a uuid: %w", u, err)
		}
	}

	return base58.Encode(buf.Bytes()), nil
}

// MustUnShorten takes a shortened value and decodes it into on or more uuids with the result
// being in the original order.
func MustUnShorten(val string) []string {
	uuids, err := UnShorten(val)
	if err != nil {
		panic(err)
	}

	return uuids
}

// UnShorten takes a shortened value and decodes it into on or more uuids with the result
// being in the original order.
func UnShorten(val string) ([]string, error) {

	var uuids []string

	data, err := base58.Decode(val)
	if err != nil {
		return nil, fmt.Errorf("failed to decode value: %w", err)
	}

	if len(data) == 0 || len(data)%uuidLen != 0 {
		return nil, fmt.Errorf("failed to decode value invalid decoded length: %d", len(data))
	}

	buf := make([]byte, 0, uuidLen)
	r := bytes.NewBuffer(data)
	for {
		n, err := io.ReadFull(r, buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				return nil, fmt.Errorf("failed to read uuid value: %w", err)
			}
		}

		u, err := uuid.FromBytes(buf)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s as a uuid: %w", u, err)
		}

		uuids = append(uuids, u.String())
	}

	return uuids, nil
}
