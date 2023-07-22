package impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/internal/types/w3c/activity_streams/vocab"
)

type ActorProperties struct {
	ObjectProperties
}

func deserializeActorProperties(d map[string]interface{}, ldAliases map[string]string, i vocab.ActorPropertySetters) error {
	if err := deserializeObjectProperties(d, ldAliases, i); err != nil {
		return err
	}

	return nil
}
