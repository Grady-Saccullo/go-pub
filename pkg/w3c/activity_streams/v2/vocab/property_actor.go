package vocab

import json_ld_v1_vocab "github.com/Grady-Saccullo/go-pub/pkg/w3c/json_ld/v1/vocab"

type PropertyActor interface {
	json_ld_v1_vocab.TypeIRI

	Object
	Object
}
