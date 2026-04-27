package service

// resolveOpenAIForwardModel determines the upstream model for OpenAI-compatible
// forwarding. Group-level default mapping only applies when the account itself
// did not match any explicit model_mapping rule.
func resolveOpenAIForwardModel(account *Account, requestedModel, defaultMappedModel string) string {
	if account == nil {
		if defaultMappedModel != "" {
			return defaultMappedModel
		}
		return requestedModel
	}

	mappedModel, matched := account.ResolveMappedModel(requestedModel)
	if !matched && defaultMappedModel != "" {
		return defaultMappedModel
	}
	return mappedModel
}

func resolveOpenAICompactForwardModel(account *Account, requestedModel string) string {
	if account == nil {
		return requestedModel
	}
	mappedModel, matched := account.ResolveCompactMappedModel(requestedModel)
	if !matched {
		return requestedModel
	}
	return mappedModel
}
