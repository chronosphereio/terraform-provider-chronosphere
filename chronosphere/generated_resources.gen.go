// Code generated by go generate; DO NOT EDIT.
package chronosphere

import (
	"context"

	"fmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/apiclients"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/bucket"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/classic_dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/collection"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/derived_label"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/derived_metric"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/drop_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/grafana_dashboard"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/mapping_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/monitor"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/notification_policy"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/notifier"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/recording_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/rollup_rule"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/service_account"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/team"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_jaeger_remote_sampling_strategy"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/trace_metrics_rule"
	configv1models "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

type generatedBucket struct{}

func (generatedBucket) slugOf(m *configv1models.Configv1Bucket) string {
	return m.Slug
}

func (generatedBucket) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Bucket,
	dryRun bool,
) (string, error) {
	req := &bucket.CreateBucketParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateBucketRequest{
			Bucket: m,
			DryRun: dryRun,
		},
	}
	resp, err := clients.ConfigV1.Bucket.CreateBucket(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.Bucket
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedBucket) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1Bucket, error) {
	req := &bucket.ReadBucketParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.Bucket.ReadBucket(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Bucket, nil
}

func (generatedBucket) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Bucket,
	params updateParams,
) error {
	req := &bucket.UpdateBucketParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: bucket.UpdateBucketBody{
			Bucket:          m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.Bucket.UpdateBucket(req)
	return err
}
func (generatedBucket) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &bucket.DeleteBucketParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.Bucket.DeleteBucket(req)
	return err
}

type generatedCollection struct{}

func (generatedCollection) slugOf(m *configv1models.Configv1Collection) string {
	return m.Slug
}

func (generatedCollection) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Collection,
	dryRun bool,
) (string, error) {
	req := &collection.CreateCollectionParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateCollectionRequest{
			Collection: m,
			DryRun:     dryRun,
		},
	}
	resp, err := clients.ConfigV1.Collection.CreateCollection(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.Collection
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedCollection) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1Collection, error) {
	req := &collection.ReadCollectionParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.Collection.ReadCollection(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Collection, nil
}

func (generatedCollection) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Collection,
	params updateParams,
) error {
	req := &collection.UpdateCollectionParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: collection.UpdateCollectionBody{
			Collection:      m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.Collection.UpdateCollection(req)
	return err
}
func (generatedCollection) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &collection.DeleteCollectionParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.Collection.DeleteCollection(req)
	return err
}

type generatedDashboard struct{}

func (generatedDashboard) slugOf(m *configv1models.Configv1Dashboard) string {
	return m.Slug
}

func (generatedDashboard) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Dashboard,
	dryRun bool,
) (string, error) {
	req := &dashboard.CreateDashboardParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateDashboardRequest{
			Dashboard: m,
			DryRun:    dryRun,
		},
	}
	resp, err := clients.ConfigV1.Dashboard.CreateDashboard(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.Dashboard
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedDashboard) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1Dashboard, error) {
	req := &dashboard.ReadDashboardParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.Dashboard.ReadDashboard(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Dashboard, nil
}

func (generatedDashboard) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Dashboard,
	params updateParams,
) error {
	req := &dashboard.UpdateDashboardParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: dashboard.UpdateDashboardBody{
			Dashboard:       m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.Dashboard.UpdateDashboard(req)
	return err
}
func (generatedDashboard) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &dashboard.DeleteDashboardParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.Dashboard.DeleteDashboard(req)
	return err
}

type generatedDerivedLabel struct{}

func (generatedDerivedLabel) slugOf(m *configv1models.Configv1DerivedLabel) string {
	return m.Slug
}

func (generatedDerivedLabel) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1DerivedLabel,
	dryRun bool,
) (string, error) {
	req := &derived_label.CreateDerivedLabelParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateDerivedLabelRequest{
			DerivedLabel: m,
			DryRun:       dryRun,
		},
	}
	resp, err := clients.ConfigV1.DerivedLabel.CreateDerivedLabel(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.DerivedLabel
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedDerivedLabel) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1DerivedLabel, error) {
	req := &derived_label.ReadDerivedLabelParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.DerivedLabel.ReadDerivedLabel(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.DerivedLabel, nil
}

func (generatedDerivedLabel) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1DerivedLabel,
	params updateParams,
) error {
	req := &derived_label.UpdateDerivedLabelParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: derived_label.UpdateDerivedLabelBody{
			DerivedLabel:    m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.DerivedLabel.UpdateDerivedLabel(req)
	return err
}
func (generatedDerivedLabel) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &derived_label.DeleteDerivedLabelParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.DerivedLabel.DeleteDerivedLabel(req)
	return err
}

type generatedDerivedMetric struct{}

func (generatedDerivedMetric) slugOf(m *configv1models.Configv1DerivedMetric) string {
	return m.Slug
}

func (generatedDerivedMetric) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1DerivedMetric,
	dryRun bool,
) (string, error) {
	req := &derived_metric.CreateDerivedMetricParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateDerivedMetricRequest{
			DerivedMetric: m,
			DryRun:        dryRun,
		},
	}
	resp, err := clients.ConfigV1.DerivedMetric.CreateDerivedMetric(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.DerivedMetric
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedDerivedMetric) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1DerivedMetric, error) {
	req := &derived_metric.ReadDerivedMetricParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.DerivedMetric.ReadDerivedMetric(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.DerivedMetric, nil
}

func (generatedDerivedMetric) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1DerivedMetric,
	params updateParams,
) error {
	req := &derived_metric.UpdateDerivedMetricParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: derived_metric.UpdateDerivedMetricBody{
			DerivedMetric:   m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.DerivedMetric.UpdateDerivedMetric(req)
	return err
}
func (generatedDerivedMetric) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &derived_metric.DeleteDerivedMetricParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.DerivedMetric.DeleteDerivedMetric(req)
	return err
}

type generatedDropRule struct{}

func (generatedDropRule) slugOf(m *configv1models.Configv1DropRule) string {
	return m.Slug
}

func (generatedDropRule) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1DropRule,
	dryRun bool,
) (string, error) {
	req := &drop_rule.CreateDropRuleParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateDropRuleRequest{
			DropRule: m,
			DryRun:   dryRun,
		},
	}
	resp, err := clients.ConfigV1.DropRule.CreateDropRule(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.DropRule
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedDropRule) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1DropRule, error) {
	req := &drop_rule.ReadDropRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.DropRule.ReadDropRule(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.DropRule, nil
}

func (generatedDropRule) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1DropRule,
	params updateParams,
) error {
	req := &drop_rule.UpdateDropRuleParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: drop_rule.UpdateDropRuleBody{
			DropRule:        m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.DropRule.UpdateDropRule(req)
	return err
}
func (generatedDropRule) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &drop_rule.DeleteDropRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.DropRule.DeleteDropRule(req)
	return err
}

type generatedGrafanaDashboard struct{}

func (generatedGrafanaDashboard) slugOf(m *configv1models.Configv1GrafanaDashboard) string {
	return m.Slug
}

func (generatedGrafanaDashboard) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1GrafanaDashboard,
	dryRun bool,
) (string, error) {
	req := &grafana_dashboard.CreateGrafanaDashboardParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateGrafanaDashboardRequest{
			GrafanaDashboard: m,
			DryRun:           dryRun,
		},
	}
	resp, err := clients.ConfigV1.GrafanaDashboard.CreateGrafanaDashboard(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.GrafanaDashboard
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedGrafanaDashboard) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1GrafanaDashboard, error) {
	req := &grafana_dashboard.ReadGrafanaDashboardParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.GrafanaDashboard.ReadGrafanaDashboard(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.GrafanaDashboard, nil
}

func (generatedGrafanaDashboard) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1GrafanaDashboard,
	params updateParams,
) error {
	req := &grafana_dashboard.UpdateGrafanaDashboardParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: grafana_dashboard.UpdateGrafanaDashboardBody{
			GrafanaDashboard: m,
			CreateIfMissing:  params.createIfMissing,
			DryRun:           params.dryRun,
		},
	}
	_, err := clients.ConfigV1.GrafanaDashboard.UpdateGrafanaDashboard(req)
	return err
}
func (generatedGrafanaDashboard) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &grafana_dashboard.DeleteGrafanaDashboardParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.GrafanaDashboard.DeleteGrafanaDashboard(req)
	return err
}

type generatedMappingRule struct{}

func (generatedMappingRule) slugOf(m *configv1models.Configv1MappingRule) string {
	return m.Slug
}

func (generatedMappingRule) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1MappingRule,
	dryRun bool,
) (string, error) {
	req := &mapping_rule.CreateMappingRuleParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateMappingRuleRequest{
			MappingRule: m,
			DryRun:      dryRun,
		},
	}
	resp, err := clients.ConfigV1.MappingRule.CreateMappingRule(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.MappingRule
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedMappingRule) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1MappingRule, error) {
	req := &mapping_rule.ReadMappingRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.MappingRule.ReadMappingRule(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.MappingRule, nil
}

func (generatedMappingRule) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1MappingRule,
	params updateParams,
) error {
	req := &mapping_rule.UpdateMappingRuleParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: mapping_rule.UpdateMappingRuleBody{
			MappingRule:     m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.MappingRule.UpdateMappingRule(req)
	return err
}
func (generatedMappingRule) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &mapping_rule.DeleteMappingRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.MappingRule.DeleteMappingRule(req)
	return err
}

type generatedMonitor struct{}

func (generatedMonitor) slugOf(m *configv1models.Configv1Monitor) string {
	return m.Slug
}

func (generatedMonitor) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Monitor,
	dryRun bool,
) (string, error) {
	req := &monitor.CreateMonitorParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateMonitorRequest{
			Monitor: m,
			DryRun:  dryRun,
		},
	}
	resp, err := clients.ConfigV1.Monitor.CreateMonitor(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.Monitor
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedMonitor) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1Monitor, error) {
	req := &monitor.ReadMonitorParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.Monitor.ReadMonitor(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Monitor, nil
}

func (generatedMonitor) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Monitor,
	params updateParams,
) error {
	req := &monitor.UpdateMonitorParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: monitor.UpdateMonitorBody{
			Monitor:         m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.Monitor.UpdateMonitor(req)
	return err
}
func (generatedMonitor) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &monitor.DeleteMonitorParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.Monitor.DeleteMonitor(req)
	return err
}

type generatedNotificationPolicy struct{}

func (generatedNotificationPolicy) slugOf(m *configv1models.Configv1NotificationPolicy) string {
	return m.Slug
}

func (generatedNotificationPolicy) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1NotificationPolicy,
	dryRun bool,
) (string, error) {
	req := &notification_policy.CreateNotificationPolicyParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateNotificationPolicyRequest{
			NotificationPolicy: m,
			DryRun:             dryRun,
		},
	}
	resp, err := clients.ConfigV1.NotificationPolicy.CreateNotificationPolicy(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.NotificationPolicy
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedNotificationPolicy) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1NotificationPolicy, error) {
	req := &notification_policy.ReadNotificationPolicyParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.NotificationPolicy.ReadNotificationPolicy(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.NotificationPolicy, nil
}

func (generatedNotificationPolicy) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1NotificationPolicy,
	params updateParams,
) error {
	req := &notification_policy.UpdateNotificationPolicyParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: notification_policy.UpdateNotificationPolicyBody{
			NotificationPolicy: m,
			CreateIfMissing:    params.createIfMissing,
			DryRun:             params.dryRun,
		},
	}
	_, err := clients.ConfigV1.NotificationPolicy.UpdateNotificationPolicy(req)
	return err
}
func (generatedNotificationPolicy) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &notification_policy.DeleteNotificationPolicyParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.NotificationPolicy.DeleteNotificationPolicy(req)
	return err
}

type generatedNotifier struct{}

func (generatedNotifier) slugOf(m *configv1models.Configv1Notifier) string {
	return m.Slug
}

func (generatedNotifier) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Notifier,
	dryRun bool,
) (string, error) {
	req := &notifier.CreateNotifierParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateNotifierRequest{
			Notifier: m,
			DryRun:   dryRun,
		},
	}
	resp, err := clients.ConfigV1.Notifier.CreateNotifier(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.Notifier
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedNotifier) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1Notifier, error) {
	req := &notifier.ReadNotifierParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.Notifier.ReadNotifier(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Notifier, nil
}

func (generatedNotifier) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Notifier,
	params updateParams,
) error {
	req := &notifier.UpdateNotifierParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: notifier.UpdateNotifierBody{
			Notifier:        m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.Notifier.UpdateNotifier(req)
	return err
}
func (generatedNotifier) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &notifier.DeleteNotifierParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.Notifier.DeleteNotifier(req)
	return err
}

type generatedRecordingRule struct{}

func (generatedRecordingRule) slugOf(m *configv1models.Configv1RecordingRule) string {
	return m.Slug
}

func (generatedRecordingRule) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1RecordingRule,
	dryRun bool,
) (string, error) {
	req := &recording_rule.CreateRecordingRuleParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateRecordingRuleRequest{
			RecordingRule: m,
			DryRun:        dryRun,
		},
	}
	resp, err := clients.ConfigV1.RecordingRule.CreateRecordingRule(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.RecordingRule
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedRecordingRule) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1RecordingRule, error) {
	req := &recording_rule.ReadRecordingRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.RecordingRule.ReadRecordingRule(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.RecordingRule, nil
}

func (generatedRecordingRule) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1RecordingRule,
	params updateParams,
) error {
	req := &recording_rule.UpdateRecordingRuleParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: recording_rule.UpdateRecordingRuleBody{
			RecordingRule:   m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.RecordingRule.UpdateRecordingRule(req)
	return err
}
func (generatedRecordingRule) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &recording_rule.DeleteRecordingRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.RecordingRule.DeleteRecordingRule(req)
	return err
}

type generatedRollupRule struct{}

func (generatedRollupRule) slugOf(m *configv1models.Configv1RollupRule) string {
	return m.Slug
}

func (generatedRollupRule) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1RollupRule,
	dryRun bool,
) (string, error) {
	req := &rollup_rule.CreateRollupRuleParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateRollupRuleRequest{
			RollupRule: m,
			DryRun:     dryRun,
		},
	}
	resp, err := clients.ConfigV1.RollupRule.CreateRollupRule(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.RollupRule
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedRollupRule) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1RollupRule, error) {
	req := &rollup_rule.ReadRollupRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.RollupRule.ReadRollupRule(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.RollupRule, nil
}

func (generatedRollupRule) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1RollupRule,
	params updateParams,
) error {
	req := &rollup_rule.UpdateRollupRuleParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: rollup_rule.UpdateRollupRuleBody{
			RollupRule:      m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.RollupRule.UpdateRollupRule(req)
	return err
}
func (generatedRollupRule) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &rollup_rule.DeleteRollupRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.RollupRule.DeleteRollupRule(req)
	return err
}

type generatedServiceAccount struct{}

func (generatedServiceAccount) slugOf(m *configv1models.Configv1ServiceAccount) string {
	return m.Slug
}

func (generatedServiceAccount) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1ServiceAccount,
	dryRun bool,
) (string, error) {
	req := &service_account.CreateServiceAccountParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateServiceAccountRequest{
			ServiceAccount: m,
			DryRun:         dryRun,
		},
	}
	resp, err := clients.ConfigV1.ServiceAccount.CreateServiceAccount(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.ServiceAccount
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedServiceAccount) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1ServiceAccount, error) {
	req := &service_account.ReadServiceAccountParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.ServiceAccount.ReadServiceAccount(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.ServiceAccount, nil
}

func (generatedServiceAccount) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &service_account.DeleteServiceAccountParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.ServiceAccount.DeleteServiceAccount(req)
	return err
}

type generatedTeam struct{}

func (generatedTeam) slugOf(m *configv1models.Configv1Team) string {
	return m.Slug
}

func (generatedTeam) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Team,
	dryRun bool,
) (string, error) {
	req := &team.CreateTeamParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateTeamRequest{
			Team:   m,
			DryRun: dryRun,
		},
	}
	resp, err := clients.ConfigV1.Team.CreateTeam(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.Team
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedTeam) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1Team, error) {
	req := &team.ReadTeamParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.Team.ReadTeam(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Team, nil
}

func (generatedTeam) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1Team,
	params updateParams,
) error {
	req := &team.UpdateTeamParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: team.UpdateTeamBody{
			Team:            m,
			CreateIfMissing: params.createIfMissing,
			DryRun:          params.dryRun,
		},
	}
	_, err := clients.ConfigV1.Team.UpdateTeam(req)
	return err
}
func (generatedTeam) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &team.DeleteTeamParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.Team.DeleteTeam(req)
	return err
}

type generatedTraceJaegerRemoteSamplingStrategy struct{}

func (generatedTraceJaegerRemoteSamplingStrategy) slugOf(m *configv1models.Configv1TraceJaegerRemoteSamplingStrategy) string {
	return m.Slug
}

func (generatedTraceJaegerRemoteSamplingStrategy) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1TraceJaegerRemoteSamplingStrategy,
	dryRun bool,
) (string, error) {
	req := &trace_jaeger_remote_sampling_strategy.CreateTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateTraceJaegerRemoteSamplingStrategyRequest{
			TraceJaegerRemoteSamplingStrategy: m,
			DryRun:                            dryRun,
		},
	}
	resp, err := clients.ConfigV1.TraceJaegerRemoteSamplingStrategy.CreateTraceJaegerRemoteSamplingStrategy(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.TraceJaegerRemoteSamplingStrategy
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedTraceJaegerRemoteSamplingStrategy) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1TraceJaegerRemoteSamplingStrategy, error) {
	req := &trace_jaeger_remote_sampling_strategy.ReadTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.TraceJaegerRemoteSamplingStrategy.ReadTraceJaegerRemoteSamplingStrategy(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.TraceJaegerRemoteSamplingStrategy, nil
}

func (generatedTraceJaegerRemoteSamplingStrategy) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1TraceJaegerRemoteSamplingStrategy,
	params updateParams,
) error {
	req := &trace_jaeger_remote_sampling_strategy.UpdateTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: trace_jaeger_remote_sampling_strategy.UpdateTraceJaegerRemoteSamplingStrategyBody{
			TraceJaegerRemoteSamplingStrategy: m,
			CreateIfMissing:                   params.createIfMissing,
			DryRun:                            params.dryRun,
		},
	}
	_, err := clients.ConfigV1.TraceJaegerRemoteSamplingStrategy.UpdateTraceJaegerRemoteSamplingStrategy(req)
	return err
}
func (generatedTraceJaegerRemoteSamplingStrategy) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &trace_jaeger_remote_sampling_strategy.DeleteTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.TraceJaegerRemoteSamplingStrategy.DeleteTraceJaegerRemoteSamplingStrategy(req)
	return err
}

type generatedTraceMetricsRule struct{}

func (generatedTraceMetricsRule) slugOf(m *configv1models.Configv1TraceMetricsRule) string {
	return m.Slug
}

func (generatedTraceMetricsRule) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1TraceMetricsRule,
	dryRun bool,
) (string, error) {
	if dryRun {
		return "", fmt.Errorf("dry run not supported for this entity type")
	}
	req := &trace_metrics_rule.CreateTraceMetricsRuleParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateTraceMetricsRuleRequest{
			TraceMetricsRule: m,
		},
	}
	resp, err := clients.ConfigV1.TraceMetricsRule.CreateTraceMetricsRule(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.TraceMetricsRule
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedTraceMetricsRule) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1TraceMetricsRule, error) {
	req := &trace_metrics_rule.ReadTraceMetricsRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.TraceMetricsRule.ReadTraceMetricsRule(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.TraceMetricsRule, nil
}

func (generatedTraceMetricsRule) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1TraceMetricsRule,
	params updateParams,
) error {
	if params.dryRun {
		return fmt.Errorf("dry run not supported for this entity type")
	}
	req := &trace_metrics_rule.UpdateTraceMetricsRuleParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: trace_metrics_rule.UpdateTraceMetricsRuleBody{
			TraceMetricsRule: m,
			CreateIfMissing:  params.createIfMissing,
		},
	}
	_, err := clients.ConfigV1.TraceMetricsRule.UpdateTraceMetricsRule(req)
	return err
}
func (generatedTraceMetricsRule) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &trace_metrics_rule.DeleteTraceMetricsRuleParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.TraceMetricsRule.DeleteTraceMetricsRule(req)
	return err
}

type generatedClassicDashboard struct{}

func (generatedClassicDashboard) slugOf(m *configv1models.Configv1GrafanaDashboard) string {
	return m.Slug
}

func (generatedClassicDashboard) create(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1GrafanaDashboard,
	dryRun bool,
) (string, error) {
	req := &classic_dashboard.CreateClassicDashboardParams{
		Context: ctx,
		Body: &configv1models.Configv1CreateClassicDashboardRequest{
			ClassicDashboard: m,
			DryRun:           dryRun,
		},
	}
	resp, err := clients.ConfigV1.ClassicDashboard.CreateClassicDashboard(req)
	if err != nil {
		return "", err
	}
	e := resp.Payload.ClassicDashboard
	if e == nil {
		return "", nil
	}
	return e.Slug, nil
}

func (generatedClassicDashboard) read(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) (*configv1models.Configv1GrafanaDashboard, error) {
	req := &classic_dashboard.ReadClassicDashboardParams{
		Context: ctx,
		Slug:    slug,
	}
	resp, err := clients.ConfigV1.ClassicDashboard.ReadClassicDashboard(req)
	if err != nil {
		return nil, err
	}
	return resp.Payload.ClassicDashboard, nil
}

func (generatedClassicDashboard) update(
	ctx context.Context,
	clients apiclients.Clients,
	m *configv1models.Configv1GrafanaDashboard,
	params updateParams,
) error {
	req := &classic_dashboard.UpdateClassicDashboardParams{
		Context: ctx,
		Slug:    m.Slug,
		Body: classic_dashboard.UpdateClassicDashboardBody{
			ClassicDashboard: m,
			CreateIfMissing:  params.createIfMissing,
			DryRun:           params.dryRun,
		},
	}
	_, err := clients.ConfigV1.ClassicDashboard.UpdateClassicDashboard(req)
	return err
}
func (generatedClassicDashboard) delete(
	ctx context.Context,
	clients apiclients.Clients,
	slug string,
) error {
	req := &classic_dashboard.DeleteClassicDashboardParams{
		Context: ctx,
		Slug:    slug,
	}
	_, err := clients.ConfigV1.ClassicDashboard.DeleteClassicDashboard(req)
	return err
}
