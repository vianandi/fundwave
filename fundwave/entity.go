package fundwave

type Fundwave struct {
	ID                int    `json:"id"`
	UserID            int    `json:"userid"`
	Name              string `json:"name"`
	Short_description string `json:"short_description"`
	Perks             string `json:"perks"`
	Becker_count      int    `json:"becker_count"`
	Goal_amount       int    `json:"goal_amount"`
	Current_amount    int    `json:"current_amount"`
	Slug              string `json:"slug"`
	FundwaveImages    []FundwaveImages
}

type FundwaveImages struct {
	ID         int    `json:"id"`
	FundwaveID int    `json:"fundwave_id"`
	Filename   string `json:"file_name"`
	IsPrimary  int    `json:"is_primary"`
}