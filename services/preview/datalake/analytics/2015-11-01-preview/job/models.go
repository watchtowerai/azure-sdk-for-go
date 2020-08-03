package job

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

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/datalake/analytics/2015-11-01-preview/job"

// DataPath a Data Lake Analytics U-SQL job data path item.
type DataPath struct {
	autorest.Response `json:"-"`
	// JobID - READ-ONLY; Gets the id of the job this data is for.
	JobID *uuid.UUID `json:"jobId,omitempty"`
	// Command - READ-ONLY; Gets the command that this job data relates to.
	Command *string `json:"command,omitempty"`
	// Paths - READ-ONLY; Gets the list of paths to all of the job data.
	Paths *[]string `json:"paths,omitempty"`
}

// ErrorDetails the Data Lake Analytics job error details.
type ErrorDetails struct {
	// Description - READ-ONLY; Gets the error message description
	Description *string `json:"description,omitempty"`
	// Details - READ-ONLY; Gets the details of the error message.
	Details *string `json:"details,omitempty"`
	// EndOffset - READ-ONLY; Gets the end offset in the job where the error was found.
	EndOffset *int32 `json:"endOffset,omitempty"`
	// ErrorID - READ-ONLY; Gets the specific identifier for the type of error encountered in the job.
	ErrorID *string `json:"errorId,omitempty"`
	// FilePath - READ-ONLY; Gets the path to any supplemental error files, if any.
	FilePath *string `json:"filePath,omitempty"`
	// HelpLink - READ-ONLY; Gets the link to MSDN or Azure help for this type of error, if any.
	HelpLink *string `json:"helpLink,omitempty"`
	// InternalDiagnostics - READ-ONLY; Gets the internal diagnostic stack trace if the user requesting the job error details has sufficient permissions it will be retrieved, otherwise it will be empty.
	InternalDiagnostics *string `json:"internalDiagnostics,omitempty"`
	// LineNumber - READ-ONLY; Gets the specific line number in the job where the error occurred.
	LineNumber *int32 `json:"lineNumber,omitempty"`
	// Message - READ-ONLY; Gets the user friendly error message for the failure.
	Message *string `json:"message,omitempty"`
	// Resolution - READ-ONLY; Gets the recommended resolution for the failure, if any.
	Resolution *string `json:"resolution,omitempty"`
	// InnerError - READ-ONLY; Gets the inner error of this specific job error message, if any.
	InnerError *InnerError `json:"InnerError,omitempty"`
	// Severity - READ-ONLY; Gets the severity level of the failure. Possible values include: 'Warning', 'Error'
	Severity SeverityTypes `json:"severity,omitempty"`
	// Source - READ-ONLY; Gets the ultimate source of the failure (usually either SYSTEM or USER).
	Source *string `json:"source,omitempty"`
	// StartOffset - READ-ONLY; Gets the start offset in the job where the error was found
	StartOffset *int32 `json:"startOffset,omitempty"`
}

// HiveJobProperties ...
type HiveJobProperties struct {
	// StatementInfo - Gets or sets the statement information for each statement in the script
	StatementInfo *[]HiveJobStatementInfo `json:"statementInfo,omitempty"`
	// LogsLocation - Gets or sets the Hive logs location
	LogsLocation *string `json:"logsLocation,omitempty"`
	// WarehouseLocation - Gets or sets the runtime version of the U-SQL engine to use
	WarehouseLocation *string `json:"warehouseLocation,omitempty"`
	// StatementCount - Gets or sets the number of statements that will be run based on the script
	StatementCount *int32 `json:"statementCount,omitempty"`
	// ExecutedStatementCount - Gets or sets the number of statements that have been run based on the script
	ExecutedStatementCount *int32 `json:"executedStatementCount,omitempty"`
	// RuntimeVersion - Gets or sets the runtime version of the U-SQL engine to use
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
	// Script - Gets or sets the U-SQL script to run
	Script *string `json:"script,omitempty"`
	// Type - Possible values include: 'TypeJobProperties', 'TypeUSQL', 'TypeHive'
	Type Type `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for HiveJobProperties.
func (hjp HiveJobProperties) MarshalJSON() ([]byte, error) {
	hjp.Type = TypeHive
	objectMap := make(map[string]interface{})
	if hjp.StatementInfo != nil {
		objectMap["statementInfo"] = hjp.StatementInfo
	}
	if hjp.LogsLocation != nil {
		objectMap["logsLocation"] = hjp.LogsLocation
	}
	if hjp.WarehouseLocation != nil {
		objectMap["warehouseLocation"] = hjp.WarehouseLocation
	}
	if hjp.StatementCount != nil {
		objectMap["statementCount"] = hjp.StatementCount
	}
	if hjp.ExecutedStatementCount != nil {
		objectMap["executedStatementCount"] = hjp.ExecutedStatementCount
	}
	if hjp.RuntimeVersion != nil {
		objectMap["runtimeVersion"] = hjp.RuntimeVersion
	}
	if hjp.Script != nil {
		objectMap["script"] = hjp.Script
	}
	if hjp.Type != "" {
		objectMap["type"] = hjp.Type
	}
	return json.Marshal(objectMap)
}

// AsUSQLJobProperties is the BasicProperties implementation for HiveJobProperties.
func (hjp HiveJobProperties) AsUSQLJobProperties() (*USQLJobProperties, bool) {
	return nil, false
}

// AsHiveJobProperties is the BasicProperties implementation for HiveJobProperties.
func (hjp HiveJobProperties) AsHiveJobProperties() (*HiveJobProperties, bool) {
	return &hjp, true
}

// AsProperties is the BasicProperties implementation for HiveJobProperties.
func (hjp HiveJobProperties) AsProperties() (*Properties, bool) {
	return nil, false
}

// AsBasicProperties is the BasicProperties implementation for HiveJobProperties.
func (hjp HiveJobProperties) AsBasicProperties() (BasicProperties, bool) {
	return &hjp, true
}

// HiveJobStatementInfo ...
type HiveJobStatementInfo struct {
	// LogLocation - Gets or sets the log location for this statement.
	LogLocation *string `json:"logLocation,omitempty"`
	// ResultPreviewLocation - Gets or sets the result preview location for this statement.
	ResultPreviewLocation *string `json:"resultPreviewLocation,omitempty"`
	// ResultLocation - Gets or sets the result location for this statement.
	ResultLocation *string `json:"resultLocation,omitempty"`
	// ErrorMessage - Gets or sets the error message for this statement.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// InfoListResult list of jobInfo items.
type InfoListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Gets the list of jobInfo items.
	Value *[]Information `json:"value,omitempty"`
	// NextLink - READ-ONLY; Gets the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Count - READ-ONLY; Gets the total count of results that are available, but might not be returned in the current page.
	Count *int64 `json:"count,omitempty"`
}

// InfoListResultIterator provides access to a complete listing of Information values.
type InfoListResultIterator struct {
	i    int
	page InfoListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *InfoListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InfoListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *InfoListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter InfoListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter InfoListResultIterator) Response() InfoListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter InfoListResultIterator) Value() Information {
	if !iter.page.NotDone() {
		return Information{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the InfoListResultIterator type.
func NewInfoListResultIterator(page InfoListResultPage) InfoListResultIterator {
	return InfoListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ilr InfoListResult) IsEmpty() bool {
	return ilr.Value == nil || len(*ilr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ilr InfoListResult) hasNextLink() bool {
	return ilr.NextLink != nil && len(*ilr.NextLink) != 0
}

// infoListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ilr InfoListResult) infoListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !ilr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ilr.NextLink)))
}

// InfoListResultPage contains a page of Information values.
type InfoListResultPage struct {
	fn  func(context.Context, InfoListResult) (InfoListResult, error)
	ilr InfoListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *InfoListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InfoListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ilr)
		if err != nil {
			return err
		}
		page.ilr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *InfoListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page InfoListResultPage) NotDone() bool {
	return !page.ilr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page InfoListResultPage) Response() InfoListResult {
	return page.ilr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page InfoListResultPage) Values() []Information {
	if page.ilr.IsEmpty() {
		return nil
	}
	return *page.ilr.Value
}

// Creates a new instance of the InfoListResultPage type.
func NewInfoListResultPage(getNextPage func(context.Context, InfoListResult) (InfoListResult, error)) InfoListResultPage {
	return InfoListResultPage{fn: getNextPage}
}

// Information the common Data Lake Analytics job information properties.
type Information struct {
	autorest.Response `json:"-"`
	// JobID - Gets or sets the job's unique identifier (a GUID).
	JobID *uuid.UUID `json:"jobId,omitempty"`
	// Name - Gets or sets the friendly name of the job.
	Name *string `json:"name,omitempty"`
	// Type - Gets or sets the job type of the current job (Hive or USql). Possible values include: 'USQL', 'Hive'
	Type TypeEnum `json:"type,omitempty"`
	// Submitter - Gets or sets the user or account that submitted the job.
	Submitter *string `json:"submitter,omitempty"`
	// ErrorMessage - READ-ONLY; Gets the error message details for the job, if the job failed.
	ErrorMessage *[]ErrorDetails `json:"errorMessage,omitempty"`
	// DegreeOfParallelism - Gets or sets the degree of parallelism used for this job. This must be greater than 0.
	DegreeOfParallelism *int32 `json:"degreeOfParallelism,omitempty"`
	// DegreeOfParallelismPercent - the degree of parallelism in percentage used for this job.
	DegreeOfParallelismPercent *float64 `json:"degreeOfParallelismPercent,omitempty"`
	// Priority - Gets or sets the priority value for the current job. Lower numbers have a higher priority. By default, a job has a priority of 1000. This must be greater than 0.
	Priority *int32 `json:"priority,omitempty"`
	// SubmitTime - READ-ONLY; Gets the time the job was submitted to the service.
	SubmitTime *date.Time `json:"submitTime,omitempty"`
	// StartTime - READ-ONLY; Gets the start time of the job.
	StartTime *date.Time `json:"startTime,omitempty"`
	// EndTime - READ-ONLY; Gets the completion time of the job.
	EndTime *date.Time `json:"endTime,omitempty"`
	// State - READ-ONLY; Gets the job state. When the job is in the Ended state, refer to Result and ErrorMessage for details. Possible values include: 'StateAccepted', 'StateCompiling', 'StateEnded', 'StateNew', 'StateQueued', 'StateRunning', 'StateScheduling', 'StateStarting', 'StatePaused', 'StateWaitingForCapacity'
	State State `json:"state,omitempty"`
	// Result - READ-ONLY; Gets the result of job execution or the current result of the running job. Possible values include: 'None', 'Succeeded', 'Cancelled', 'Failed'
	Result Result `json:"result,omitempty"`
	// StateAuditRecords - READ-ONLY; Gets the job state audit records, indicating when various operations have been performed on this job.
	StateAuditRecords *[]StateAuditRecord `json:"stateAuditRecords,omitempty"`
	// HierarchyQueueNode - READ-ONLY; the name of hierarchy queue node this job is assigned to, null if job has not been assigned yet or the account doesn't have hierarchy queue.
	HierarchyQueueNode *string `json:"hierarchyQueueNode,omitempty"`
	// Properties - Gets or sets the job specific properties.
	Properties BasicProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Information.
func (i Information) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.JobID != nil {
		objectMap["jobId"] = i.JobID
	}
	if i.Name != nil {
		objectMap["name"] = i.Name
	}
	if i.Type != "" {
		objectMap["type"] = i.Type
	}
	if i.Submitter != nil {
		objectMap["submitter"] = i.Submitter
	}
	if i.DegreeOfParallelism != nil {
		objectMap["degreeOfParallelism"] = i.DegreeOfParallelism
	}
	if i.DegreeOfParallelismPercent != nil {
		objectMap["degreeOfParallelismPercent"] = i.DegreeOfParallelismPercent
	}
	if i.Priority != nil {
		objectMap["priority"] = i.Priority
	}
	objectMap["properties"] = i.Properties
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Information struct.
func (i *Information) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "jobId":
			if v != nil {
				var jobID uuid.UUID
				err = json.Unmarshal(*v, &jobID)
				if err != nil {
					return err
				}
				i.JobID = &jobID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				i.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar TypeEnum
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				i.Type = typeVar
			}
		case "submitter":
			if v != nil {
				var submitter string
				err = json.Unmarshal(*v, &submitter)
				if err != nil {
					return err
				}
				i.Submitter = &submitter
			}
		case "errorMessage":
			if v != nil {
				var errorMessage []ErrorDetails
				err = json.Unmarshal(*v, &errorMessage)
				if err != nil {
					return err
				}
				i.ErrorMessage = &errorMessage
			}
		case "degreeOfParallelism":
			if v != nil {
				var degreeOfParallelism int32
				err = json.Unmarshal(*v, &degreeOfParallelism)
				if err != nil {
					return err
				}
				i.DegreeOfParallelism = &degreeOfParallelism
			}
		case "degreeOfParallelismPercent":
			if v != nil {
				var degreeOfParallelismPercent float64
				err = json.Unmarshal(*v, &degreeOfParallelismPercent)
				if err != nil {
					return err
				}
				i.DegreeOfParallelismPercent = &degreeOfParallelismPercent
			}
		case "priority":
			if v != nil {
				var priority int32
				err = json.Unmarshal(*v, &priority)
				if err != nil {
					return err
				}
				i.Priority = &priority
			}
		case "submitTime":
			if v != nil {
				var submitTime date.Time
				err = json.Unmarshal(*v, &submitTime)
				if err != nil {
					return err
				}
				i.SubmitTime = &submitTime
			}
		case "startTime":
			if v != nil {
				var startTime date.Time
				err = json.Unmarshal(*v, &startTime)
				if err != nil {
					return err
				}
				i.StartTime = &startTime
			}
		case "endTime":
			if v != nil {
				var endTime date.Time
				err = json.Unmarshal(*v, &endTime)
				if err != nil {
					return err
				}
				i.EndTime = &endTime
			}
		case "state":
			if v != nil {
				var state State
				err = json.Unmarshal(*v, &state)
				if err != nil {
					return err
				}
				i.State = state
			}
		case "result":
			if v != nil {
				var resultVar Result
				err = json.Unmarshal(*v, &resultVar)
				if err != nil {
					return err
				}
				i.Result = resultVar
			}
		case "stateAuditRecords":
			if v != nil {
				var stateAuditRecords []StateAuditRecord
				err = json.Unmarshal(*v, &stateAuditRecords)
				if err != nil {
					return err
				}
				i.StateAuditRecords = &stateAuditRecords
			}
		case "hierarchyQueueNode":
			if v != nil {
				var hierarchyQueueNode string
				err = json.Unmarshal(*v, &hierarchyQueueNode)
				if err != nil {
					return err
				}
				i.HierarchyQueueNode = &hierarchyQueueNode
			}
		case "properties":
			if v != nil {
				properties, err := unmarshalBasicProperties(*v)
				if err != nil {
					return err
				}
				i.Properties = properties
			}
		}
	}

	return nil
}

// InnerError the Data Lake Analytics job error details.
type InnerError struct {
	// DiagnosticCode - READ-ONLY; Gets the diagnostic error code.
	DiagnosticCode *int32 `json:"diagnosticCode,omitempty"`
	// Severity - READ-ONLY; Gets the severity level of the failure. Possible values include: 'Warning', 'Error'
	Severity SeverityTypes `json:"severity,omitempty"`
	// Details - READ-ONLY; Gets the details of the error message.
	Details *string `json:"details,omitempty"`
	// Component - READ-ONLY; Gets the component that failed.
	Component *string `json:"component,omitempty"`
	// ErrorID - READ-ONLY; Gets the specific identifier for the type of error encountered in the job.
	ErrorID *string `json:"errorId,omitempty"`
	// HelpLink - READ-ONLY; Gets the link to MSDN or Azure help for this type of error, if any.
	HelpLink *string `json:"helpLink,omitempty"`
	// InternalDiagnostics - READ-ONLY; Gets the internal diagnostic stack trace if the user requesting the job error details has sufficient permissions it will be retrieved, otherwise it will be empty.
	InternalDiagnostics *string `json:"internalDiagnostics,omitempty"`
	// Message - READ-ONLY; Gets the user friendly error message for the failure.
	Message *string `json:"message,omitempty"`
	// Resolution - READ-ONLY; Gets the recommended resolution for the failure, if any.
	Resolution *string `json:"resolution,omitempty"`
	// Source - READ-ONLY; Gets the ultimate source of the failure (usually either SYSTEM or USER).
	Source *string `json:"source,omitempty"`
	// Description - READ-ONLY; Gets the error message description
	Description *string `json:"description,omitempty"`
}

// BasicProperties the common Data Lake Analytics job properties.
type BasicProperties interface {
	AsUSQLJobProperties() (*USQLJobProperties, bool)
	AsHiveJobProperties() (*HiveJobProperties, bool)
	AsProperties() (*Properties, bool)
}

// Properties the common Data Lake Analytics job properties.
type Properties struct {
	// RuntimeVersion - Gets or sets the runtime version of the U-SQL engine to use
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
	// Script - Gets or sets the U-SQL script to run
	Script *string `json:"script,omitempty"`
	// Type - Possible values include: 'TypeJobProperties', 'TypeUSQL', 'TypeHive'
	Type Type `json:"type,omitempty"`
}

func unmarshalBasicProperties(body []byte) (BasicProperties, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["type"] {
	case string(TypeUSQL):
		var usjp USQLJobProperties
		err := json.Unmarshal(body, &usjp)
		return usjp, err
	case string(TypeHive):
		var hjp HiveJobProperties
		err := json.Unmarshal(body, &hjp)
		return hjp, err
	default:
		var p Properties
		err := json.Unmarshal(body, &p)
		return p, err
	}
}
func unmarshalBasicPropertiesArray(body []byte) ([]BasicProperties, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	pArray := make([]BasicProperties, len(rawMessages))

	for index, rawMessage := range rawMessages {
		p, err := unmarshalBasicProperties(*rawMessage)
		if err != nil {
			return nil, err
		}
		pArray[index] = p
	}
	return pArray, nil
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	p.Type = TypeJobProperties
	objectMap := make(map[string]interface{})
	if p.RuntimeVersion != nil {
		objectMap["runtimeVersion"] = p.RuntimeVersion
	}
	if p.Script != nil {
		objectMap["script"] = p.Script
	}
	if p.Type != "" {
		objectMap["type"] = p.Type
	}
	return json.Marshal(objectMap)
}

// AsUSQLJobProperties is the BasicProperties implementation for Properties.
func (p Properties) AsUSQLJobProperties() (*USQLJobProperties, bool) {
	return nil, false
}

// AsHiveJobProperties is the BasicProperties implementation for Properties.
func (p Properties) AsHiveJobProperties() (*HiveJobProperties, bool) {
	return nil, false
}

// AsProperties is the BasicProperties implementation for Properties.
func (p Properties) AsProperties() (*Properties, bool) {
	return &p, true
}

// AsBasicProperties is the BasicProperties implementation for Properties.
func (p Properties) AsBasicProperties() (BasicProperties, bool) {
	return &p, true
}

// Resource the Data Lake Analytics U-SQL job resources.
type Resource struct {
	// Name - Gets or set the name of the resource.
	Name *string `json:"name,omitempty"`
	// ResourcePath - Gets or sets the path to the resource.
	ResourcePath *string `json:"resourcePath,omitempty"`
	// Type - Gets or sets the job resource type. Possible values include: 'VertexResource', 'StatisticsResource'
	Type ResourceType `json:"type,omitempty"`
}

// StateAuditRecord the Data Lake Analytics U-SQL job state audit records for tracking the lifecycle of a job.
type StateAuditRecord struct {
	// NewState - READ-ONLY; Gets the new state the job is in.
	NewState *string `json:"newState,omitempty"`
	// TimeStamp - READ-ONLY; Gets the time stamp that the state change took place.
	TimeStamp *date.Time `json:"timeStamp,omitempty"`
	// RequestedByUser - READ-ONLY; Gets the user who requests the change.
	RequestedByUser *string `json:"requestedByUser,omitempty"`
	// Details - READ-ONLY; Gets  the details of the audit log.
	Details *string `json:"details,omitempty"`
}

// Statistics the Data Lake Analytics U-SQL job execution statistics.
type Statistics struct {
	autorest.Response `json:"-"`
	// LastUpdateTimeUtc - READ-ONLY; Gets the last update time for the statistics.
	LastUpdateTimeUtc *date.Time `json:"lastUpdateTimeUtc,omitempty"`
	// Stages - READ-ONLY; Gets the list of stages for the job.
	Stages *[]StatisticsVertexStage `json:"stages,omitempty"`
}

// StatisticsVertexStage the Data Lake Analytics U-SQL job statistics vertex stage information.
type StatisticsVertexStage struct {
	// DataRead - READ-ONLY; Gets the amount of data read, in bytes.
	DataRead *int64 `json:"dataRead,omitempty"`
	// DataReadCrossPod - READ-ONLY; Gets the amount of data read across multiple pods, in bytes.
	DataReadCrossPod *int64 `json:"dataReadCrossPod,omitempty"`
	// DataReadIntraPod - READ-ONLY; Gets the amount of data read in one pod, in bytes.
	DataReadIntraPod *int64 `json:"dataReadIntraPod,omitempty"`
	// DataToRead - READ-ONLY; Gets the amount of data remaining to be read, in bytes.
	DataToRead *int64 `json:"dataToRead,omitempty"`
	// DataWritten - READ-ONLY; Gets the amount of data written, in bytes.
	DataWritten *int64 `json:"dataWritten,omitempty"`
	// DuplicateDiscardCount - READ-ONLY; Gets the number of duplicates that were discarded.
	DuplicateDiscardCount *int32 `json:"duplicateDiscardCount,omitempty"`
	// FailedCount - READ-ONLY; Gets the number of failures that occurred in this stage.
	FailedCount *int32 `json:"failedCount,omitempty"`
	// MaxVertexDataRead - READ-ONLY; Gets the maximum amount of data read in a single vertex, in bytes.
	MaxVertexDataRead *int64 `json:"maxVertexDataRead,omitempty"`
	// MinVertexDataRead - READ-ONLY; Gets the minimum amount of data read in a single vertex, in bytes.
	MinVertexDataRead *int64 `json:"minVertexDataRead,omitempty"`
	// ReadFailureCount - READ-ONLY; Gets the number of read failures in this stage.
	ReadFailureCount *int32 `json:"readFailureCount,omitempty"`
	// RevocationCount - READ-ONLY; Gets the number of vertices that were revoked during this stage.
	RevocationCount *int32 `json:"revocationCount,omitempty"`
	// RunningCount - READ-ONLY; Gets the number of currently running vertices in this stage.
	RunningCount *int32 `json:"runningCount,omitempty"`
	// ScheduledCount - READ-ONLY; Gets the number of currently scheduled vertices in this stage
	ScheduledCount *int32 `json:"scheduledCount,omitempty"`
	// StageName - READ-ONLY; Gets the name of this stage in job execution.
	StageName *string `json:"stageName,omitempty"`
	// SucceededCount - READ-ONLY; Gets the number of vertices that succeeded in this stage.
	SucceededCount *int32 `json:"succeededCount,omitempty"`
	// TempDataWritten - READ-ONLY; Gets the amount of temporary data written, in bytes.
	TempDataWritten *int64 `json:"tempDataWritten,omitempty"`
	// TotalCount - READ-ONLY; Gets the total vertex count for this stage.
	TotalCount *int32 `json:"totalCount,omitempty"`
	// TotalFailedTime - READ-ONLY; Gets the amount of time that failed vertices took up in this stage.
	TotalFailedTime *string `json:"totalFailedTime,omitempty"`
	// TotalProgress - READ-ONLY; Gets the current progress of this stage, as a percentage.
	TotalProgress *int32 `json:"totalProgress,omitempty"`
	// TotalSucceededTime - READ-ONLY; Gets the amount of time all successful vertices took in this stage.
	TotalSucceededTime *string `json:"totalSucceededTime,omitempty"`
}

// USQLJobProperties ...
type USQLJobProperties struct {
	// Resources - Gets or sets the list of resources that are required by the job
	Resources *[]Resource `json:"resources,omitempty"`
	// Statistics - Gets or sets the job specific statistics.
	Statistics *Statistics `json:"statistics,omitempty"`
	// DebugData - Gets or sets the job specific debug data locations.
	DebugData *DataPath `json:"debugData,omitempty"`
	// AlgebraFilePath - READ-ONLY; Gets the U-SQL algebra file path after the job has completed
	AlgebraFilePath *string `json:"algebraFilePath,omitempty"`
	// TotalCompilationTime - READ-ONLY; Gets the total time this job spent compiling. This value should not be set by the user and will be ignored if it is.
	TotalCompilationTime *string `json:"totalCompilationTime,omitempty"`
	// TotalPauseTime - READ-ONLY; Gets the total time this job spent paused. This value should not be set by the user and will be ignored if it is.
	TotalPauseTime *string `json:"totalPauseTime,omitempty"`
	// TotalQueuedTime - READ-ONLY; Gets the total time this job spent queued. This value should not be set by the user and will be ignored if it is.
	TotalQueuedTime *string `json:"totalQueuedTime,omitempty"`
	// TotalRunningTime - READ-ONLY; Gets the total time this job spent executing. This value should not be set by the user and will be ignored if it is.
	TotalRunningTime *string `json:"totalRunningTime,omitempty"`
	// RootProcessNodeID - READ-ONLY; Gets the ID used to identify the job manager coordinating job execution. This value should not be set by the user and will be ignored if it is.
	RootProcessNodeID *string `json:"rootProcessNodeId,omitempty"`
	// YarnApplicationID - READ-ONLY; Gets the ID used to identify the yarn application executing the job. This value should not be set by the user and will be ignored if it is.
	YarnApplicationID *string `json:"yarnApplicationId,omitempty"`
	// YarnApplicationTimeStamp - READ-ONLY; Gets the timestamp (in ticks) for the yarn application executing the job. This value should not be set by the user and will be ignored if it is.
	YarnApplicationTimeStamp *int64 `json:"yarnApplicationTimeStamp,omitempty"`
	// CompileMode - Gets or sets the compile mode for the job. Possible values include: 'Semantic', 'Full', 'SingleBox'
	CompileMode CompileMode `json:"compileMode,omitempty"`
	// RuntimeVersion - Gets or sets the runtime version of the U-SQL engine to use
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
	// Script - Gets or sets the U-SQL script to run
	Script *string `json:"script,omitempty"`
	// Type - Possible values include: 'TypeJobProperties', 'TypeUSQL', 'TypeHive'
	Type Type `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for USQLJobProperties.
func (usjp USQLJobProperties) MarshalJSON() ([]byte, error) {
	usjp.Type = TypeUSQL
	objectMap := make(map[string]interface{})
	if usjp.Resources != nil {
		objectMap["resources"] = usjp.Resources
	}
	if usjp.Statistics != nil {
		objectMap["statistics"] = usjp.Statistics
	}
	if usjp.DebugData != nil {
		objectMap["debugData"] = usjp.DebugData
	}
	if usjp.CompileMode != "" {
		objectMap["compileMode"] = usjp.CompileMode
	}
	if usjp.RuntimeVersion != nil {
		objectMap["runtimeVersion"] = usjp.RuntimeVersion
	}
	if usjp.Script != nil {
		objectMap["script"] = usjp.Script
	}
	if usjp.Type != "" {
		objectMap["type"] = usjp.Type
	}
	return json.Marshal(objectMap)
}

// AsUSQLJobProperties is the BasicProperties implementation for USQLJobProperties.
func (usjp USQLJobProperties) AsUSQLJobProperties() (*USQLJobProperties, bool) {
	return &usjp, true
}

// AsHiveJobProperties is the BasicProperties implementation for USQLJobProperties.
func (usjp USQLJobProperties) AsHiveJobProperties() (*HiveJobProperties, bool) {
	return nil, false
}

// AsProperties is the BasicProperties implementation for USQLJobProperties.
func (usjp USQLJobProperties) AsProperties() (*Properties, bool) {
	return nil, false
}

// AsBasicProperties is the BasicProperties implementation for USQLJobProperties.
func (usjp USQLJobProperties) AsBasicProperties() (BasicProperties, bool) {
	return &usjp, true
}
