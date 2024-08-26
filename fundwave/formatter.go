package fundwave

type FundwaveFormatter struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	Name              string `json:"name"`
	Short_description string `json:"short_description"`
	ImageURL          string `json:"image_url"`
	Goal_amount        int    `json:"goal_amount"`
	Current_amount     int    `json:"current_amount"`
}

func FormmatFundwave(fundwave Fundwave) FundwaveFormatter {
	// var fundwaveFormatter FundwaveFormatter
	fundwaveFormatter := FundwaveFormatter{}
	fundwaveFormatter.ID = fundwave.ID
	fundwaveFormatter.UserID = fundwave.UserID
	fundwaveFormatter.Name = fundwave.Name
	fundwaveFormatter.Short_description = fundwave.Short_description
	// fundwaveFormatter.ImageURL = fundwave.ImageURL
	fundwaveFormatter.Goal_amount = fundwave.Goal_amount
	fundwaveFormatter.Current_amount = fundwave.Current_amount
	fundwaveFormatter.ImageURL = ""

	if len(fundwave.FundwaveImages) > 0 {
		fundwaveFormatter.ImageURL = fundwave.FundwaveImages[0].Filename
	}

	return fundwaveFormatter
}

func FormatFundwaves(fundwaves []Fundwave) []FundwaveFormatter {
	if len(fundwaves) == 0 {
		return []FundwaveFormatter{}
	}

	var fundwaveFormatters []FundwaveFormatter

	for _, fundwave := range fundwaves {
		fundwaveFormatter := FormmatFundwave(fundwave)
		fundwaveFormatters = append(fundwaveFormatters, fundwaveFormatter)
	}

	return fundwaveFormatters
}