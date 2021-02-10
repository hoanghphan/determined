// This is a generated file.  Editing it will make you sad.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (b *BindMountV0) ParsedSchema() interface{} {
	return schemas.ParsedBindMountV0()
}

func (b *BindMountV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/bind-mount.json")
}

func (b *BindMountV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/bind-mount.json")
}

func (c *CheckpointStorageConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedCheckpointStorageConfigV0()
}

func (c *CheckpointStorageConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/checkpoint-storage.json")
}

func (c *CheckpointStorageConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/checkpoint-storage.json")
}

func (d *DataLayerGCSConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedDataLayerGCSConfigV0()
}

func (d *DataLayerGCSConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/data-layer-gcs.json")
}

func (d *DataLayerGCSConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/data-layer-gcs.json")
}

func (d *DataLayerS3ConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedDataLayerS3ConfigV0()
}

func (d *DataLayerS3ConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/data-layer-s3.json")
}

func (d *DataLayerS3ConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/data-layer-s3.json")
}

func (d *DataLayerSharedFSConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedDataLayerSharedFSConfigV0()
}

func (d *DataLayerSharedFSConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/data-layer-shared-fs.json")
}

func (d *DataLayerSharedFSConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/data-layer-shared-fs.json")
}

func (d *DataLayerConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedDataLayerConfigV0()
}

func (d *DataLayerConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/data-layer.json")
}

func (d *DataLayerConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/data-layer.json")
}

func (e *EnvironmentImageV0) ParsedSchema() interface{} {
	return schemas.ParsedEnvironmentImageV0()
}

func (e *EnvironmentImageV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/environment-image.json")
}

func (e *EnvironmentImageV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/environment-image.json")
}

func (e *EnvironmentVariablesV0) ParsedSchema() interface{} {
	return schemas.ParsedEnvironmentVariablesV0()
}

func (e *EnvironmentVariablesV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/environment-variables.json")
}

func (e *EnvironmentVariablesV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/environment-variables.json")
}

func (e *EnvironmentConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedEnvironmentConfigV0()
}

func (e *EnvironmentConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/environment.json")
}

func (e *EnvironmentConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/environment.json")
}

func (e *ExperimentConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedExperimentConfigV0()
}

func (e *ExperimentConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/experiment.json")
}

func (e *ExperimentConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/experiment.json")
}

func (g *GCSConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedGCSConfigV0()
}

func (g *GCSConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/gcs.json")
}

func (g *GCSConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/gcs.json")
}

func (h *HDFSConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedHDFSConfigV0()
}

func (h *HDFSConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hdfs.json")
}

func (h *HDFSConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hdfs.json")
}

func (c *CategoricalHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedCategoricalHyperparameterV0()
}

func (c *CategoricalHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json")
}

func (c *CategoricalHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json")
}

func (c *ConstHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedConstHyperparameterV0()
}

func (c *ConstHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-const.json")
}

func (c *ConstHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-const.json")
}

func (d *DoubleHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedDoubleHyperparameterV0()
}

func (d *DoubleHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-double.json")
}

func (d *DoubleHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-double.json")
}

func (i *IntHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedIntHyperparameterV0()
}

func (i *IntHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-int.json")
}

func (i *IntHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-int.json")
}

func (l *LogHyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedLogHyperparameterV0()
}

func (l *LogHyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-log.json")
}

func (l *LogHyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter-log.json")
}

func (h *HyperparameterV0) ParsedSchema() interface{} {
	return schemas.ParsedHyperparameterV0()
}

func (h *HyperparameterV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/hyperparameter.json")
}

func (h *HyperparameterV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/hyperparameter.json")
}

func (i *InternalConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedInternalConfigV0()
}

func (i *InternalConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/internal.json")
}

func (i *InternalConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/internal.json")
}

func (l *LengthV0) ParsedSchema() interface{} {
	return schemas.ParsedLengthV0()
}

func (l *LengthV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/length.json")
}

func (l *LengthV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/length.json")
}

func (o *OptimizationsConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedOptimizationsConfigV0()
}

func (o *OptimizationsConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/optimizations.json")
}

func (o *OptimizationsConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/optimizations.json")
}

func (r *ResourcesConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedResourcesConfigV0()
}

func (r *ResourcesConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/resources.json")
}

func (r *ResourcesConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/resources.json")
}

func (s *S3ConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedS3ConfigV0()
}

func (s *S3ConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/s3.json")
}

func (s *S3ConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/s3.json")
}

func (a *AdaptiveASHASearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedAdaptiveASHASearcherConfigV0()
}

func (a *AdaptiveASHASearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json")
}

func (a *AdaptiveASHASearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json")
}

func (a *AdaptiveSimpleSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedAdaptiveSimpleSearcherConfigV0()
}

func (a *AdaptiveSimpleSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive-simple.json")
}

func (a *AdaptiveSimpleSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive-simple.json")
}

func (a *AdaptiveSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedAdaptiveSearcherConfigV0()
}

func (a *AdaptiveSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive.json")
}

func (a *AdaptiveSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-adaptive.json")
}

func (a *AsyncHalvingSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedAsyncHalvingSearcherConfigV0()
}

func (a *AsyncHalvingSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-async-halving.json")
}

func (a *AsyncHalvingSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-async-halving.json")
}

func (g *GridSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedGridSearcherConfigV0()
}

func (g *GridSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-grid.json")
}

func (g *GridSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-grid.json")
}

func (p *PBTSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedPBTSearcherConfigV0()
}

func (p *PBTSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-pbt.json")
}

func (p *PBTSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-pbt.json")
}

func (r *RandomSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedRandomSearcherConfigV0()
}

func (r *RandomSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-random.json")
}

func (r *RandomSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-random.json")
}

func (s *SingleSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSingleSearcherConfigV0()
}

func (s *SingleSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-single.json")
}

func (s *SingleSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-single.json")
}

func (s *SyncHalvingSearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSyncHalvingSearcherConfigV0()
}

func (s *SyncHalvingSearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher-sync-halving.json")
}

func (s *SyncHalvingSearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher-sync-halving.json")
}

func (s *SearcherConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSearcherConfigV0()
}

func (s *SearcherConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/searcher.json")
}

func (s *SearcherConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/searcher.json")
}

func (s *SharedFSConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedSharedFSConfigV0()
}

func (s *SharedFSConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/shared-fs.json")
}

func (s *SharedFSConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/shared-fs.json")
}

func (t *TestRootV0) ParsedSchema() interface{} {
	return schemas.ParsedTestRootV0()
}

func (t *TestRootV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-root.json")
}

func (t *TestRootV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-root.json")
}

func (t *TestSubV0) ParsedSchema() interface{} {
	return schemas.ParsedTestSubV0()
}

func (t *TestSubV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-sub.json")
}

func (t *TestSubV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-sub.json")
}

func (t *TestUnionAV0) ParsedSchema() interface{} {
	return schemas.ParsedTestUnionAV0()
}

func (t *TestUnionAV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-union-a.json")
}

func (t *TestUnionAV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-union-a.json")
}

func (t *TestUnionBV0) ParsedSchema() interface{} {
	return schemas.ParsedTestUnionBV0()
}

func (t *TestUnionBV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-union-b.json")
}

func (t *TestUnionBV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-union-b.json")
}

func (t *TestUnionV0) ParsedSchema() interface{} {
	return schemas.ParsedTestUnionV0()
}

func (t *TestUnionV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/test-union.json")
}

func (t *TestUnionV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/test-union.json")
}
