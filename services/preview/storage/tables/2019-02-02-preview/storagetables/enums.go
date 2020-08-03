package storagetables

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

// GeoReplicationStatusType enumerates the values for geo replication status type.
type GeoReplicationStatusType string

const (
	// Bootstrap ...
	Bootstrap GeoReplicationStatusType = "bootstrap"
	// Live ...
	Live GeoReplicationStatusType = "live"
	// Unavailable ...
	Unavailable GeoReplicationStatusType = "unavailable"
)

// PossibleGeoReplicationStatusTypeValues returns an array of possible values for the GeoReplicationStatusType const type.
func PossibleGeoReplicationStatusTypeValues() []GeoReplicationStatusType {
	return []GeoReplicationStatusType{Bootstrap, Live, Unavailable}
}

// OdataMetadataFormat enumerates the values for odata metadata format.
type OdataMetadataFormat string

const (
	// Applicationjsonodatafullmetadata ...
	Applicationjsonodatafullmetadata OdataMetadataFormat = "application/json;odata=fullmetadata"
	// Applicationjsonodataminimalmetadata ...
	Applicationjsonodataminimalmetadata OdataMetadataFormat = "application/json;odata=minimalmetadata"
	// Applicationjsonodatanometadata ...
	Applicationjsonodatanometadata OdataMetadataFormat = "application/json;odata=nometadata"
)

// PossibleOdataMetadataFormatValues returns an array of possible values for the OdataMetadataFormat const type.
func PossibleOdataMetadataFormatValues() []OdataMetadataFormat {
	return []OdataMetadataFormat{Applicationjsonodatafullmetadata, Applicationjsonodataminimalmetadata, Applicationjsonodatanometadata}
}

// ResponseFormat enumerates the values for response format.
type ResponseFormat string

const (
	// ReturnContent ...
	ReturnContent ResponseFormat = "return-content"
	// ReturnNoContent ...
	ReturnNoContent ResponseFormat = "return-no-content"
)

// PossibleResponseFormatValues returns an array of possible values for the ResponseFormat const type.
func PossibleResponseFormatValues() []ResponseFormat {
	return []ResponseFormat{ReturnContent, ReturnNoContent}
}
