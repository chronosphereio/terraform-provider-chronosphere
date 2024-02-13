package chronosphere

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

// RecordingRuleFromModel maps an API model to an intschema model.
func RecordingRuleFromModel(m *models.Configv1RecordingRule) (*intschema.RecordingRule, error) {
	return recordingRuleConverter{}.fromModel(m)
}

func resourceRecordingRule() *schema.Resource {
	r := newGenericResource[
		*models.Configv1RecordingRule,
		intschema.RecordingRule,
		*intschema.RecordingRule,
	](
		"recording_rule",
		recordingRuleConverter{},
		generatedRecordingRule{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.RecordingRule,
		CustomizeDiff: r.ValidateDryRun(&RecordingRuleDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// RecordingRuleDryRunCount tracks how many times dry run is run during validation for testing.
var RecordingRuleDryRunCount atomic.Int64

type recordingRuleConverter struct{}

func (recordingRuleConverter) toModel(
	r *intschema.RecordingRule,
) (*models.Configv1RecordingRule, error) {
	intervalSecs, err := durationToSecs(r.Interval)
	if err != nil {
		return nil, err
	}
	return &models.Configv1RecordingRule{
		Name:           r.Name,
		Slug:           r.Slug,
		BucketSlug:     r.BucketId.Slug(),
		ExecutionGroup: r.ExecutionGroup.Slug(),
		PrometheusExpr: r.Expr,
		IntervalSecs:   intervalSecs,
		LabelPolicy: &models.Configv1RecordingRuleLabelPolicy{
			Add: r.Labels,
		},
		MetricName: r.MetricName,
	}, nil
}

func (recordingRuleConverter) fromModel(
	m *models.Configv1RecordingRule,
) (*intschema.RecordingRule, error) {
	r := &intschema.RecordingRule{
		Name:           m.Name,
		Slug:           m.Slug,
		BucketId:       tfid.Slug(m.BucketSlug),
		ExecutionGroup: tfid.Slug(m.ExecutionGroup),
		Expr:           m.PrometheusExpr,
		Interval:       durationFromSecs(m.IntervalSecs),
		MetricName:     m.MetricName,
	}
	if m.LabelPolicy != nil {
		r.Labels = m.LabelPolicy.Add
	}
	return r, nil
}
