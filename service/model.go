package service

type Chronicle struct {
	Debtor Debtor
}

// Debtor struct
type Debtor struct {
	ID                  uint               `json:"id"`
	Name                string             `json:"name,omitempty"`                                                                                // Должник (наименование)
	FullName            string             `json:"full_name,omitempty"`                                                                          // Должник (полное наименование)
}