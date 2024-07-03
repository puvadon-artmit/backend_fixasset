package model

type ItemPostingGroups struct {
	Odata_context string                    `json:"@odata.context"`
	Value         []GenProductPostingGroups `json:"value"`
}

type GenProductPostingGroups struct {
	GenProductPostingGroupsID  string  `gorm:"type:uuid;primaryKey" json:"gen_product_posting_groups_id"`
	Code                       *string `json:"Code"`
	Description                *string `json:"Description"`
	Def_VAT_Prod_Posting_Group *string `json:"Def_VAT_Prod_Posting_Group"`
	Auto_Insert_Default        *bool   `json:"Auto_Insert_Default"`
}

type ItemPostingGroupMaliwan struct {
	Odata_context string                          `json:"@odata.context"`
	Value         []GenProductPostingGroupMaliwan `json:"value"`
}

type GenProductPostingGroupMaliwan struct {
	Code                       *string `json:"Code"`
	Description                *string `json:"Description"`
	Def_VAT_Prod_Posting_Group *string `json:"Def_VAT_Prod_Posting_Group"`
	Auto_Insert_Default        *bool   `json:"Auto_Insert_Default"`
}
