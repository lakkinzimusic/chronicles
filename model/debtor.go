package model


import (
	pb "github.com/lakkinzimusic/chronicles/proto"
)
// Debtor struct
type Debtor struct {
	ID                  uint               `json:"id"`
	Name                string             `json:"name,omitempty"`                                                                                // Должник (наименование)
	FullName            string            `json:"full_name,omitempty"`                                                                          // Должник (полное наименование)
}

func MarshalDebtor(debtors []*pb.Debtor)  []*Debtor {
	var result []*Debtor
	for _, item := range debtors{
		result = append(result, &Debtor{
			ID: uint(item.ID),
			Name: item.Name,
			FullName: item.FullName,
		})
	}
	return result
}

func UnmarshalDebtor(debtors []*Debtor)  []*pb.Debtor {
	var result []*pb.Debtor
	for _, item := range debtors{
		result = append(result, &pb.Debtor{
			ID: uint32(item.ID),
			Name: item.Name,
			FullName: item.FullName,
		})
	}
	return result
}