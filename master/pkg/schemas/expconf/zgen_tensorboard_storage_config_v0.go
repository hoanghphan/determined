// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (t TensorboardStorageConfigV0) GetUnionMember() interface{} {
	if t.SharedFSConfigV0 != nil {
		return nil
	}
	if t.HDFSConfig != nil {
		return nil
	}
	if t.S3Config != nil {
		return nil
	}
	if t.GCSConfig != nil {
		return nil
	}
	panic("no union member defined")
}

func (t TensorboardStorageConfigV0) GetSaveExperimentBest() int {
	if t.SharedFSConfigV0 != nil {
		return t.SharedFSConfigV0.GetSaveExperimentBest()
	}
	if t.HDFSConfig != nil {
		return t.HDFSConfig.GetSaveExperimentBest()
	}
	if t.S3Config != nil {
		return t.S3Config.GetSaveExperimentBest()
	}
	if t.GCSConfig != nil {
		return t.GCSConfig.GetSaveExperimentBest()
	}
	panic("no union member defined")
}

func (t TensorboardStorageConfigV0) GetSaveTrialBest() int {
	if t.SharedFSConfigV0 != nil {
		return t.SharedFSConfigV0.GetSaveTrialBest()
	}
	if t.HDFSConfig != nil {
		return t.HDFSConfig.GetSaveTrialBest()
	}
	if t.S3Config != nil {
		return t.S3Config.GetSaveTrialBest()
	}
	if t.GCSConfig != nil {
		return t.GCSConfig.GetSaveTrialBest()
	}
	panic("no union member defined")
}

func (t TensorboardStorageConfigV0) GetSaveTrialLatest() int {
	if t.SharedFSConfigV0 != nil {
		return t.SharedFSConfigV0.GetSaveTrialLatest()
	}
	if t.HDFSConfig != nil {
		return t.HDFSConfig.GetSaveTrialLatest()
	}
	if t.S3Config != nil {
		return t.S3Config.GetSaveTrialLatest()
	}
	if t.GCSConfig != nil {
		return t.GCSConfig.GetSaveTrialLatest()
	}
	panic("no union member defined")
}

func (t TensorboardStorageConfigV0) WithDefaults() TensorboardStorageConfigV0 {
	return schemas.WithDefaults(t).(TensorboardStorageConfigV0)
}

func (t TensorboardStorageConfigV0) Merge(other TensorboardStorageConfigV0) TensorboardStorageConfigV0 {
	return schemas.Merge(t, other).(TensorboardStorageConfigV0)
}

func (t TensorboardStorageConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedTensorboardStorageConfigV0()
}

func (t TensorboardStorageConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/tensorboard-storage.json")
}

func (t TensorboardStorageConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/tensorboard-storage.json")
}