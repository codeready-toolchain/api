package v1alpha1

const (
	// AnnotationKeyPrefix is the prefix used for annotation key values
	AnnotationKeyPrefix = LabelKeyPrefix

	// UserIDUserAnnotationKey is used to set an annotation value in the User resource on the member cluster, that
	// contains the user's User ID as set in the user's JWT token.
	UserIDUserAnnotationKey = AnnotationKeyPrefix + "sso-user-id"

	// AccountIDUserAnnotationKey is used to set an annotation value in the User resource on the member cluster, that
	// contains the user's Account ID as set in the user's JWT token.
	AccountIDUserAnnotationKey = AnnotationKeyPrefix + "sso-account_id"

	// EmailUserAnnotationKey is used to set an annotation value in the User resource on the member cluster, that
	// contains the user's Email as set in the user's JWT token.
	EmailUserAnnotationKey = AnnotationKeyPrefix + "email"
)
