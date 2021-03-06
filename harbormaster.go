package gonduit

import (
	"github.com/uber/gonduit/requests"
	"github.com/uber/gonduit/responses"
)

// HarbormasterBuildableSearchMethod is method name on Phabricator API.
const HarbormasterBuildableSearchMethod = "harbormaster.buildable.search"

// HarbormasterBuildableSearch performs a call to harbormaster.builable.search.
func (c *Conn) HarbormasterBuildableSearch(
	req requests.HarbormasterBuildableSearchRequest,
) (*responses.HarbormasterBuildableSearchResponse, error) {
	var res responses.HarbormasterBuildableSearchResponse

	if err := c.Call(HarbormasterBuildableSearchMethod, &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
