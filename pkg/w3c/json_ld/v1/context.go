package json_ld_v1

func GetJsonLDContext(ldAlias map[string]string, ns string) *string {
	if s, ok := ldAlias[ns]; ok {
		return &s
	}

	return nil
}
