package alertsmanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AlertModificationEvent enumerates the values for alert modification event.
type AlertModificationEvent string

const (
	// AlertCreated ...
	AlertCreated AlertModificationEvent = "AlertCreated"
	// MonitorConditionChange ...
	MonitorConditionChange AlertModificationEvent = "MonitorConditionChange"
	// StateChange ...
	StateChange AlertModificationEvent = "StateChange"
)

// PossibleAlertModificationEventValues returns an array of possible values for the AlertModificationEvent const type.
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return []AlertModificationEvent{AlertCreated, MonitorConditionChange, StateChange}
}

// AlertsSortByFields enumerates the values for alerts sort by fields.
type AlertsSortByFields string

const (
	// AlertsSortByFieldsAlertState ...
	AlertsSortByFieldsAlertState AlertsSortByFields = "alertState"
	// AlertsSortByFieldsLastModifiedDateTime ...
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = "lastModifiedDateTime"
	// AlertsSortByFieldsMonitorCondition ...
	AlertsSortByFieldsMonitorCondition AlertsSortByFields = "monitorCondition"
	// AlertsSortByFieldsName ...
	AlertsSortByFieldsName AlertsSortByFields = "name"
	// AlertsSortByFieldsSeverity ...
	AlertsSortByFieldsSeverity AlertsSortByFields = "severity"
	// AlertsSortByFieldsStartDateTime ...
	AlertsSortByFieldsStartDateTime AlertsSortByFields = "startDateTime"
	// AlertsSortByFieldsTargetResource ...
	AlertsSortByFieldsTargetResource AlertsSortByFields = "targetResource"
	// AlertsSortByFieldsTargetResourceGroup ...
	AlertsSortByFieldsTargetResourceGroup AlertsSortByFields = "targetResourceGroup"
	// AlertsSortByFieldsTargetResourceName ...
	AlertsSortByFieldsTargetResourceName AlertsSortByFields = "targetResourceName"
	// AlertsSortByFieldsTargetResourceType ...
	AlertsSortByFieldsTargetResourceType AlertsSortByFields = "targetResourceType"
)

// PossibleAlertsSortByFieldsValues returns an array of possible values for the AlertsSortByFields const type.
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return []AlertsSortByFields{AlertsSortByFieldsAlertState, AlertsSortByFieldsLastModifiedDateTime, AlertsSortByFieldsMonitorCondition, AlertsSortByFieldsName, AlertsSortByFieldsSeverity, AlertsSortByFieldsStartDateTime, AlertsSortByFieldsTargetResource, AlertsSortByFieldsTargetResourceGroup, AlertsSortByFieldsTargetResourceName, AlertsSortByFieldsTargetResourceType}
}

// AlertsSummaryGroupByFields enumerates the values for alerts summary group by fields.
type AlertsSummaryGroupByFields string

const (
	// AlertsSummaryGroupByFieldsAlertRule ...
	AlertsSummaryGroupByFieldsAlertRule AlertsSummaryGroupByFields = "alertRule"
	// AlertsSummaryGroupByFieldsAlertState ...
	AlertsSummaryGroupByFieldsAlertState AlertsSummaryGroupByFields = "alertState"
	// AlertsSummaryGroupByFieldsMonitorCondition ...
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = "monitorCondition"
	// AlertsSummaryGroupByFieldsMonitorService ...
	AlertsSummaryGroupByFieldsMonitorService AlertsSummaryGroupByFields = "monitorService"
	// AlertsSummaryGroupByFieldsSeverity ...
	AlertsSummaryGroupByFieldsSeverity AlertsSummaryGroupByFields = "severity"
	// AlertsSummaryGroupByFieldsSignalType ...
	AlertsSummaryGroupByFieldsSignalType AlertsSummaryGroupByFields = "signalType"
)

// PossibleAlertsSummaryGroupByFieldsValues returns an array of possible values for the AlertsSummaryGroupByFields const type.
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return []AlertsSummaryGroupByFields{AlertsSummaryGroupByFieldsAlertRule, AlertsSummaryGroupByFieldsAlertState, AlertsSummaryGroupByFieldsMonitorCondition, AlertsSummaryGroupByFieldsMonitorService, AlertsSummaryGroupByFieldsSeverity, AlertsSummaryGroupByFieldsSignalType}
}

// AlertState enumerates the values for alert state.
type AlertState string

const (
	// AlertStateAcknowledged ...
	AlertStateAcknowledged AlertState = "Acknowledged"
	// AlertStateClosed ...
	AlertStateClosed AlertState = "Closed"
	// AlertStateNew ...
	AlertStateNew AlertState = "New"
)

// PossibleAlertStateValues returns an array of possible values for the AlertState const type.
func PossibleAlertStateValues() []AlertState {
	return []AlertState{AlertStateAcknowledged, AlertStateClosed, AlertStateNew}
}

// MonitorCondition enumerates the values for monitor condition.
type MonitorCondition string

const (
	// Fired ...
	Fired MonitorCondition = "Fired"
	// Resolved ...
	Resolved MonitorCondition = "Resolved"
)

// PossibleMonitorConditionValues returns an array of possible values for the MonitorCondition const type.
func PossibleMonitorConditionValues() []MonitorCondition {
	return []MonitorCondition{Fired, Resolved}
}

// MonitorService enumerates the values for monitor service.
type MonitorService string

const (
	// ActivityLogAdministrative ...
	ActivityLogAdministrative MonitorService = "ActivityLog Administrative"
	// ActivityLogAutoscale ...
	ActivityLogAutoscale MonitorService = "ActivityLog Autoscale"
	// ActivityLogPolicy ...
	ActivityLogPolicy MonitorService = "ActivityLog Policy"
	// ActivityLogRecommendation ...
	ActivityLogRecommendation MonitorService = "ActivityLog Recommendation"
	// ActivityLogSecurity ...
	ActivityLogSecurity MonitorService = "ActivityLog Security"
	// ApplicationInsights ...
	ApplicationInsights MonitorService = "Application Insights"
	// LogAnalytics ...
	LogAnalytics MonitorService = "Log Analytics"
	// Nagios ...
	Nagios MonitorService = "Nagios"
	// Platform ...
	Platform MonitorService = "Platform"
	// SCOM ...
	SCOM MonitorService = "SCOM"
	// ServiceHealth ...
	ServiceHealth MonitorService = "ServiceHealth"
	// SmartDetector ...
	SmartDetector MonitorService = "SmartDetector"
	// VMInsights ...
	VMInsights MonitorService = "VM Insights"
	// Zabbix ...
	Zabbix MonitorService = "Zabbix"
)

// PossibleMonitorServiceValues returns an array of possible values for the MonitorService const type.
func PossibleMonitorServiceValues() []MonitorService {
	return []MonitorService{ActivityLogAdministrative, ActivityLogAutoscale, ActivityLogPolicy, ActivityLogRecommendation, ActivityLogSecurity, ApplicationInsights, LogAnalytics, Nagios, Platform, SCOM, ServiceHealth, SmartDetector, VMInsights, Zabbix}
}

// Severity enumerates the values for severity.
type Severity string

const (
	// Sev0 ...
	Sev0 Severity = "Sev0"
	// Sev1 ...
	Sev1 Severity = "Sev1"
	// Sev2 ...
	Sev2 Severity = "Sev2"
	// Sev3 ...
	Sev3 Severity = "Sev3"
	// Sev4 ...
	Sev4 Severity = "Sev4"
)

// PossibleSeverityValues returns an array of possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{Sev0, Sev1, Sev2, Sev3, Sev4}
}

// SignalType enumerates the values for signal type.
type SignalType string

const (
	// Log ...
	Log SignalType = "Log"
	// Metric ...
	Metric SignalType = "Metric"
	// Unknown ...
	Unknown SignalType = "Unknown"
)

// PossibleSignalTypeValues returns an array of possible values for the SignalType const type.
func PossibleSignalTypeValues() []SignalType {
	return []SignalType{Log, Metric, Unknown}
}

// SmartGroupModificationEvent enumerates the values for smart group modification event.
type SmartGroupModificationEvent string

const (
	// SmartGroupModificationEventAlertAdded ...
	SmartGroupModificationEventAlertAdded SmartGroupModificationEvent = "AlertAdded"
	// SmartGroupModificationEventAlertRemoved ...
	SmartGroupModificationEventAlertRemoved SmartGroupModificationEvent = "AlertRemoved"
	// SmartGroupModificationEventSmartGroupCreated ...
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = "SmartGroupCreated"
	// SmartGroupModificationEventStateChange ...
	SmartGroupModificationEventStateChange SmartGroupModificationEvent = "StateChange"
)

// PossibleSmartGroupModificationEventValues returns an array of possible values for the SmartGroupModificationEvent const type.
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return []SmartGroupModificationEvent{SmartGroupModificationEventAlertAdded, SmartGroupModificationEventAlertRemoved, SmartGroupModificationEventSmartGroupCreated, SmartGroupModificationEventStateChange}
}

// SmartGroupsSortByFields enumerates the values for smart groups sort by fields.
type SmartGroupsSortByFields string

const (
	// SmartGroupsSortByFieldsAlertsCount ...
	SmartGroupsSortByFieldsAlertsCount SmartGroupsSortByFields = "alertsCount"
	// SmartGroupsSortByFieldsLastModifiedDateTime ...
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = "lastModifiedDateTime"
	// SmartGroupsSortByFieldsSeverity ...
	SmartGroupsSortByFieldsSeverity SmartGroupsSortByFields = "severity"
	// SmartGroupsSortByFieldsStartDateTime ...
	SmartGroupsSortByFieldsStartDateTime SmartGroupsSortByFields = "startDateTime"
	// SmartGroupsSortByFieldsState ...
	SmartGroupsSortByFieldsState SmartGroupsSortByFields = "state"
)

// PossibleSmartGroupsSortByFieldsValues returns an array of possible values for the SmartGroupsSortByFields const type.
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return []SmartGroupsSortByFields{SmartGroupsSortByFieldsAlertsCount, SmartGroupsSortByFieldsLastModifiedDateTime, SmartGroupsSortByFieldsSeverity, SmartGroupsSortByFieldsStartDateTime, SmartGroupsSortByFieldsState}
}

// State enumerates the values for state.
type State string

const (
	// StateAcknowledged ...
	StateAcknowledged State = "Acknowledged"
	// StateClosed ...
	StateClosed State = "Closed"
	// StateNew ...
	StateNew State = "New"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateAcknowledged, StateClosed, StateNew}
}

// TimeRange enumerates the values for time range.
type TimeRange string

const (
	// Oned ...
	Oned TimeRange = "1d"
	// Oneh ...
	Oneh TimeRange = "1h"
	// Sevend ...
	Sevend TimeRange = "7d"
	// ThreeZerod ...
	ThreeZerod TimeRange = "30d"
)

// PossibleTimeRangeValues returns an array of possible values for the TimeRange const type.
func PossibleTimeRangeValues() []TimeRange {
	return []TimeRange{Oned, Oneh, Sevend, ThreeZerod}
}
