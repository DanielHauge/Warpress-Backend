package bnet

// ProfileService has OAuth Profile APIs. See Client.
type ProfileService struct {
	client *Client
}

type WoWProfile struct {
	Characters []WOWCharacter `json:"characters"`
}

func (s *ProfileService) WOW() (*WoWProfile, *Response, error) {
	req, err := s.client.NewRequest("GET", "wow/user/characters", nil)
	if err != nil {
		return nil, nil, err
	}

	var profile WoWProfile
	resp, err := s.client.Do(req, &profile)
	if err != nil {
		return nil, resp, err
	}

	return &profile, resp, nil
}
