// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: interservice/teams/teams.proto

package teams

import (
	"context"

	version "github.com/chef/automate/api/external/common/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TeamsServiceServer interface (at compile time)
var _ TeamsServiceServer = &TeamsServiceServerMock{}

// NewTeamsServiceServerMock gives you a fresh instance of TeamsServiceServerMock.
func NewTeamsServiceServerMock() *TeamsServiceServerMock {
	return &TeamsServiceServerMock{validateRequests: true}
}

// NewTeamsServiceServerMockWithoutValidation gives you a fresh instance of
// TeamsServiceServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTeamsServiceServerMockWithoutValidation() *TeamsServiceServerMock {
	return &TeamsServiceServerMock{}
}

// TeamsServiceServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TeamsServiceServerMock struct {
	validateRequests        bool
	GetVersionFunc          func(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	GetTeamFunc             func(context.Context, *GetTeamReq) (*GetTeamResp, error)
	ListTeamsFunc           func(context.Context, *ListTeamsReq) (*ListTeamsResp, error)
	CreateTeamFunc          func(context.Context, *CreateTeamReq) (*CreateTeamResp, error)
	UpdateTeamFunc          func(context.Context, *UpdateTeamReq) (*UpdateTeamResp, error)
	DeleteTeamFunc          func(context.Context, *DeleteTeamReq) (*DeleteTeamResp, error)
	AddTeamMembersFunc      func(context.Context, *AddTeamMembersReq) (*AddTeamMembersResp, error)
	RemoveTeamMembersFunc   func(context.Context, *RemoveTeamMembersReq) (*RemoveTeamMembersResp, error)
	GetTeamsForMemberFunc   func(context.Context, *GetTeamsForMemberReq) (*GetTeamsForMemberResp, error)
	GetTeamMembershipFunc   func(context.Context, *GetTeamMembershipReq) (*GetTeamMembershipResp, error)
	PurgeUserMembershipFunc func(context.Context, *PurgeUserMembershipReq) (*PurgeUserMembershipResp, error)
}

func (m *TeamsServiceServerMock) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetVersionFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetVersion' not implemented")
}

func (m *TeamsServiceServerMock) GetTeam(ctx context.Context, req *GetTeamReq) (*GetTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeam' not implemented")
}

func (m *TeamsServiceServerMock) ListTeams(ctx context.Context, req *ListTeamsReq) (*ListTeamsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListTeamsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListTeams' not implemented")
}

func (m *TeamsServiceServerMock) CreateTeam(ctx context.Context, req *CreateTeamReq) (*CreateTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateTeam' not implemented")
}

func (m *TeamsServiceServerMock) UpdateTeam(ctx context.Context, req *UpdateTeamReq) (*UpdateTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateTeam' not implemented")
}

func (m *TeamsServiceServerMock) DeleteTeam(ctx context.Context, req *DeleteTeamReq) (*DeleteTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteTeam' not implemented")
}

func (m *TeamsServiceServerMock) AddTeamMembers(ctx context.Context, req *AddTeamMembersReq) (*AddTeamMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.AddTeamMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'AddTeamMembers' not implemented")
}

func (m *TeamsServiceServerMock) RemoveTeamMembers(ctx context.Context, req *RemoveTeamMembersReq) (*RemoveTeamMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.RemoveTeamMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'RemoveTeamMembers' not implemented")
}

func (m *TeamsServiceServerMock) GetTeamsForMember(ctx context.Context, req *GetTeamsForMemberReq) (*GetTeamsForMemberResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamsForMemberFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeamsForMember' not implemented")
}

func (m *TeamsServiceServerMock) GetTeamMembership(ctx context.Context, req *GetTeamMembershipReq) (*GetTeamMembershipResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamMembershipFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeamMembership' not implemented")
}

func (m *TeamsServiceServerMock) PurgeUserMembership(ctx context.Context, req *PurgeUserMembershipReq) (*PurgeUserMembershipResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.PurgeUserMembershipFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'PurgeUserMembership' not implemented")
}

// Reset resets all overridden functions
func (m *TeamsServiceServerMock) Reset() {
	m.GetVersionFunc = nil
	m.GetTeamFunc = nil
	m.ListTeamsFunc = nil
	m.CreateTeamFunc = nil
	m.UpdateTeamFunc = nil
	m.DeleteTeamFunc = nil
	m.AddTeamMembersFunc = nil
	m.RemoveTeamMembersFunc = nil
	m.GetTeamsForMemberFunc = nil
	m.GetTeamMembershipFunc = nil
	m.PurgeUserMembershipFunc = nil
}
