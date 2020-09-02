/*
Copyright © 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"
	"regexp"
)

const (
	// MainEndpoint defines suffix of the root endpoint
	MainEndpoint = ""

	// GroupsEndpoint defines suffix of the groups request endpoint
	GroupsEndpoint = "groups"
	// DeleteOrganizationsEndpoint deletes all {organizations}(comma separated array). DEBUG only
	DeleteOrganizationsEndpoint = "organizations/{organizations}"
	// DeleteClustersEndpoint deletes all {clusters}(comma separated array). DEBUG only
	DeleteClustersEndpoint = "clusters/{clusters}"
	// OrganizationsEndpoint returns all organizations
	OrganizationsEndpoint = "organizations"
	// ClustersEndpoint returns reports for selected clusters
	ClustersEndpoint = "clusters"
	// ReportEndpoint returns report for provided {organization} and {cluster}
	ReportEndpoint = "report/{organization}/{cluster}"
	// ReportForClusterEndpoint returns report for provided {cluster} (w/o organization)
	ReportForClusterEndpoint = "report/{cluster}"
	// LikeRuleEndpoint likes rule with {rule_id} for {cluster} using current user(from auth header)
	LikeRuleEndpoint = "clusters/{cluster}/rules/{rule_id}/like"
	// DislikeRuleEndpoint dislikes rule with {rule_id} for {cluster} using current user(from auth header)
	DislikeRuleEndpoint = "clusters/{cluster}/rules/{rule_id}/dislike"
	// ResetVoteOnRuleEndpoint resets vote on rule with {rule_id} for {cluster} using current user(from auth header)
	ResetVoteOnRuleEndpoint = "clusters/{cluster}/rules/{rule_id}/reset_vote"
	// GetVoteOnRuleEndpoint is an endpoint to get vote on rule. DEBUG only
	GetVoteOnRuleEndpoint = "clusters/{cluster}/rules/{rule_id}/get_vote"
	// RuleEndpoint is an endpoint to create&delete a rule. DEBUG only
	RuleEndpoint = "rules/{rule_id}"
	// RuleErrorKeyEndpoint is for endpoints to create&delete a rule_error_key (DEBUG only)
	// and for endpoint to get a rule
	RuleErrorKeyEndpoint = "rules/{rule_id}/error_keys/{error_key}"
	// RuleGroupsEndpoint is a simple redirect endpoint to the insights-content-service API specified in configuration
	RuleGroupsEndpoint = "groups"
	// ClustersForOrganizationEndpoint returns all clusters for {organization}
	ClustersForOrganizationEndpoint = "organizations/{organization}/clusters"
	// DisableRuleForClusterEndpoint disables a rule for specified cluster
	DisableRuleForClusterEndpoint = "clusters/{cluster}/rules/{rule_id}/disable"
	// EnableRuleForClusterEndpoint re-enables a rule for specified cluster
	EnableRuleForClusterEndpoint = "clusters/{cluster}/rules/{rule_id}/enable"
	// MetricsEndpoint returns prometheus metrics
	MetricsEndpoint = "metrics"
)

// MakeURLToEndpoint creates URL to endpoint, use constants from file endpoints.go
func MakeURLToEndpoint(apiPrefix, endpoint string, args ...interface{}) string {
	re := regexp.MustCompile(`\{[a-zA-Z_0-9]+\}`)
	endpoint = re.ReplaceAllString(endpoint, "%v")
	return apiPrefix + fmt.Sprintf(endpoint, args...)
}
