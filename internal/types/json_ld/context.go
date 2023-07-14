package json_ld

func GetJsonLDContext(ldAlias map[string]string, ns string) *string {
	if s, ok := ldAlias[ns]; ok {
		return &s
	}

	return nil
}
