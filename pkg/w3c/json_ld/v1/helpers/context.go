package helpers

func GetJsonLDContext(ldAlias map[string]string, ns string) *string {
	if s, ok := ldAlias[ns]; ok {
		return &s
	}

	return nil
}
