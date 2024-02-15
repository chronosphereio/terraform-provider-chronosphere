// Copyright 2023 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chronosphere

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"go.uber.org/atomic"
)

func TraceJaegerRemoteSamplingStrategyFromModel(
	m *models.Configv1TraceJaegerRemoteSamplingStrategy,
) (*intschema.TraceJaegerRemoteSamplingStrategy, error) {
	return (traceJRSConverter{}).fromModel(m)
}

func resourceTraceJRSStrategy() *schema.Resource {
	r := newGenericResource[
		*models.Configv1TraceJaegerRemoteSamplingStrategy,
		intschema.TraceJaegerRemoteSamplingStrategy,
		*intschema.TraceJaegerRemoteSamplingStrategy,
	](
		"trace_jaeger_remote_sampling_strategy",
		traceJRSConverter{},
		generatedTraceJaegerRemoteSamplingStrategy{},
	)

	return &schema.Resource{
		CreateContext: r.CreateContext,
		ReadContext:   r.ReadContext,
		UpdateContext: r.UpdateContext,
		DeleteContext: r.DeleteContext,
		Schema:        tfschema.TraceJaegerRemoteSamplingStrategy,
		CustomizeDiff: r.ValidateDryRun(&TraceJaegerRemoteSamplingStrategyDryRunCount),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// TraceJaegerRemoteSamplingStrategyDryRunCount tracks how many times dry run is run during validation for testing.
var TraceJaegerRemoteSamplingStrategyDryRunCount atomic.Int64

type traceJRSConverter struct{}

func (c traceJRSConverter) toModel(
	s *intschema.TraceJaegerRemoteSamplingStrategy,
) (*models.Configv1TraceJaegerRemoteSamplingStrategy, error) {
	if s == nil {
		return nil, nil
	}
	a, err := toModelAppliedStrategy(s.AppliedStrategy)
	if err != nil {
		return nil, err
	}
	m := &models.Configv1TraceJaegerRemoteSamplingStrategy{
		Slug:            s.Slug,
		Name:            s.Name,
		ServiceName:     s.ServiceName,
		AppliedStrategy: a,
	}
	return m, nil
}

func toModelAppliedStrategy(
	s intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategy,
) (*models.TraceJaegerRemoteSamplingStrategyAppliedStrategy, error) {
	m := &models.TraceJaegerRemoteSamplingStrategyAppliedStrategy{}
	if s.ProbabilisticStrategy != nil {
		m.ProbabilisticStrategy = toModelProbabilisticStrategy(s.ProbabilisticStrategy)
	}
	if s.RateLimitingStrategy != nil {
		m.RateLimitingStrategy = toModelRateLimitingStrategy(s.RateLimitingStrategy)
	}
	if s.PerOperationStrategies != nil {
		o, err := toModelPerOperationStrategies(s.PerOperationStrategies)
		if err != nil {
			return nil, err
		}
		m.PerOperationStrategies = o
	}
	return m, nil
}

func toModelPerOperationStrategies(
	s *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategies,
) (*models.TraceJaegerRemoteSamplingStrategyPerOperationSamplingStrategies, error) {
	m := &models.TraceJaegerRemoteSamplingStrategyPerOperationSamplingStrategies{
		DefaultSamplingRate:              s.DefaultSamplingRate,
		DefaultLowerBoundTracesPerSecond: s.DefaultLowerBoundTracesPerSecond,
		DefaultUpperBoundTracesPerSecond: s.DefaultUpperBoundTracesPerSecond,
		PerOperationStrategies: make(
			[]*models.PerOperationSamplingStrategiesPerOperationSamplingStrategy,
			0,
			len(s.PerOperationStrategies),
		),
	}
	for _, perOp := range s.PerOperationStrategies {
		mPerOp := &models.PerOperationSamplingStrategiesPerOperationSamplingStrategy{
			Operation:                     perOp.Operation,
			ProbabilisticSamplingStrategy: toModelPerOperationProbabilisticStrategy(&perOp.ProbabilisticStrategy),
		}
		m.PerOperationStrategies = append(m.PerOperationStrategies, mPerOp)
	}
	return m, nil
}

func toModelPerOperationProbabilisticStrategy(
	s *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategiesPerOperationStrategiesProbabilisticStrategy,
) *models.TraceJaegerRemoteSamplingStrategyProbabilisticStrategy {
	if s == nil {
		return nil
	}
	return &models.TraceJaegerRemoteSamplingStrategyProbabilisticStrategy{
		SamplingRate: s.SamplingRate,
	}
}

func toModelRateLimitingStrategy(
	s *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyRateLimitingStrategy,
) *models.TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy {
	if s == nil {
		return nil
	}
	return &models.TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy{
		MaxTracesPerSecond: int32(s.MaxTracesPerSecond),
	}
}

func toModelProbabilisticStrategy(
	s *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyProbabilisticStrategy,
) *models.TraceJaegerRemoteSamplingStrategyProbabilisticStrategy {
	if s == nil {
		return nil
	}
	return &models.TraceJaegerRemoteSamplingStrategyProbabilisticStrategy{
		SamplingRate: s.SamplingRate,
	}
}

func (traceJRSConverter) fromModel(
	m *models.Configv1TraceJaegerRemoteSamplingStrategy,
) (*intschema.TraceJaegerRemoteSamplingStrategy, error) {
	if m == nil {
		return nil, nil
	}
	a, err := fromModelAppliedStrategy(m.AppliedStrategy)
	if err != nil {
		return nil, err
	}
	s := &intschema.TraceJaegerRemoteSamplingStrategy{
		Slug:            m.Slug,
		Name:            m.Name,
		ServiceName:     m.ServiceName,
		AppliedStrategy: *a,
	}

	return s, nil
}

func fromModelAppliedStrategy(
	m *models.TraceJaegerRemoteSamplingStrategyAppliedStrategy,
) (*intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategy, error) {
	if m == nil {
		return nil, nil
	}
	s := &intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategy{}
	if m.ProbabilisticStrategy != nil {
		s.ProbabilisticStrategy = fromModelProbabilisticStrategy(m.ProbabilisticStrategy)
	}
	if m.RateLimitingStrategy != nil {
		s.RateLimitingStrategy = fromModelRateLimitingStrategy(m.RateLimitingStrategy)
	}
	if m.PerOperationStrategies != nil {
		o, err := fromModelPerOperationStrategies(m.PerOperationStrategies)
		if err != nil {
			return nil, err
		}
		s.PerOperationStrategies = o
	}
	return s, nil
}

func fromModelRateLimitingStrategy(
	m *models.TraceJaegerRemoteSamplingStrategyRateLimitingSamplingStrategy,
) *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyRateLimitingStrategy {
	if m == nil {
		return nil
	}
	return &intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyRateLimitingStrategy{
		MaxTracesPerSecond: int64(m.MaxTracesPerSecond),
	}
}

func fromModelProbabilisticStrategy(
	m *models.TraceJaegerRemoteSamplingStrategyProbabilisticStrategy,
) *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyProbabilisticStrategy {
	if m == nil {
		return nil
	}
	return &intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyProbabilisticStrategy{
		SamplingRate: m.SamplingRate,
	}
}

func fromModelPerOpProbabilisticStrategy(
	m *models.TraceJaegerRemoteSamplingStrategyProbabilisticStrategy,
) *intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategiesPerOperationStrategiesProbabilisticStrategy {
	if m == nil {
		return nil
	}
	return &intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategiesPerOperationStrategiesProbabilisticStrategy{
		SamplingRate: m.SamplingRate,
	}
}

func fromModelPerOperationStrategies(
	m *models.TraceJaegerRemoteSamplingStrategyPerOperationSamplingStrategies,
) (*intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategies, error) {
	s := &intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategies{
		DefaultSamplingRate:              m.DefaultSamplingRate,
		DefaultLowerBoundTracesPerSecond: m.DefaultLowerBoundTracesPerSecond,
		DefaultUpperBoundTracesPerSecond: m.DefaultUpperBoundTracesPerSecond,
		PerOperationStrategies: make(
			[]intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategiesPerOperationStrategies,
			0,
			len(m.PerOperationStrategies),
		),
	}
	for _, mPerOp := range m.PerOperationStrategies {
		perOp := intschema.TraceJaegerRemoteSamplingStrategyAppliedStrategyPerOperationStrategiesPerOperationStrategies{
			Operation: mPerOp.Operation,
		}
		if perOpProb := fromModelPerOpProbabilisticStrategy(mPerOp.ProbabilisticSamplingStrategy); perOpProb != nil {
			perOp.ProbabilisticStrategy = *perOpProb
		}
		s.PerOperationStrategies = append(s.PerOperationStrategies, perOp)
	}
	return s, nil
}
