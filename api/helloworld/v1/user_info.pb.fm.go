// Code generated by protoc-gen-fieldmask. DO NOT EDIT.
// source: api/helloworld/v1/user_info.proto

package v1

import (
	pbfieldmask "github.com/yeqown/protoc-gen-fieldmask/proto/fieldmask"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (x *GetUserInfoRequest) FieldMaskWithMode(mode pbfieldmask.MaskMode) *GetUserInfoRequest_FieldMask {
	fm := &GetUserInfoRequest_FieldMask{
		maskMode: mode,
		mask:     pbfieldmask.New(x.FieldMask),
	}

	return fm
}

// FieldMask_Prune generates *GetUserInfoRequest_FieldMask with filter mode, so that
// only the fields in the GetUserInfoRequest.FieldMask will be
// appended into the GetUserInfoRequest.
func (x *GetUserInfoRequest) FieldMask_Filter() *GetUserInfoRequest_FieldMask {
	return x.FieldMaskWithMode(pbfieldmask.MaskMode_Filter)
}

// FieldMask_Prune generates *GetUserInfoRequest_FieldMask with prune mode, so that
// only the fields NOT in the GetUserInfoRequest.FieldMask will be
// appended into the GetUserInfoRequest.
func (x *GetUserInfoRequest) FieldMask_Prune() *GetUserInfoRequest_FieldMask {
	return x.FieldMaskWithMode(pbfieldmask.MaskMode_Prune)
}

// GetUserInfoRequest_FieldMask provide provide helper functions to deal with FieldMask.
type GetUserInfoRequest_FieldMask struct {
	maskMode pbfieldmask.MaskMode
	mask     pbfieldmask.NestedFieldMask
}

// _fm_GetUserInfoRequest is built in variable for GetUserInfoRequest to call FieldMask.Append
var _fm_GetUserInfoRequest = new(GetUserInfoRequest)

// MaskIn_Id append GetUserInfoRequest.Id into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskIn_Id() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoRequest, "id")

	return x
}

// MaskedIn_Id indicates the field GetUserInfoRequest.Id
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedIn_Id() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("id")
}

// _fm_GetUserInfoResponse is built in variable for GetUserInfoResponse to call FieldMask.Append
var _fm_GetUserInfoResponse = new(GetUserInfoResponse)

// MaskOut_ProfileInfo append GetUserInfoResponse.ProfileInfo into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_ProfileInfo() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "profile_info")

	return x
}

// MaskedOut_ProfileInfo indicates the field GetUserInfoRequest.ProfileInfo
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_ProfileInfo() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info")
}

// MaskOut_ProfileInfo_UserId append GetUserInfoResponse.UserId into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_ProfileInfo_UserId() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "profile_info.user_id")

	return x
}

// MaskedOut_ProfileInfo_UserId indicates the field GetUserInfoRequest.UserId
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_ProfileInfo_UserId() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info.user_id")
}

// MaskOut_ProfileInfo_Name append GetUserInfoResponse.Name into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_ProfileInfo_Name() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "profile_info.name")

	return x
}

// MaskedOut_ProfileInfo_Name indicates the field GetUserInfoRequest.Name
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_ProfileInfo_Name() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info.name")
}

// MaskOut_ProfileInfo_Email append GetUserInfoResponse.Email into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_ProfileInfo_Email() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "profile_info.email")

	return x
}

// MaskedOut_ProfileInfo_Email indicates the field GetUserInfoRequest.Email
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_ProfileInfo_Email() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info.email")
}

// MaskOut_VipInfo append GetUserInfoResponse.VipInfo into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_VipInfo() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "vip_info")

	return x
}

// MaskedOut_VipInfo indicates the field GetUserInfoRequest.VipInfo
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_VipInfo() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("vip_info")
}

// MaskOut_VipInfo_Level append GetUserInfoResponse.Level into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_VipInfo_Level() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "vip_info.level")

	return x
}

// MaskedOut_VipInfo_Level indicates the field GetUserInfoRequest.Level
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_VipInfo_Level() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("vip_info.level")
}

// MaskOut_VipInfo_Exp append GetUserInfoResponse.Exp into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_VipInfo_Exp() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "vip_info.exp")

	return x
}

// MaskedOut_VipInfo_Exp indicates the field GetUserInfoRequest.Exp
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_VipInfo_Exp() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("vip_info.exp")
}

// MaskOut_SocialInfo append GetUserInfoResponse.SocialInfo into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_SocialInfo() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "social_info")

	return x
}

// MaskedOut_SocialInfo indicates the field GetUserInfoRequest.SocialInfo
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_SocialInfo() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("social_info")
}

// MaskOut_SocialInfo_Following append GetUserInfoResponse.Following into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_SocialInfo_Following() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "social_info.following")

	return x
}

// MaskedOut_SocialInfo_Following indicates the field GetUserInfoRequest.Following
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_SocialInfo_Following() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("social_info.following")
}

// MaskOut_SocialInfo_Follower append GetUserInfoResponse.Follower into
// GetUserInfoRequest.FieldMask.
func (x *GetUserInfoRequest) MaskOut_SocialInfo_Follower() *GetUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_GetUserInfoResponse, "social_info.follower")

	return x
}

// MaskedOut_SocialInfo_Follower indicates the field GetUserInfoRequest.Follower
// exists in the GetUserInfoRequest.FieldMask or not.
func (x *GetUserInfoRequest_FieldMask) MaskedOut_SocialInfo_Follower() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("social_info.follower")
}

// Mask only affects the fields in the GetUserInfoRequest.
func (x *GetUserInfoRequest_FieldMask) Mask(m *GetUserInfoResponse) *GetUserInfoResponse {
	switch x.maskMode {
	case pbfieldmask.MaskMode_Filter:
		x.mask.Filter(m)
	case pbfieldmask.MaskMode_Prune:
		x.mask.Prune(m)
	}

	return m
}

func (x *UpdateUserInfoRequest) FieldMaskWithMode(mode pbfieldmask.MaskMode) *UpdateUserInfoRequest_FieldMask {
	fm := &UpdateUserInfoRequest_FieldMask{
		maskMode: mode,
		mask:     pbfieldmask.New(x.FieldMask),
	}

	return fm
}

// FieldMask_Prune generates *UpdateUserInfoRequest_FieldMask with filter mode, so that
// only the fields in the UpdateUserInfoRequest.FieldMask will be
// appended into the UpdateUserInfoRequest.
func (x *UpdateUserInfoRequest) FieldMask_Filter() *UpdateUserInfoRequest_FieldMask {
	return x.FieldMaskWithMode(pbfieldmask.MaskMode_Filter)
}

// FieldMask_Prune generates *UpdateUserInfoRequest_FieldMask with prune mode, so that
// only the fields NOT in the UpdateUserInfoRequest.FieldMask will be
// appended into the UpdateUserInfoRequest.
func (x *UpdateUserInfoRequest) FieldMask_Prune() *UpdateUserInfoRequest_FieldMask {
	return x.FieldMaskWithMode(pbfieldmask.MaskMode_Prune)
}

// UpdateUserInfoRequest_FieldMask provide provide helper functions to deal with FieldMask.
type UpdateUserInfoRequest_FieldMask struct {
	maskMode pbfieldmask.MaskMode
	mask     pbfieldmask.NestedFieldMask
}

// _fm_UpdateUserInfoRequest is built in variable for UpdateUserInfoRequest to call FieldMask.Append
var _fm_UpdateUserInfoRequest = new(UpdateUserInfoRequest)

// MaskIn_Id append UpdateUserInfoRequest.Id into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskIn_Id() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoRequest, "id")

	return x
}

// MaskedIn_Id indicates the field UpdateUserInfoRequest.Id
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedIn_Id() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("id")
}

// MaskIn_Name append UpdateUserInfoRequest.Name into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskIn_Name() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoRequest, "name")

	return x
}

// MaskedIn_Name indicates the field UpdateUserInfoRequest.Name
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedIn_Name() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("name")
}

// MaskIn_Email append UpdateUserInfoRequest.Email into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskIn_Email() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoRequest, "email")

	return x
}

// MaskedIn_Email indicates the field UpdateUserInfoRequest.Email
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedIn_Email() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("email")
}

// _fm_UpdateUserInfoResponse is built in variable for UpdateUserInfoResponse to call FieldMask.Append
var _fm_UpdateUserInfoResponse = new(UpdateUserInfoResponse)

// MaskOut_ProfileInfo append UpdateUserInfoResponse.ProfileInfo into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskOut_ProfileInfo() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoResponse, "profile_info")

	return x
}

// MaskedOut_ProfileInfo indicates the field UpdateUserInfoRequest.ProfileInfo
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedOut_ProfileInfo() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info")
}

// MaskOut_ProfileInfo_UserId append UpdateUserInfoResponse.UserId into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskOut_ProfileInfo_UserId() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoResponse, "profile_info.user_id")

	return x
}

// MaskedOut_ProfileInfo_UserId indicates the field UpdateUserInfoRequest.UserId
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedOut_ProfileInfo_UserId() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info.user_id")
}

// MaskOut_ProfileInfo_Name append UpdateUserInfoResponse.Name into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskOut_ProfileInfo_Name() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoResponse, "profile_info.name")

	return x
}

// MaskedOut_ProfileInfo_Name indicates the field UpdateUserInfoRequest.Name
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedOut_ProfileInfo_Name() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info.name")
}

// MaskOut_ProfileInfo_Email append UpdateUserInfoResponse.Email into
// UpdateUserInfoRequest.FieldMask.
func (x *UpdateUserInfoRequest) MaskOut_ProfileInfo_Email() *UpdateUserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UpdateUserInfoResponse, "profile_info.email")

	return x
}

// MaskedOut_ProfileInfo_Email indicates the field UpdateUserInfoRequest.Email
// exists in the UpdateUserInfoRequest.FieldMask or not.
func (x *UpdateUserInfoRequest_FieldMask) MaskedOut_ProfileInfo_Email() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("profile_info.email")
}

// Mask only affects the fields in the UpdateUserInfoRequest.
func (x *UpdateUserInfoRequest_FieldMask) Mask(m *UpdateUserInfoResponse) *UpdateUserInfoResponse {
	switch x.maskMode {
	case pbfieldmask.MaskMode_Filter:
		x.mask.Filter(m)
	case pbfieldmask.MaskMode_Prune:
		x.mask.Prune(m)
	}

	return m
}
