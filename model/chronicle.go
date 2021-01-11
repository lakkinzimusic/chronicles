package model

import (
	pb "github.com/lakkinzimusic/chronicles/proto"
	"github.com/golang/protobuf/ptypes"
	"time"
)

// Chronicle struct
type Chronicle struct {
	Date time.Time
	Debtors []*Debtor
}

func MarshalChronicle(chronicle *pb.Chronicle) *Chronicle {
	date, _ := ptypes.Timestamp(chronicle.Date)
	return &Chronicle{
		Date: date,
		Debtors: MarshalDebtor(chronicle.Debtors),
	}
}

func UnmarshalChronicle(chronicle *Chronicle) *pb.Chronicle {
	date, _ := ptypes.TimestampProto(chronicle.Date)
	return &pb.Chronicle{
		Date: date,
		Debtors: UnmarshalDebtor(chronicle.Debtors),
	}
}