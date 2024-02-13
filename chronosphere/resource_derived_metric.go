package chronosphere

import (
	"fmt"

	"go.uber.org/atomic"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DerivedMetricFromModel maps an API model to an intschema model.
func DerivedMetricFromModel(m *models.Configv1DerivedMetric) (*intschema.DerivedMetric, error) {
	return derivedMetricConverter{}.fromModel(m)
}

func resourceDerivedMetric() *schema.Resource {
	r := newGenericResource[
		*models.Configv1DerivedMetric,
		intschema.DerivedMetric,
		*intschema.DerivedMetric,
	](
		"derived_metric",
		derivedMetricConverter{},
		generatedDerivedMetric{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.DerivedMetric,
		CustomizeDiff: r.ValidateDryRun(&DerivedMetricDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// DerivedMetricRunCount tracks how many times dry run is run during validation for testing.
var DerivedMetricDryRunCount atomic.Int64

type derivedMetricConverter struct{}

func (derivedMetricConverter) toModel(
	m *intschema.DerivedMetric,
) (*models.Configv1DerivedMetric, error) {
	return &models.Configv1DerivedMetric{
		Name:        m.Name,
		Slug:        m.Slug,
		MetricName:  m.MetricName,
		Queries:     sliceutil.Map(m.Queries, derivedMetricSelectorQueryToModel),
		Description: m.Description,
	}, nil
}

func (derivedMetricConverter) fromModel(
	m *models.Configv1DerivedMetric,
) (*intschema.DerivedMetric, error) {
	queries, err := sliceutil.MapErr(m.Queries, derivedMetricSelectorQueryFromModel)
	if err != nil {
		return nil, err
	}
	return &intschema.DerivedMetric{
		Name:        m.Name,
		Slug:        m.Slug,
		MetricName:  m.MetricName,
		Queries:     queries,
		Description: m.Description,
	}, nil
}

func derivedMetricSelectorQueryToModel(
	q intschema.DerivedMetricQueries,
) *models.DerivedMetricSelectorQuery {
	return &models.DerivedMetricSelectorQuery{
		Query: &models.DerivedMetricQuery{
			PrometheusExpr: q.Query.Expr,
			Variables:      sliceutil.Map(q.Query.Variables, derivedMetricQueryVariableToModel),
		},
		Selector: derivedMetricSelectorToModel(q.Selector),
	}
}

func derivedMetricSelectorQueryFromModel(
	q *models.DerivedMetricSelectorQuery,
) (intschema.DerivedMetricQueries, error) {
	selector, err := derivedMetricSelectorFromModel(q.Selector)
	if err != nil {
		return intschema.DerivedMetricQueries{}, err
	}
	return intschema.DerivedMetricQueries{
		Query:    derivedMetricQueryFromModel(q.Query),
		Selector: selector,
	}, nil
}

func derivedMetricSelectorToModel(
	s *intschema.DerivedMetricQueriesSelector,
) *models.DerivedMetricSelector {
	if s == nil {
		return nil
	}
	return &models.DerivedMetricSelector{
		Labels: derivedMetricLabelsToModel(s.Labels),
	}
}

func derivedMetricSelectorFromModel(
	s *models.DerivedMetricSelector,
) (*intschema.DerivedMetricQueriesSelector, error) {
	if s == nil {
		return nil, nil
	}
	labels, err := derivedMetricLabelsFromModel(s.Labels)
	if err != nil {
		return nil, err
	}
	return &intschema.DerivedMetricQueriesSelector{
		Labels: labels,
	}, nil
}

func derivedMetricQueryFromModel(
	q *models.DerivedMetricQuery,
) intschema.DerivedMetricQueriesQuery {
	if q == nil {
		return intschema.DerivedMetricQueriesQuery{}
	}
	return intschema.DerivedMetricQueriesQuery{
		Expr:      q.PrometheusExpr,
		Variables: sliceutil.Map(q.Variables, derivedMetricQueryVariableFromModel),
	}
}

func derivedMetricQueryVariableToModel(
	v intschema.DerivedMetricQueriesQueryVariables,
) *models.DerivedMetricVariable {
	return &models.DerivedMetricVariable{
		DefaultPrometheusSelector: v.DefaultSelector,
		Name:                      v.Name,
	}
}

func derivedMetricQueryVariableFromModel(
	v *models.DerivedMetricVariable,
) intschema.DerivedMetricQueriesQueryVariables {
	return intschema.DerivedMetricQueriesQueryVariables{
		DefaultSelector: v.DefaultPrometheusSelector,
		Name:            v.Name,
	}
}

func derivedMetricLabelsToModel(
	labels map[string]string,
) []*models.Configv1DerivedMetricLabelMatcher {
	var result []*models.Configv1DerivedMetricLabelMatcher
	for _, k := range sortedKeys(labels) {
		result = append(result, &models.Configv1DerivedMetricLabelMatcher{
			Type:  models.Configv1DerivedMetricLabelMatcherMatcherTypeEXACT,
			Name:  k,
			Value: labels[k],
		})
	}
	return result
}

func derivedMetricLabelsFromModel(
	labels []*models.Configv1DerivedMetricLabelMatcher,
) (map[string]string, error) {
	m := make(map[string]string)
	for _, l := range labels {
		if l.Type != models.Configv1DerivedMetricLabelMatcherMatcherTypeEXACT {
			return nil, fmt.Errorf(
				"cannot parse server labels: type must be EXACT, got %q", l.Type)
		}
		if _, ok := m[l.Name]; ok {
			return nil, fmt.Errorf(
				"cannot parse server labels: duplicate name found: %q", l.Name)
		}
		m[l.Name] = l.Value
	}
	return m, nil
}
