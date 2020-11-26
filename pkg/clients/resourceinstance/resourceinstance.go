package resourceinstance

import (
	"strconv"

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
	// StatePendingReclamation means that delete has been initiated
	StatePendingReclamation = "pending_reclamation"
	errGetResPlaID          = "error getting resource plan ID"
	errGetResGroupID        = "error getting resource group ID"
)

// LateInitializeSpec fills optional and unassigned fields with the values in *rcv2.ResourceInstance object.
func LateInitializeSpec(client ibmc.ClientSession, spec *v1alpha1.ResourceInstanceParameters, in *rcv2.ResourceInstance) error { // nolint:gocyclo
	if spec.AllowCleanup == nil {
		spec.AllowCleanup = in.AllowCleanup
	}
	if spec.Parameters == nil {
		spec.Parameters = ibmc.MapToRawExtension(in.Parameters)
	}
	if spec.EntityLock == nil {
		spec.EntityLock = reference.ToPtrValue(strconv.FormatBool(*in.Locked))
	}
	if spec.Tags == nil {
		tags, err := ibmc.GetResourceInstanceTags(client, reference.FromPtrValue(in.TargetCrn))
		if err != nil {
			return err
		}
		spec.Tags = tags
	}
	return nil
}

// GenerateCreateResourceInstanceOptions produces CreateResourceInstanceOptions object from ResourceInstanceParameters object.
func GenerateCreateResourceInstanceOptions(client ibmc.ClientSession, in v1alpha1.ResourceInstanceParameters, o *rcv2.CreateResourceInstanceOptions) error {
	rgID, err := ibmc.GetResourceGroupID(client, in.ResourceGroupName)
	if err != nil {
		return errors.Wrap(err, errGetResGroupID)
	}

	rPlanID, err := ibmc.GetResourcePlanID(client, in.ServiceName, in.ResourcePlanName)
	if err != nil {
		return errors.Wrap(err, errGetResPlaID)
	}

	o.Name = reference.ToPtrValue(in.Name)
	o.Target = reference.ToPtrValue(in.Target)
	o.ResourceGroup = rgID
	o.ResourcePlanID = rPlanID
	o.Tags = in.Tags
	o.AllowCleanup = in.AllowCleanup
	o.Parameters = ibmc.RawExtensionToMap(in.Parameters)
	o.EntityLock = in.EntityLock
	return nil
}

// GenerateUpdateResourceInstanceOptions produces UpdateResourceInstanceOptions object from ResourceInstanceParameters object.
func GenerateUpdateResourceInstanceOptions(client ibmc.ClientSession, id string, in v1alpha1.ResourceInstanceParameters, o *rcv2.UpdateResourceInstanceOptions) error {
	rPlanID, err := ibmc.GetResourcePlanID(client, in.ServiceName, in.ResourcePlanName)
	if err != nil {
		return errors.Wrap(err, errGetResPlaID)
	}

	o.ID = reference.ToPtrValue(id)
	o.Name = reference.ToPtrValue(in.Name)
	o.Parameters = ibmc.RawExtensionToMap(in.Parameters)
	o.ResourcePlanID = rPlanID
	o.AllowCleanup = in.AllowCleanup
	return nil
}

// GenerateObservation produces ResourceInstanceObservation object from *rcv2.ResourceInstance object.
func GenerateObservation(client ibmc.ClientSession, in *rcv2.ResourceInstance) (v1alpha1.ResourceInstanceObservation, error) {
	o := v1alpha1.ResourceInstanceObservation{
		ID:                  reference.FromPtrValue(in.ID),
		GUID:                reference.FromPtrValue(in.Guid),
		Crn:                 reference.FromPtrValue(in.Crn),
		URL:                 reference.FromPtrValue(in.URL),
		AccountID:           reference.FromPtrValue(in.AccountID),
		ResourceGroupID:     reference.FromPtrValue(in.ResourceGroupID),
		ResourceGroupCrn:    reference.FromPtrValue(in.ResourceGroupCrn),
		ResourceID:          reference.FromPtrValue(in.ResourceID),
		ResourcePlanID:      reference.FromPtrValue(in.ResourcePlanID),
		TargetCrn:           reference.FromPtrValue(in.TargetCrn),
		State:               reference.FromPtrValue(in.State),
		Type:                reference.FromPtrValue(in.Type),
		SubType:             reference.FromPtrValue(in.SubType),
		LastOperation:       ibmc.MapToRawExtension(in.LastOperation),
		DashboardURL:        reference.FromPtrValue(in.DashboardURL),
		PlanHistory:         GeneratePlanHistory(in.PlanHistory),
		ResourceAliasesURL:  reference.FromPtrValue(in.ResourceAliasesURL),
		ResourceBindingsURL: reference.FromPtrValue(in.ResourceBindingsURL),
		ResourceKeysURL:     reference.FromPtrValue(in.ResourceKeysURL),
		CreatedAt:           ibmc.DateTimeToMetaV1Time(in.CreatedAt),
		CreatedBy:           reference.FromPtrValue(in.CreatedBy),
		UpdatedAt:           ibmc.DateTimeToMetaV1Time(in.UpdatedAt),
		UpdatedBy:           reference.FromPtrValue(in.UpdatedBy),
		DeletedAt:           ibmc.DateTimeToMetaV1Time(in.DeletedAt),
		DeletedBy:           reference.FromPtrValue(in.DeletedBy),
		ScheduledReclaimAt:  ibmc.DateTimeToMetaV1Time(in.ScheduledReclaimAt),
		ScheduledReclaimBy:  reference.FromPtrValue(in.ScheduledReclaimBy),
		RestoredAt:          ibmc.DateTimeToMetaV1Time(in.RestoredAt),
		RestoredBy:          reference.FromPtrValue(in.RestoredBy),
	}
	// ServiceEndpoints can be found in instance.Parameters["service-endpoints"]
	return o, nil
}

// GenerateTarget generates Target from Crn
func GenerateTarget(in *rcv2.ResourceInstance) string {
	if in.Crn == nil {
		return ""
	}
	crn, err := crn.Parse(*in.Crn)
	if err != nil {
		return ""
	}
	return crn.Region
}

// GeneratePlanHistory generates []v1alpha1.PlanHistoryItem from []rcv2.PlanHistoryItem
func GeneratePlanHistory(in []rcv2.PlanHistoryItem) []v1alpha1.PlanHistoryItem {
	if in == nil {
		return nil
	}
	o := make([]v1alpha1.PlanHistoryItem, 0)
	for _, phi := range in {
		o = append(o, GeneratePlanHistoryItem(phi))
	}
	return o
}

// GeneratePlanHistoryItem generates v1alpha1.PlanHistoryItem from rcv2.PlanHistoryItem
func GeneratePlanHistoryItem(in rcv2.PlanHistoryItem) v1alpha1.PlanHistoryItem {
	planHistoryItem := v1alpha1.PlanHistoryItem{
		ResourcePlanID: reference.FromPtrValue(in.ResourcePlanID),
		StartDate:      ibmc.DateTimeToMetaV1Time(in.StartDate),
	}
	return planHistoryItem
}

// IsUpToDate checks whether current state is up-to-date compared to the given set of parameters.
func IsUpToDate(client ibmc.ClientSession, in *v1alpha1.ResourceInstanceParameters, observed *rcv2.ResourceInstance, l logging.Logger) (bool, error) {
	desired := in.DeepCopy()
	actual, err := GenerateResourceInstanceParameters(client, observed)
	if err != nil {
		return false, err
	}

	l.Info(cmp.Diff(desired, actual))

	return cmp.Equal(desired, actual, cmpopts.EquateEmpty(), cmpopts.IgnoreFields(v1alpha1.ResourceInstanceParameters{})), nil
}

// GenerateResourceInstanceParameters generates service instance parameters from resource instance
func GenerateResourceInstanceParameters(client ibmc.ClientSession, in *rcv2.ResourceInstance) (*v1alpha1.ResourceInstanceParameters, error) {

	rgName, err := ibmc.GetResourceGroupName(client, reference.FromPtrValue(in.ResourceGroupID))
	if err != nil {
		return nil, err
	}

	sName := ibmc.GetServiceName(in)
	pName, err := ibmc.GetResourcePlanName(client, sName, reference.FromPtrValue(in.ResourcePlanID))
	if err != nil {
		return nil, err
	}

	tags, err := ibmc.GetResourceInstanceTags(client, reference.FromPtrValue(in.Crn))
	if err != nil {
		return nil, err
	}

	o := &v1alpha1.ResourceInstanceParameters{
		Name:              reference.FromPtrValue(in.Name),
		Target:            GenerateTarget(in),
		ResourceGroupName: rgName,
		ServiceName:       sName,
		ResourcePlanName:  reference.FromPtrValue(pName),
		AllowCleanup:      in.AllowCleanup,
		Parameters:        ibmc.MapToRawExtension(in.Parameters),
		Tags:              tags,
		EntityLock:        reference.ToPtrValue(strconv.FormatBool(*in.Locked)),
	}
	return o, nil
}
