package activity_streams_v2_impl

import (
	"github.com/Grady-Saccullo/activity-pub-go/pkg/w3c/activity_streams/v2/vocab"
)

type ActorProperties struct {
	ObjectProperties
}

func deserializeActorProperties(d map[string]interface{}, ldAliases map[string]string, i activity_streams_v2_vocab.ActorPropertySetters) error {
	if err := deserializeObjectProperties(d, ldAliases, i); err != nil {
		return err
	}

	return nil
}
