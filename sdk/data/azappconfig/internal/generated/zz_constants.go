//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

type LabelFields string

const (
	LabelFieldsName LabelFields = "name"
)

// PossibleLabelFieldsValues returns the possible values for the LabelFields const type.
func PossibleLabelFieldsValues() []LabelFields {
	return []LabelFields{
		LabelFieldsName,
	}
}

type SettingFields string

const (
	SettingFieldsContentType  SettingFields = "content_type"
	SettingFieldsEtag         SettingFields = "etag"
	SettingFieldsKey          SettingFields = "key"
	SettingFieldsLabel        SettingFields = "label"
	SettingFieldsLastModified SettingFields = "last_modified"
	SettingFieldsLocked       SettingFields = "locked"
	SettingFieldsTags         SettingFields = "tags"
	SettingFieldsValue        SettingFields = "value"
)

// PossibleSettingFieldsValues returns the possible values for the SettingFields const type.
func PossibleSettingFieldsValues() []SettingFields {
	return []SettingFields{
		SettingFieldsContentType,
		SettingFieldsEtag,
		SettingFieldsKey,
		SettingFieldsLabel,
		SettingFieldsLastModified,
		SettingFieldsLocked,
		SettingFieldsTags,
		SettingFieldsValue,
	}
}
