package uuidpb

import (
	"encoding/binary"

	"github.com/google/uuid"
)

func New(id uuid.UUID) *UUID {
	return &UUID{
		Hi: binary.BigEndian.Uint64(id[:8]),
		Lo: binary.BigEndian.Uint64(id[8:]),
	}
}

func (idProto *UUID) ToUUID() uuid.UUID {
	var id uuid.UUID
	binary.BigEndian.PutUint64(id[:8], idProto.GetHi())
	binary.BigEndian.PutUint64(id[8:], idProto.GetLo())
	return id
}
