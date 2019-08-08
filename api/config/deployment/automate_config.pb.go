// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/deployment/automate_config.proto

package deployment

import (
	fmt "fmt"
	applications "github.com/chef/automate/api/config/applications"
	authn "github.com/chef/automate/api/config/authn"
	authz "github.com/chef/automate/api/config/authz"
	backup_gateway "github.com/chef/automate/api/config/backup_gateway"
	bifrost "github.com/chef/automate/api/config/bifrost"
	bookshelf "github.com/chef/automate/api/config/bookshelf"
	cereal "github.com/chef/automate/api/config/cereal"
	cfgmgmt "github.com/chef/automate/api/config/cfgmgmt"
	compliance "github.com/chef/automate/api/config/compliance"
	cs_nginx "github.com/chef/automate/api/config/cs_nginx"
	data_feed "github.com/chef/automate/api/config/data_feed"
	data_lifecycle "github.com/chef/automate/api/config/data_lifecycle"
	dex "github.com/chef/automate/api/config/dex"
	elasticsearch "github.com/chef/automate/api/config/elasticsearch"
	erchef "github.com/chef/automate/api/config/erchef"
	es_sidecar "github.com/chef/automate/api/config/es_sidecar"
	esgateway "github.com/chef/automate/api/config/esgateway"
	event "github.com/chef/automate/api/config/event"
	event_feed "github.com/chef/automate/api/config/event_feed"
	event_gateway "github.com/chef/automate/api/config/event_gateway"
	gateway "github.com/chef/automate/api/config/gateway"
	ingest "github.com/chef/automate/api/config/ingest"
	license_control "github.com/chef/automate/api/config/license_control"
	load_balancer "github.com/chef/automate/api/config/load_balancer"
	local_user "github.com/chef/automate/api/config/local_user"
	nodemanager "github.com/chef/automate/api/config/nodemanager"
	notifications "github.com/chef/automate/api/config/notifications"
	pg_gateway "github.com/chef/automate/api/config/pg_gateway"
	pg_sidecar "github.com/chef/automate/api/config/pg_sidecar"
	postgresql "github.com/chef/automate/api/config/postgresql"
	prometheus "github.com/chef/automate/api/config/prometheus"
	secrets "github.com/chef/automate/api/config/secrets"
	session "github.com/chef/automate/api/config/session"
	shared "github.com/chef/automate/api/config/shared"
	teams "github.com/chef/automate/api/config/teams"
	ui "github.com/chef/automate/api/config/ui"
	workflow_nginx "github.com/chef/automate/api/config/workflow_nginx"
	workflow_server "github.com/chef/automate/api/config/workflow_server"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AutomateConfig struct {
	Global               *shared.GlobalConfig           `protobuf:"bytes,19,opt,name=global,proto3" json:"global,omitempty" toml:"global,omitempty" mapstructure:"global,omitempty"`
	AuthN                *authn.ConfigRequest           `protobuf:"bytes,1,opt,name=auth_n,json=authN,proto3" json:"auth_n,omitempty" toml:"auth_n,omitempty" mapstructure:"auth_n,omitempty"`
	AuthZ                *authz.ConfigRequest           `protobuf:"bytes,2,opt,name=auth_z,json=authZ,proto3" json:"auth_z,omitempty" toml:"auth_z,omitempty" mapstructure:"auth_z,omitempty"`
	Compliance           *compliance.ConfigRequest      `protobuf:"bytes,10,opt,name=compliance,proto3" json:"compliance,omitempty" toml:"compliance,omitempty" mapstructure:"compliance,omitempty"`
	ConfigMgmt           *cfgmgmt.ConfigRequest         `protobuf:"bytes,6,opt,name=config_mgmt,json=configMgmt,proto3" json:"config_mgmt,omitempty" toml:"config_mgmt,omitempty" mapstructure:"config_mgmt,omitempty"`
	Deployment           *ConfigRequest                 `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty" toml:"deployment,omitempty" mapstructure:"deployment,omitempty"`
	Dex                  *dex.ConfigRequest             `protobuf:"bytes,4,opt,name=dex,proto3" json:"dex,omitempty" toml:"dex,omitempty" mapstructure:"dex,omitempty"`
	Elasticsearch        *elasticsearch.ConfigRequest   `protobuf:"bytes,7,opt,name=elasticsearch,proto3" json:"elasticsearch,omitempty" toml:"elasticsearch,omitempty" mapstructure:"elasticsearch,omitempty"`
	Esgateway            *esgateway.ConfigRequest       `protobuf:"bytes,31,opt,name=esgateway,proto3" json:"esgateway,omitempty" toml:"esgateway,omitempty" mapstructure:"esgateway,omitempty"`
	EsSidecar            *es_sidecar.ConfigRequest      `protobuf:"bytes,11,opt,name=es_sidecar,json=esSidecar,proto3" json:"es_sidecar,omitempty" toml:"es_sidecar,omitempty" mapstructure:"es_sidecar,omitempty"`
	Gateway              *gateway.ConfigRequest         `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty" toml:"gateway,omitempty" mapstructure:"gateway,omitempty"`
	Ingest               *ingest.ConfigRequest          `protobuf:"bytes,13,opt,name=ingest,proto3" json:"ingest,omitempty" toml:"ingest,omitempty" mapstructure:"ingest,omitempty"`
	LoadBalancer         *load_balancer.ConfigRequest   `protobuf:"bytes,8,opt,name=load_balancer,json=loadBalancer,proto3" json:"load_balancer,omitempty" toml:"load_balancer,omitempty" mapstructure:"load_balancer,omitempty"`
	LocalUser            *local_user.ConfigRequest      `protobuf:"bytes,12,opt,name=local_user,json=localUser,proto3" json:"local_user,omitempty" toml:"local_user,omitempty" mapstructure:"local_user,omitempty"`
	LicenseControl       *license_control.ConfigRequest `protobuf:"bytes,16,opt,name=license_control,json=licenseControl,proto3" json:"license_control,omitempty" toml:"license_control,omitempty" mapstructure:"license_control,omitempty"`
	Notifications        *notifications.ConfigRequest   `protobuf:"bytes,14,opt,name=notifications,proto3" json:"notifications,omitempty" toml:"notifications,omitempty" mapstructure:"notifications,omitempty"`
	Postgresql           *postgresql.ConfigRequest      `protobuf:"bytes,15,opt,name=postgresql,proto3" json:"postgresql,omitempty" toml:"postgresql,omitempty" mapstructure:"postgresql,omitempty"`
	Session              *session.ConfigRequest         `protobuf:"bytes,17,opt,name=session,proto3" json:"session,omitempty" toml:"session,omitempty" mapstructure:"session,omitempty"`
	Teams                *teams.ConfigRequest           `protobuf:"bytes,18,opt,name=teams,proto3" json:"teams,omitempty" toml:"teams,omitempty" mapstructure:"teams,omitempty"`
	UI                   *ui.ConfigRequest              `protobuf:"bytes,9,opt,name=u_i,json=uI,proto3" json:"u_i,omitempty" toml:"u_i,omitempty" mapstructure:"u_i,omitempty"`
	DataLifecycle        *data_lifecycle.ConfigRequest  `protobuf:"bytes,20,opt,name=data_lifecycle,json=dataLifecycle,proto3" json:"data_lifecycle,omitempty" toml:"data_lifecycle,omitempty" mapstructure:"data_lifecycle,omitempty"`
	Secrets              *secrets.ConfigRequest         `protobuf:"bytes,21,opt,name=secrets,proto3" json:"secrets,omitempty" toml:"secrets,omitempty" mapstructure:"secrets,omitempty"`
	BackupGateway        *backup_gateway.ConfigRequest  `protobuf:"bytes,29,opt,name=backup_gateway,json=backupGateway,proto3" json:"backup_gateway,omitempty" toml:"backup_gateway,omitempty" mapstructure:"backup_gateway,omitempty"`
	PgSidecar            *pg_sidecar.ConfigRequest      `protobuf:"bytes,35,opt,name=pg_sidecar,json=pgSidecar,proto3" json:"pg_sidecar,omitempty" toml:"pg_sidecar,omitempty" mapstructure:"pg_sidecar,omitempty"`
	PgGateway            *pg_gateway.ConfigRequest      `protobuf:"bytes,34,opt,name=pg_gateway,json=pgGateway,proto3" json:"pg_gateway,omitempty" toml:"pg_gateway,omitempty" mapstructure:"pg_gateway,omitempty"`
	Applications         *applications.ConfigRequest    `protobuf:"bytes,36,opt,name=applications,proto3" json:"applications,omitempty" toml:"applications,omitempty" mapstructure:"applications,omitempty"`
	Bookshelf            *bookshelf.ConfigRequest       `protobuf:"bytes,22,opt,name=bookshelf,proto3" json:"bookshelf,omitempty" toml:"bookshelf,omitempty" mapstructure:"bookshelf,omitempty"`
	Bifrost              *bifrost.ConfigRequest         `protobuf:"bytes,23,opt,name=bifrost,proto3" json:"bifrost,omitempty" toml:"bifrost,omitempty" mapstructure:"bifrost,omitempty"`
	Erchef               *erchef.ConfigRequest          `protobuf:"bytes,24,opt,name=erchef,proto3" json:"erchef,omitempty" toml:"erchef,omitempty" mapstructure:"erchef,omitempty"`
	CsNginx              *cs_nginx.ConfigRequest        `protobuf:"bytes,25,opt,name=cs_nginx,json=csNginx,proto3" json:"cs_nginx,omitempty" toml:"cs_nginx,omitempty" mapstructure:"cs_nginx,omitempty"`
	Workflow             *workflow_server.ConfigRequest `protobuf:"bytes,27,opt,name=workflow,proto3" json:"workflow,omitempty" toml:"workflow,omitempty" mapstructure:"workflow,omitempty"`
	WorkflowNginx        *workflow_nginx.ConfigRequest  `protobuf:"bytes,28,opt,name=workflow_nginx,json=workflowNginx,proto3" json:"workflow_nginx,omitempty" toml:"workflow_nginx,omitempty" mapstructure:"workflow_nginx,omitempty"`
	EventService         *event.ConfigRequest           `protobuf:"bytes,30,opt,name=event_service,json=eventService,proto3" json:"event_service,omitempty" toml:"event_service,omitempty" mapstructure:"event_service,omitempty"`
	Nodemanager          *nodemanager.ConfigRequest     `protobuf:"bytes,33,opt,name=nodemanager,proto3" json:"nodemanager,omitempty" toml:"nodemanager,omitempty" mapstructure:"nodemanager,omitempty"`
	EventGateway         *event_gateway.ConfigRequest   `protobuf:"bytes,37,opt,name=event_gateway,json=eventGateway,proto3" json:"event_gateway,omitempty" toml:"event_gateway,omitempty" mapstructure:"event_gateway,omitempty"`
	Prometheus           *prometheus.ConfigRequest      `protobuf:"bytes,32,opt,name=prometheus,proto3" json:"prometheus,omitempty" toml:"prometheus,omitempty" mapstructure:"prometheus,omitempty"`
	DataFeedService      *data_feed.ConfigRequest       `protobuf:"bytes,38,opt,name=data_feed_service,json=dataFeedService,proto3" json:"data_feed_service,omitempty" toml:"data_feed_service,omitempty" mapstructure:"data_feed_service,omitempty"`
	EventFeedService     *event_feed.ConfigRequest      `protobuf:"bytes,39,opt,name=event_feed_service,json=eventFeedService,proto3" json:"event_feed_service,omitempty" toml:"event_feed_service,omitempty" mapstructure:"event_feed_service,omitempty"`
	Cereal               *cereal.ConfigRequest          `protobuf:"bytes,40,opt,name=cereal,proto3" json:"cereal,omitempty" toml:"cereal,omitempty" mapstructure:"cereal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                          `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *AutomateConfig) Reset()         { *m = AutomateConfig{} }
func (m *AutomateConfig) String() string { return proto.CompactTextString(m) }
func (*AutomateConfig) ProtoMessage()    {}
func (*AutomateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff33eb1f2117d94c, []int{0}
}

func (m *AutomateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutomateConfig.Unmarshal(m, b)
}
func (m *AutomateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutomateConfig.Marshal(b, m, deterministic)
}
func (m *AutomateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutomateConfig.Merge(m, src)
}
func (m *AutomateConfig) XXX_Size() int {
	return xxx_messageInfo_AutomateConfig.Size(m)
}
func (m *AutomateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AutomateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AutomateConfig proto.InternalMessageInfo

func (m *AutomateConfig) GetGlobal() *shared.GlobalConfig {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *AutomateConfig) GetAuthN() *authn.ConfigRequest {
	if m != nil {
		return m.AuthN
	}
	return nil
}

func (m *AutomateConfig) GetAuthZ() *authz.ConfigRequest {
	if m != nil {
		return m.AuthZ
	}
	return nil
}

func (m *AutomateConfig) GetCompliance() *compliance.ConfigRequest {
	if m != nil {
		return m.Compliance
	}
	return nil
}

func (m *AutomateConfig) GetConfigMgmt() *cfgmgmt.ConfigRequest {
	if m != nil {
		return m.ConfigMgmt
	}
	return nil
}

func (m *AutomateConfig) GetDeployment() *ConfigRequest {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *AutomateConfig) GetDex() *dex.ConfigRequest {
	if m != nil {
		return m.Dex
	}
	return nil
}

func (m *AutomateConfig) GetElasticsearch() *elasticsearch.ConfigRequest {
	if m != nil {
		return m.Elasticsearch
	}
	return nil
}

func (m *AutomateConfig) GetEsgateway() *esgateway.ConfigRequest {
	if m != nil {
		return m.Esgateway
	}
	return nil
}

func (m *AutomateConfig) GetEsSidecar() *es_sidecar.ConfigRequest {
	if m != nil {
		return m.EsSidecar
	}
	return nil
}

func (m *AutomateConfig) GetGateway() *gateway.ConfigRequest {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *AutomateConfig) GetIngest() *ingest.ConfigRequest {
	if m != nil {
		return m.Ingest
	}
	return nil
}

func (m *AutomateConfig) GetLoadBalancer() *load_balancer.ConfigRequest {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *AutomateConfig) GetLocalUser() *local_user.ConfigRequest {
	if m != nil {
		return m.LocalUser
	}
	return nil
}

func (m *AutomateConfig) GetLicenseControl() *license_control.ConfigRequest {
	if m != nil {
		return m.LicenseControl
	}
	return nil
}

func (m *AutomateConfig) GetNotifications() *notifications.ConfigRequest {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *AutomateConfig) GetPostgresql() *postgresql.ConfigRequest {
	if m != nil {
		return m.Postgresql
	}
	return nil
}

func (m *AutomateConfig) GetSession() *session.ConfigRequest {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *AutomateConfig) GetTeams() *teams.ConfigRequest {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *AutomateConfig) GetUI() *ui.ConfigRequest {
	if m != nil {
		return m.UI
	}
	return nil
}

func (m *AutomateConfig) GetDataLifecycle() *data_lifecycle.ConfigRequest {
	if m != nil {
		return m.DataLifecycle
	}
	return nil
}

func (m *AutomateConfig) GetSecrets() *secrets.ConfigRequest {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *AutomateConfig) GetBackupGateway() *backup_gateway.ConfigRequest {
	if m != nil {
		return m.BackupGateway
	}
	return nil
}

func (m *AutomateConfig) GetPgSidecar() *pg_sidecar.ConfigRequest {
	if m != nil {
		return m.PgSidecar
	}
	return nil
}

func (m *AutomateConfig) GetPgGateway() *pg_gateway.ConfigRequest {
	if m != nil {
		return m.PgGateway
	}
	return nil
}

func (m *AutomateConfig) GetApplications() *applications.ConfigRequest {
	if m != nil {
		return m.Applications
	}
	return nil
}

func (m *AutomateConfig) GetBookshelf() *bookshelf.ConfigRequest {
	if m != nil {
		return m.Bookshelf
	}
	return nil
}

func (m *AutomateConfig) GetBifrost() *bifrost.ConfigRequest {
	if m != nil {
		return m.Bifrost
	}
	return nil
}

func (m *AutomateConfig) GetErchef() *erchef.ConfigRequest {
	if m != nil {
		return m.Erchef
	}
	return nil
}

func (m *AutomateConfig) GetCsNginx() *cs_nginx.ConfigRequest {
	if m != nil {
		return m.CsNginx
	}
	return nil
}

func (m *AutomateConfig) GetWorkflow() *workflow_server.ConfigRequest {
	if m != nil {
		return m.Workflow
	}
	return nil
}

func (m *AutomateConfig) GetWorkflowNginx() *workflow_nginx.ConfigRequest {
	if m != nil {
		return m.WorkflowNginx
	}
	return nil
}

func (m *AutomateConfig) GetEventService() *event.ConfigRequest {
	if m != nil {
		return m.EventService
	}
	return nil
}

func (m *AutomateConfig) GetNodemanager() *nodemanager.ConfigRequest {
	if m != nil {
		return m.Nodemanager
	}
	return nil
}

func (m *AutomateConfig) GetEventGateway() *event_gateway.ConfigRequest {
	if m != nil {
		return m.EventGateway
	}
	return nil
}

func (m *AutomateConfig) GetPrometheus() *prometheus.ConfigRequest {
	if m != nil {
		return m.Prometheus
	}
	return nil
}

func (m *AutomateConfig) GetDataFeedService() *data_feed.ConfigRequest {
	if m != nil {
		return m.DataFeedService
	}
	return nil
}

func (m *AutomateConfig) GetEventFeedService() *event_feed.ConfigRequest {
	if m != nil {
		return m.EventFeedService
	}
	return nil
}

func (m *AutomateConfig) GetCereal() *cereal.ConfigRequest {
	if m != nil {
		return m.Cereal
	}
	return nil
}

func init() {
	proto.RegisterType((*AutomateConfig)(nil), "chef.automate.domain.deployment.AutomateConfig")
}

func init() {
	proto.RegisterFile("api/config/deployment/automate_config.proto", fileDescriptor_ff33eb1f2117d94c)
}

var fileDescriptor_ff33eb1f2117d94c = []byte{
	// 1197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x98, 0x69, 0x73, 0xd4, 0x36,
	0x18, 0xc7, 0x07, 0x42, 0x2e, 0xe5, 0x44, 0x3d, 0x50, 0x43, 0x4b, 0x68, 0x28, 0xe4, 0x80, 0x78,
	0x53, 0x98, 0xbe, 0xe8, 0xab, 0x36, 0x61, 0xda, 0x4c, 0x69, 0x48, 0xa7, 0xd0, 0xb4, 0x85, 0x61,
	0xc6, 0xa3, 0xf5, 0x6a, 0xbd, 0x9a, 0xd8, 0x96, 0xb1, 0x6c, 0x72, 0x7c, 0xe6, 0x7e, 0x88, 0x8e,
	0x25, 0x1f, 0x92, 0x22, 0xcb, 0xfb, 0x72, 0xf5, 0xfc, 0xf4, 0x9f, 0xc7, 0x8f, 0x9e, 0x43, 0x5a,
	0xf0, 0x14, 0xa7, 0x74, 0x10, 0xb0, 0x64, 0x4c, 0xc3, 0xc1, 0x88, 0xa4, 0x11, 0xbb, 0x8a, 0x49,
	0x92, 0x0f, 0x70, 0x91, 0xb3, 0x18, 0xe7, 0xc4, 0x97, 0x26, 0x2f, 0xcd, 0x58, 0xce, 0xe0, 0x66,
	0x30, 0x21, 0x63, 0xaf, 0xb6, 0x79, 0x23, 0x16, 0x63, 0x9a, 0x78, 0xed, 0xb6, 0x8d, 0x67, 0x8a,
	0x1a, 0x4e, 0xd3, 0x88, 0x06, 0x38, 0xa7, 0x2c, 0xe1, 0xd5, 0x9a, 0x9f, 0x91, 0x8f, 0x05, 0xe1,
	0xb9, 0x94, 0xdb, 0x78, 0xac, 0xd2, 0x45, 0x3e, 0x49, 0xa6, 0xc3, 0xae, 0xed, 0x98, 0xa7, 0x60,
	0x43, 0x1c, 0x9c, 0x17, 0xa9, 0x1f, 0xe2, 0x9c, 0x5c, 0xe0, 0x2b, 0x3b, 0xbf, 0xad, 0xf2, 0x74,
	0x9c, 0x31, 0x9e, 0xdb, 0xc1, 0x5d, 0x15, 0x64, 0xec, 0x9c, 0x4f, 0x48, 0x34, 0xb6, 0xa3, 0x4f,
	0x14, 0x94, 0x64, 0x65, 0xb4, 0xfa, 0x3f, 0x89, 0x7c, 0x2a, 0x03, 0xde, 0xeb, 0x62, 0x30, 0x0e,
	0xe3, 0x30, 0xee, 0x00, 0xf7, 0x54, 0x90, 0xc5, 0x69, 0x44, 0x71, 0x12, 0x10, 0x3b, 0xfb, 0x40,
	0x61, 0xf9, 0x04, 0x67, 0x64, 0x34, 0x08, 0x23, 0x36, 0xc4, 0x51, 0x65, 0xdf, 0x51, 0xb5, 0xb8,
	0x9f, 0x84, 0x34, 0xb9, 0xec, 0x8f, 0xf8, 0x08, 0xe7, 0xd8, 0x8f, 0xe8, 0x98, 0x04, 0x57, 0x41,
	0x44, 0xfa, 0xbd, 0x54, 0x72, 0xcd, 0xca, 0x3e, 0xd2, 0xd8, 0x0e, 0x07, 0xf6, 0xd5, 0x30, 0x46,
	0x98, 0xe7, 0x34, 0xe0, 0x04, 0x67, 0xc1, 0xa4, 0xff, 0x20, 0x09, 0x77, 0x26, 0xc7, 0x9e, 0x86,
	0xfa, 0x9c, 0x8e, 0x48, 0x80, 0xb3, 0xfe, 0x53, 0x72, 0x8a, 0xaa, 0xd9, 0x41, 0x93, 0x90, 0x74,
	0x25, 0xdc, 0x40, 0xe1, 0x22, 0x1a, 0x90, 0x84, 0x8b, 0x3a, 0xcc, 0x33, 0x16, 0xf5, 0xc7, 0x21,
	0x62, 0x78, 0xe4, 0x0f, 0x71, 0x54, 0x26, 0x40, 0xd6, 0xff, 0x71, 0x11, 0x0b, 0x70, 0xe4, 0x17,
	0xbc, 0x8b, 0x55, 0xfb, 0x43, 0xc2, 0x46, 0x24, 0xc6, 0x09, 0x0e, 0xbb, 0xe0, 0x7d, 0x0d, 0xce,
	0xe9, 0xd8, 0x5d, 0xff, 0xaa, 0x1f, 0x29, 0xe3, 0x79, 0x98, 0x11, 0xfe, 0x31, 0x9a, 0x82, 0x0d,
	0xfd, 0xa9, 0x0f, 0x2f, 0x0d, 0xdd, 0x87, 0xa7, 0xb1, 0x19, 0x8b, 0x49, 0x3e, 0x21, 0x05, 0xef,
	0x3f, 0x68, 0x4e, 0x82, 0x8c, 0xe4, 0x53, 0x81, 0x9c, 0x53, 0x36, 0x45, 0x6b, 0xcb, 0x09, 0x8e,
	0x3b, 0xf4, 0xb6, 0x14, 0xac, 0xa0, 0xfd, 0xc5, 0x78, 0xc1, 0xb2, 0xf3, 0x71, 0xc4, 0x2e, 0x5c,
	0xc5, 0x3b, 0xb0, 0xf1, 0x9c, 0x64, 0x9f, 0xa6, 0x39, 0x5c, 0xd1, 0xb3, 0xdc, 0x87, 0xb0, 0x6b,
	0x36, 0x87, 0x31, 0x21, 0xa3, 0x29, 0x8a, 0x4d, 0x28, 0x77, 0xb3, 0x6a, 0x0d, 0x05, 0x24, 0x23,
	0xd8, 0x9e, 0x2f, 0x5b, 0xff, 0xdd, 0x03, 0xab, 0x87, 0xd5, 0xa0, 0x7a, 0x29, 0x00, 0xf8, 0x33,
	0x98, 0x93, 0x8d, 0x0e, 0x7d, 0xf6, 0xf0, 0xd6, 0xce, 0xd2, 0xf3, 0x1d, 0x4f, 0x1f, 0x67, 0x34,
	0x19, 0x67, 0xd8, 0xab, 0x06, 0xde, 0xb1, 0x20, 0xe5, 0xce, 0x37, 0xd5, 0xbe, 0x52, 0xa1, 0x1c,
	0x40, 0x7e, 0x82, 0x6e, 0x09, 0x85, 0x5d, 0xcf, 0x3a, 0x10, 0xc5, 0x2c, 0xf3, 0xaa, 0xbd, 0xd2,
	0xab, 0x37, 0xb3, 0xe5, 0xe2, 0x69, 0xa3, 0x70, 0x8d, 0x6e, 0xf7, 0x29, 0x5c, 0xdb, 0x14, 0xde,
	0xc3, 0x33, 0x00, 0xda, 0x0e, 0x8f, 0x80, 0x50, 0xf9, 0xc1, 0xae, 0xd2, 0x72, 0xf5, 0x57, 0xe9,
	0x8a, 0x8a, 0x10, 0x3c, 0x01, 0x4b, 0x55, 0x1c, 0xcb, 0x29, 0x83, 0xe6, 0x84, 0xee, 0xd3, 0x0e,
	0x5d, 0x39, 0x8a, 0x6e, 0xaa, 0x95, 0x3f, 0x5f, 0x87, 0x71, 0x0e, 0x4f, 0x01, 0x68, 0x1b, 0x3c,
	0x9a, 0x11, 0x62, 0x9e, 0xd7, 0x73, 0x7b, 0x30, 0xf5, 0x5a, 0x0b, 0xfc, 0x11, 0xcc, 0x8c, 0xc8,
	0x25, 0xba, 0x23, 0x84, 0xb6, 0xbb, 0x84, 0x2e, 0x0d, 0x85, 0x72, 0x0f, 0xfc, 0x1b, 0xac, 0x68,
	0xa3, 0x01, 0xcd, 0x0b, 0x91, 0x03, 0xeb, 0xe1, 0x6b, 0xa4, 0xa1, 0xa6, 0xcb, 0xc0, 0x57, 0x60,
	0xb1, 0x99, 0x21, 0x68, 0x53, 0x68, 0x3e, 0xb3, 0x6b, 0xd6, 0x94, 0xa1, 0xd7, 0x6e, 0x87, 0x27,
	0x00, 0xb4, 0x43, 0x06, 0x2d, 0x09, 0xb1, 0xfd, 0x0e, 0xb1, 0x1a, 0xbb, 0xa9, 0xf6, 0x56, 0x1a,
	0xe0, 0x21, 0x98, 0xaf, 0xfd, 0x9a, 0xb5, 0x06, 0x0c, 0xa7, 0xd4, 0x9e, 0x10, 0xf5, 0x3e, 0x78,
	0x04, 0xe6, 0xe4, 0x80, 0x42, 0x2b, 0x42, 0x61, 0xcf, 0x1e, 0x72, 0xc9, 0x18, 0x22, 0xd5, 0x4e,
	0x78, 0x06, 0x56, 0xb4, 0x59, 0x84, 0x16, 0x1c, 0x81, 0xd7, 0x48, 0x43, 0x70, 0xb9, 0x34, 0x1e,
	0x55, 0x36, 0xf8, 0x1a, 0x80, 0x76, 0x66, 0xa1, 0x65, 0x57, 0x6a, 0xb5, 0x9c, 0x19, 0x2c, 0x61,
	0x39, 0xe3, 0x24, 0x83, 0x1f, 0xc0, 0x9a, 0x31, 0x62, 0xd1, 0xba, 0xd0, 0x7c, 0xd1, 0xa1, 0xa9,
	0xc3, 0x86, 0xf0, 0x6a, 0x65, 0x7e, 0x29, 0xad, 0xf0, 0x1f, 0xb0, 0xa2, 0xcd, 0x41, 0xb4, 0x2a,
	0xb4, 0xbf, 0xb7, 0x6b, 0x6b, 0xa8, 0x99, 0x7d, 0x9a, 0xb1, 0x8c, 0x42, 0x3b, 0x31, 0xd1, 0x9a,
	0x23, 0x63, 0x5a, 0xcc, 0xac, 0xaf, 0xd6, 0x02, 0x7f, 0x01, 0xf3, 0xd5, 0x9c, 0x42, 0x77, 0x5d,
	0x95, 0x5f, 0x41, 0x66, 0xda, 0x54, 0xcb, 0xf0, 0x27, 0x30, 0x2b, 0xa6, 0x18, 0x82, 0xae, 0xe6,
	0x26, 0x10, 0xb3, 0xb9, 0x89, 0x45, 0x78, 0x00, 0x66, 0x0a, 0x9f, 0xa2, 0x45, 0xb1, 0x7d, 0xd3,
	0xd8, 0x5e, 0x50, 0x63, 0xd3, 0xed, 0xe2, 0x37, 0xf8, 0x0e, 0xac, 0xea, 0x57, 0x4f, 0xf4, 0xb9,
	0xd8, 0xfc, 0xbc, 0xa3, 0x49, 0x68, 0xac, 0x19, 0xe3, 0xd2, 0x7a, 0x52, 0x1b, 0x65, 0x50, 0xc4,
	0x94, 0x47, 0x5f, 0xb8, 0x83, 0x22, 0xa0, 0x9b, 0x41, 0x11, 0xcb, 0xf0, 0x5f, 0xb0, 0xaa, 0x3f,
	0x47, 0xd0, 0x37, 0xd6, 0x24, 0x90, 0xc7, 0xa5, 0xa3, 0xa6, 0x83, 0xd2, 0x7a, 0xdc, 0xb6, 0x8d,
	0xf6, 0x7a, 0x83, 0x1e, 0xb9, 0x92, 0x20, 0xec, 0x6a, 0x1b, 0x69, 0x58, 0xb7, 0x0d, 0xa9, 0x56,
	0xfb, 0xb8, 0xe5, 0x56, 0xeb, 0x68, 0x69, 0x69, 0x58, 0xfb, 0xf6, 0x17, 0x58, 0x56, 0x1f, 0x80,
	0xe8, 0x3b, 0x6b, 0xf1, 0xd7, 0xe3, 0x4e, 0x21, 0xcd, 0xe2, 0x57, 0x6d, 0xf0, 0x77, 0xb0, 0xd8,
	0xbc, 0xc0, 0xd0, 0x97, 0x56, 0x17, 0x2b, 0xc9, 0x06, 0x33, 0x5d, 0x6c, 0x0c, 0xe5, 0xf9, 0x56,
	0xef, 0x3e, 0x74, 0xcf, 0x75, 0xbe, 0x15, 0x64, 0x9e, 0x6f, 0xb5, 0x5c, 0xf6, 0x4a, 0xf9, 0xd4,
	0x43, 0xc8, 0xd5, 0x2b, 0x25, 0x63, 0xf6, 0x4a, 0xb9, 0x0a, 0x8f, 0xc1, 0x42, 0xfd, 0xd4, 0x42,
	0x5f, 0x59, 0x67, 0x49, 0x3d, 0x7a, 0x2b, 0xca, 0x74, 0x26, 0xe0, 0xa7, 0xe5, 0x32, 0xfc, 0x03,
	0x2c, 0xd4, 0x97, 0x39, 0x74, 0xdf, 0xd5, 0xc7, 0x8c, 0x2b, 0x9f, 0xa1, 0xd7, 0x88, 0x94, 0xf5,
	0xa5, 0xdf, 0x26, 0xd1, 0xd7, 0xae, 0xfa, 0xd2, 0x59, 0x33, 0x7d, 0x6b, 0xab, 0xf4, 0xf5, 0x14,
	0xac, 0xc8, 0xdb, 0x5e, 0xe9, 0x02, 0x0d, 0x08, 0x7a, 0xe0, 0xea, 0x1a, 0x02, 0x35, 0x93, 0x43,
	0x2c, 0xbe, 0x95, 0xdb, 0xe1, 0x9f, 0x60, 0x49, 0x79, 0xa1, 0xa0, 0x6f, 0x85, 0xda, 0xa0, 0xab,
	0xd5, 0x36, 0xa0, 0xa1, 0xa9, 0x6a, 0x88, 0xcb, 0x83, 0x7a, 0xd5, 0x45, 0x8f, 0x5d, 0xfd, 0x5b,
	0x43, 0xad, 0xae, 0xd6, 0xd5, 0x51, 0xb6, 0xef, 0xe6, 0xb1, 0x81, 0x1e, 0xba, 0x6a, 0xad, 0xc1,
	0x6e, 0xb4, 0xef, 0xc6, 0x02, 0xdf, 0x81, 0xbb, 0xcd, 0x15, 0xbb, 0x89, 0xe6, 0x13, 0x57, 0x79,
	0x34, 0xb8, 0xa1, 0xba, 0x56, 0x1a, 0x7e, 0x25, 0x64, 0x54, 0x07, 0xf5, 0x03, 0x80, 0xed, 0x95,
	0xbc, 0xd1, 0xde, 0x76, 0x8d, 0xdd, 0x96, 0x37, 0xc4, 0xd7, 0x85, 0x45, 0x55, 0x3f, 0x04, 0x73,
	0xf2, 0x12, 0x8f, 0x76, 0xac, 0x67, 0x5f, 0x5d, 0xc9, 0x05, 0x62, 0x96, 0x8e, 0x5c, 0x7d, 0x75,
	0x67, 0x61, 0x63, 0xfd, 0xfe, 0xd1, 0xc1, 0x7b, 0x2f, 0xa4, 0xf9, 0xa4, 0x18, 0x96, 0x97, 0xde,
	0x81, 0xf8, 0xe3, 0xa5, 0x16, 0x19, 0x58, 0xff, 0x75, 0x18, 0xce, 0x89, 0x77, 0xc2, 0x8b, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0x00, 0x4c, 0xb5, 0x01, 0x13, 0x00, 0x00,
}
