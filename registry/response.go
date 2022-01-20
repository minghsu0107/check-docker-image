package registry

type tagsResponse struct {
	Tags []string `json:"tags"`
}

func (registry *Registry) Tags(imageName, imageTag string) (bool, error) {
	url := registry.url("/v2/%s/tags/list", imageName)

	var response tagsResponse
	var err error
	for {
		url, err = registry.getPaginatedJSON(url, &response)
		switch err {
		case ErrNoMorePages:
			for _, remoteTag := range response.Tags {
				if remoteTag == imageTag {
					return true, nil
				}
			}
			return false, nil
		case nil:
			for _, remoteTag := range response.Tags {
				if remoteTag == imageTag {
					return true, nil
				}
			}
			continue
		default:
			return false, err
		}
	}
}
