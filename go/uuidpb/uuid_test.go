package uuidpb

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func assertProtoEqual(t *testing.T, expected proto.Message, actual proto.Message) {
	assert.Truef(t, proto.Equal(expected, actual), cmp.Diff(expected, actual, protocmp.Transform()))
}

func TestNew(t *testing.T) {
	id := uuid.MustParse("e63ebbbb-d54e-4598-8ceb-58fe1a353675")
	idProto := New(id)
	expected := &UUID{
		Hi: 16590904492691441048,
		Lo: 10154307633221547637,
	}
	assertProtoEqual(t, expected, idProto)
}

func TestToUUID(t *testing.T) {
	idProto := &UUID{
		Hi: 16590904492691441048,
		Lo: 10154307633221547637,
	}
	id := idProto.ToUUID()
	expected := uuid.MustParse("e63ebbbb-d54e-4598-8ceb-58fe1a353675")
	assert.Equal(t, expected, id)
}
