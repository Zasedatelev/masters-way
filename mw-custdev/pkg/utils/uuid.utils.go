package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const DEFAULT_STRING_LAYOUT = "2006-01-02T15:04:05.000Z07:00"

func ToNullUuid(someString string) uuid.NullUUID {
	parseUuid, err := uuid.Parse(someString)

	nullUuid := uuid.NullUUID{
		UUID:  parseUuid,
		Valid: false,
	}

	if err != nil {
		nullUuid.Valid = true
	}

	return nullUuid
}

func ConvertPgUUIDToUUID(pgUUID pgtype.UUID) uuid.UUID {
	var uuid uuid.UUID
	copy(uuid[:], pgUUID.Bytes[:])
	return uuid
}

func MarshalPgUUID(pgUUID pgtype.UUID) *string {
	if pgUUID.Valid {
		u, err := uuid.FromBytes(pgUUID.Bytes[:])
		if err == nil {
			str := u.String()
			return &str
		}
	}
	return nil
}

func MarshalPgTimestamp(timestamp pgtype.Timestamp) *string {
	if timestamp.Valid {
		str := timestamp.Time.Format(DEFAULT_STRING_LAYOUT)
		return &str
	} else {
		return nil
	}
}

func MarshalPgText(pgText pgtype.Text) *string {
	if pgText.Valid {
		str := pgText.String
		return &str
	}
	return nil
}
