package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var idlerlog = logf.Log.WithName("idler-resource")

func (r *Idler) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,delete=/validate-codeready-toolchain-v1alpha1-idler,mutating=false,failurePolicy=fail,groups=toolchain.dev.openshift.com,resources=idlers,versions=v1alpha1,name=vidler.kb.io

var _ webhook.Validator = &Idler{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Idler) ValidateCreate() error {
	idlerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Idler) ValidateUpdate(old runtime.Object) error {
	idlerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Idler) ValidateDelete() error {
	idlerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
