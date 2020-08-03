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

// CompileMode enumerates the values for compile mode.
type CompileMode string

const (
	// Full ...
	Full CompileMode = "Full"
	// Semantic ...
	Semantic CompileMode = "Semantic"
	// SingleBox ...
	SingleBox CompileMode = "SingleBox"
)

// PossibleCompileModeValues returns an array of possible values for the CompileMode const type.
func PossibleCompileModeValues() []CompileMode {
	return []CompileMode{Full, Semantic, SingleBox}
}

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// JobManagerResource ...
	JobManagerResource ResourceType = "JobManagerResource"
	// JobManagerResourceInUserFolder ...
	JobManagerResourceInUserFolder ResourceType = "JobManagerResourceInUserFolder"
	// StatisticsResource ...
	StatisticsResource ResourceType = "StatisticsResource"
	// StatisticsResourceInUserFolder ...
	StatisticsResourceInUserFolder ResourceType = "StatisticsResourceInUserFolder"
	// VertexResource ...
	VertexResource ResourceType = "VertexResource"
	// VertexResourceInUserFolder ...
	VertexResourceInUserFolder ResourceType = "VertexResourceInUserFolder"
)

// PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{JobManagerResource, JobManagerResourceInUserFolder, StatisticsResource, StatisticsResourceInUserFolder, VertexResource, VertexResourceInUserFolder}
}

// Result enumerates the values for result.
type Result string

const (
	// Cancelled ...
	Cancelled Result = "Cancelled"
	// Failed ...
	Failed Result = "Failed"
	// None ...
	None Result = "None"
	// Succeeded ...
	Succeeded Result = "Succeeded"
)

// PossibleResultValues returns an array of possible values for the Result const type.
func PossibleResultValues() []Result {
	return []Result{Cancelled, Failed, None, Succeeded}
}

// SeverityTypes enumerates the values for severity types.
type SeverityTypes string

const (
	// Error ...
	Error SeverityTypes = "Error"
	// Info ...
	Info SeverityTypes = "Info"
	// Warning ...
	Warning SeverityTypes = "Warning"
)

// PossibleSeverityTypesValues returns an array of possible values for the SeverityTypes const type.
func PossibleSeverityTypesValues() []SeverityTypes {
	return []SeverityTypes{Error, Info, Warning}
}

// State enumerates the values for state.
type State string

const (
	// StateAccepted ...
	StateAccepted State = "Accepted"
	// StateCompiling ...
	StateCompiling State = "Compiling"
	// StateEnded ...
	StateEnded State = "Ended"
	// StateNew ...
	StateNew State = "New"
	// StatePaused ...
	StatePaused State = "Paused"
	// StateQueued ...
	StateQueued State = "Queued"
	// StateRunning ...
	StateRunning State = "Running"
	// StateScheduling ...
	StateScheduling State = "Scheduling"
	// StateStarting ...
	StateStarting State = "Starting"
	// StateWaitingForCapacity ...
	StateWaitingForCapacity State = "WaitingForCapacity"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateAccepted, StateCompiling, StateEnded, StateNew, StatePaused, StateQueued, StateRunning, StateScheduling, StateStarting, StateWaitingForCapacity}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeHive ...
	TypeHive Type = "Hive"
	// TypeJobProperties ...
	TypeJobProperties Type = "JobProperties"
	// TypeUSQL ...
	TypeUSQL Type = "USql"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeHive, TypeJobProperties, TypeUSQL}
}

// TypeEnum enumerates the values for type enum.
type TypeEnum string

const (
	// Hive ...
	Hive TypeEnum = "Hive"
	// USQL ...
	USQL TypeEnum = "USql"
)

// PossibleTypeEnumValues returns an array of possible values for the TypeEnum const type.
func PossibleTypeEnumValues() []TypeEnum {
	return []TypeEnum{Hive, USQL}
}
