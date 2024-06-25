package v1alpha1

const (
	// FeatureAnnotationKeyPrefix is used in feature annotation keys in Space, NSTemplate and other resources
	// to refer to the corresponding feature toggle from the configuration:
	// "toolchain.dev.openshift.com/feature/<feature-name>
	FeatureAnnotationKeyPrefix = FeatureToggleNameAnnotationKey + "/"

	// FeatureToggleNameAnnotationKey is used for referring tier template objects to feature toggles defined in configuration
	FeatureToggleNameAnnotationKey = LabelKeyPrefix + "feature"
)
