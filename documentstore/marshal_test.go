package documentstore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestMarshalUnmarshal_RoundTrip(t *testing.T) {
	in := testUser{ID: "0", Name: "luke_skywalker"}

	doc, err := MarshalDocument(in)
	require.NoError(t, err)
	require.NotNil(t, doc)

	var out testUser
	err = UnmarshalDocument(doc, &out)
	require.NoError(t, err)

	require.Equal(t, in, out)
}

func TestUnmarshalDocument_Validation(t *testing.T) {
	var out testUser

	err := UnmarshalDocument(nil, &out)
	require.ErrorIs(t, err, ErrNilDocument)

	err = UnmarshalDocument(&Document{Fields: map[string]DocumentField{}}, nil)
	require.ErrorIs(t, err, ErrNilOutput)
}
