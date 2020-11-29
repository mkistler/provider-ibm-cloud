package resourcekey

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reference"

	"github.com/IBM-Cloud/bluemix-go/crn"
	rcv2 "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"

	"github.com/crossplane-contrib/provider-ibm-cloud/apis/resourcecontrollerv2/v1alpha1"
	ibmc "github.com/crossplane-contrib/provider-ibm-cloud/pkg/clients"
)

const (
	// StateActive represents a service instance in a running, available, and ready state
	StateActive = "active"
	// StateInactive represents a service instance in a not running state
	StateInactive = "inactive"
)

// LateInitializeSpec fills optional and unassigned fields with the values in *rcv2.ResourceKey object.
func LateInitializeSpec(client ibmc.ClientSession, spec *v1alpha1.ResourceKeyParameters, in *rcv2.ResourceKey) error { // nolint:gocyclo
	if spec.Role == nil {
		spec.Role = in.Role
	}
	return nil
}

// GenerateCreateResourceKeyOptions produces CreateResourceKeyOptions object from ResourceKeyParameters object.
func GenerateCreateResourceKeyOptions(client ibmc.ClientSession, in v1alpha1.ResourceKeyParameters, o *rcv2.CreateResourceKeyOptions) error {
	o.Name = reference.ToPtrValue(in.Name)
	o.Source = in.Source
	o.Parameters = in.Parameters
	o.Role = in.Role
	return nil
}

// GenerateUpdateResourceKeyOptions produces UpdateResourceKeyOptions object from ResourceKeyParameters object.
func GenerateUpdateResourceKeyOptions(client ibmc.ClientSession, id string, in v1alpha1.ResourceKeyParameters, o *rcv2.UpdateResourceKeyOptions) error {
	o.ID = reference.ToPtrValue(id)
	o.Name = reference.ToPtrValue(in.Name)
	return nil
}

// GenerateObservation produces ResourceKeyObservation object from *rcv2.ResourceKey object.
func GenerateObservation(client ibmc.ClientSession, in *rcv2.ResourceKey) (v1alpha1.ResourceKeyObservation, error) {
	o := v1alpha1.ResourceKeyObservation{
		ID:                  reference.FromPtrValue(in.ID),
		Guid:                reference.FromPtrValue(in.Guid),
		Crn:                 reference.FromPtrValue(in.Crn),
		URL:                 reference.FromPtrValue(in.URL),
		AccountID:           reference.FromPtrValue(in.AccountID),
		ResourceGroupID:     reference.FromPtrValue(in.ResourceGroupID),
		SourceCrn:           reference.FromPtrValue(in.SourceCrn),
		State:               reference.FromPtrValue(in.State),
		IamCompatible:       ibmc.BoolValue(in.IamCompatible),
		ResourceInstanceURL: reference.FromPtrValue(in.ResourceInstanceURL),
		CreatedAt:           ibmc.DateTimeToMetaV1Time(in.CreatedAt),
		UpdatedAt:           ibmc.DateTimeToMetaV1Time(in.UpdatedAt),
		DeletedAt:           ibmc.DateTimeToMetaV1Time(in.DeletedAt),
		CreatedBy:           reference.FromPtrValue(in.CreatedBy),
		UpdatedBy:           reference.FromPtrValue(in.UpdatedBy),
		DeletedBy:           reference.FromPtrValue(in.DeletedBy),
	}
	// ServiceEndpoints can be found in instance.Parameters["service-endpoints"]
	return o, nil
}

// IsUpToDate checks whether current state is up-to-date compared to the given set of parameters.
func IsUpToDate(client ibmc.ClientSession, in *v1alpha1.ResourceKeyParameters, observed *rcv2.ResourceKey, l logging.Logger) (bool, error) {
	desired := in.DeepCopy()
	actual, err := GenerateResourceKeyParameters(client, observed)
	if err != nil {
		return false, err
	}

	l.Info(cmp.Diff(desired, actual))

	return cmp.Equal(desired, actual, cmpopts.EquateEmpty(), cmpopts.IgnoreFields(v1alpha1.ResourceKeyParameters{})), nil
}

// GenerateResourceKeyParameters generates service instance parameters from resource instance
func GenerateResourceKeyParameters(client ibmc.ClientSession, in *rcv2.ResourceKey) (*v1alpha1.ResourceKeyParameters, error) {

	o := &v1alpha1.ResourceKeyParameters{
		Name: reference.FromPtrValue(in.Name),
		Role: in.Role,
	}
	return o, nil
}
