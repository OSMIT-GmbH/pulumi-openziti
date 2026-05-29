package provider

import (
	"context"
	"github.com/openziti/edge-api/rest_model"
	"github.com/pulumi/pulumi-go-provider/infer"
)

// ─── V0 base types with old pulumi:"id" tag ─────────────────────

type BaseStateEntityV0 struct {
	Links       Links  `pulumi:"_links"`
	CreatedAt   string `pulumi:"createdAt"`
	ID          string `pulumi:"id"`
	Tags        Tags   `pulumi:"tags,optional"`
	UpdatedAt   string `pulumi:"updatedAt"`
	Assimilated bool   `pulumi:"_assimilated"`
}

func (v0 BaseStateEntityV0) toV1() BaseStateEntity {
	return BaseStateEntity{
		Links:       v0.Links,
		CreatedAt:   v0.CreatedAt,
		ID:          v0.ID,
		Tags:        v0.Tags,
		UpdatedAt:   v0.UpdatedAt,
		Assimilated: v0.Assimilated,
	}
}

type EntityRefV0 struct {
	Links  Links  `pulumi:"_links"`
	Entity string `pulumi:"entity,optional"`
	ID     string `pulumi:"id,optional"`
	Name   string `pulumi:"name,optional"`
}

func (v0 EntityRefV0) toV1() EntityRef {
	return EntityRef{
		Links:  v0.Links,
		Entity: v0.Entity,
		ID:     v0.ID,
		Name:   v0.Name,
	}
}

// ─── Identity enrollment V0 types ──────────────────────────────

type IdentityEnrollmentsOttV0 struct {
	ExpiresAt string `pulumi:"expiresAt,optional"`
	ID        string `pulumi:"id,optional"`
	JWT       string `pulumi:"jwt,optional"`
	Token     string `pulumi:"token,optional"`
}

func (v0 IdentityEnrollmentsOttV0) toV1() IdentityEnrollmentsOtt {
	return IdentityEnrollmentsOtt{
		ExpiresAt: v0.ExpiresAt,
		ID:        v0.ID,
		JWT:       v0.JWT,
		Token:     v0.Token,
	}
}

type IdentityEnrollmentsOttcaV0 struct {
	Ca        *EntityRefV0 `pulumi:"ca,optional"`
	CaID      string       `pulumi:"caId,optional"`
	ExpiresAt string       `pulumi:"expiresAt,optional"`
	ID        string       `pulumi:"id,optional"`
	JWT       string       `pulumi:"jwt,optional"`
	Token     string       `pulumi:"token,optional"`
}

func (v0 IdentityEnrollmentsOttcaV0) toV1() IdentityEnrollmentsOttca {
	return IdentityEnrollmentsOttca{
		Ca: func() *EntityRef {
			if v0.Ca == nil {
				return nil
			}
			v := v0.Ca.toV1()
			return &v
		}(),
		CaID:      v0.CaID,
		ExpiresAt: v0.ExpiresAt,
		ID:        v0.ID,
		JWT:       v0.JWT,
		Token:     v0.Token,
	}
}

type IdentityEnrollmentsUpdbV0 struct {
	ExpiresAt string `pulumi:"expiresAt,optional"`
	ID        string `pulumi:"id,optional"`
	JWT       string `pulumi:"jwt,optional"`
	Token     string `pulumi:"token,optional"`
}

func (v0 IdentityEnrollmentsUpdbV0) toV1() IdentityEnrollmentsUpdb {
	return IdentityEnrollmentsUpdb{
		ExpiresAt: v0.ExpiresAt,
		ID:        v0.ID,
		JWT:       v0.JWT,
		Token:     v0.Token,
	}
}

type IdentityEnrollmentsV0 struct {
	Ott   *IdentityEnrollmentsOttV0   `pulumi:"ott,optional"`
	Ottca *IdentityEnrollmentsOttcaV0 `pulumi:"ottca,optional"`
	Updb  *IdentityEnrollmentsUpdbV0  `pulumi:"updb,optional"`
}

func (v0 IdentityEnrollmentsV0) toV1() IdentityEnrollments {
	return IdentityEnrollments{
		Ott: func() *IdentityEnrollmentsOtt {
			if v0.Ott == nil {
				return nil
			}
			v := v0.Ott.toV1()
			return &v
		}(),
		Ottca: func() *IdentityEnrollmentsOttca {
			if v0.Ottca == nil {
				return nil
			}
			v := v0.Ottca.toV1()
			return &v
		}(),
		Updb: func() *IdentityEnrollmentsUpdb {
			if v0.Updb == nil {
				return nil
			}
			v := v0.Updb.toV1()
			return &v
		}(),
	}
}

// ─── ServicePolicy V0 ──────────────────────────────────────────

type ServicePolicyStateV0 struct {
	ServicePolicyArgs
	BaseStateEntityV0
	IdentityRoles            rest_model.Roles    `pulumi:"identityRoles"`
	IdentityRolesDisplay     NamedRoles          `pulumi:"identityRolesDisplay"`
	PostureCheckRoles        rest_model.Roles    `pulumi:"postureCheckRoles"`
	PostureCheckRolesDisplay NamedRoles          `pulumi:"postureCheckRolesDisplay"`
	Semantic                 rest_model.Semantic `pulumi:"semantic"`
	ServiceRoles             rest_model.Roles    `pulumi:"serviceRoles"`
	ServiceRolesDisplay      NamedRoles          `pulumi:"serviceRolesDisplay"`
	Type                     rest_model.DialBind `pulumi:"type"`
}

func migrateServicePolicyV0(_ context.Context, v0 ServicePolicyStateV0) (infer.MigrationResult[ServicePolicyState], error) {
	return infer.MigrationResult[ServicePolicyState]{
		Result: &ServicePolicyState{
			ServicePolicyArgs:          v0.ServicePolicyArgs,
			BaseStateEntity:            v0.BaseStateEntityV0.toV1(),
			IdentityRoles:              v0.IdentityRoles,
			IdentityRolesDisplay:       v0.IdentityRolesDisplay,
			PostureCheckRoles:          v0.PostureCheckRoles,
			PostureCheckRolesDisplay:   v0.PostureCheckRolesDisplay,
			Semantic:                   v0.Semantic,
			ServiceRoles:               v0.ServiceRoles,
			ServiceRolesDisplay:        v0.ServiceRolesDisplay,
			Type:                       v0.Type,
		},
	}, nil
}

// ─── Service V0 ────────────────────────────────────────────────

type ServiceStateV0 struct {
	ServiceArgs
	BaseStateEntityV0
	Config             map[string]map[string]interface{} `pulumi:"config"`
	Configs            []string                          `pulumi:"configs"`
	EncryptionRequired bool                              `pulumi:"encryptionRequired"`
	Name               string                            `pulumi:"name"`
	Permissions        rest_model.DialBindArray          `pulumi:"permissions"`
	PostureQueries     []PostureQueriesType              `pulumi:"postureQueries"`
	RoleAttributes     rest_model.Attributes             `pulumi:"roleAttributes"`
	TerminatorStrategy string                            `pulumi:"terminatorStrategy"`
}

func migrateServiceV0(_ context.Context, v0 ServiceStateV0) (infer.MigrationResult[ServiceState], error) {
	return infer.MigrationResult[ServiceState]{
		Result: &ServiceState{
			ServiceArgs:        v0.ServiceArgs,
			BaseStateEntity:    v0.BaseStateEntityV0.toV1(),
			Config:             v0.Config,
			Configs:            v0.Configs,
			EncryptionRequired: v0.EncryptionRequired,
			Name:               v0.Name,
			Permissions:        v0.Permissions,
			PostureQueries:     v0.PostureQueries,
			RoleAttributes:     v0.RoleAttributes,
			TerminatorStrategy: v0.TerminatorStrategy,
		},
	}, nil
}

// ─── EdgeRouter V0 ─────────────────────────────────────────────

type EdgeRouterStateV0 struct {
	BaseStateEntityV0
	CommonEdgeRouterProperties
	CertPem               *string `pulumi:"certPem,optional"`
	EnrollmentCreatedAt   *string `pulumi:"enrollmentCreatedAt,optional"`
	EnrollmentExpiresAt   *string `pulumi:"enrollmentExpiresAt,optional"`
	EnrollmentJWT         *string `pulumi:"enrollmentJwt,optional" provider:"secret,output"`
	EnrollmentToken       *string `pulumi:"enrollmentToken,optional" provider:"secret,output"`
	Fingerprint           string  `pulumi:"fingerprint,optional"`
	IsTunnelerEnabled     bool    `pulumi:"isTunnelerEnabled"`
	IsVerified            bool    `pulumi:"isVerified"`
	RoleAttributes        rest_model.Attributes `pulumi:"roleAttributes"`
	UnverifiedCertPem     *string `pulumi:"unverifiedCertPem,optional"`
	UnverifiedFingerprint *string `pulumi:"unverifiedFingerprint,optional"`
	VersionInfo VersionInfo
}

func migrateEdgeRouterV0(_ context.Context, v0 EdgeRouterStateV0) (infer.MigrationResult[EdgeRouterState], error) {
	return infer.MigrationResult[EdgeRouterState]{
		Result: &EdgeRouterState{
			BaseStateEntity:            v0.BaseStateEntityV0.toV1(),
			CommonEdgeRouterProperties: v0.CommonEdgeRouterProperties,
			CertPem:               v0.CertPem,
			EnrollmentCreatedAt:   v0.EnrollmentCreatedAt,
			EnrollmentExpiresAt:   v0.EnrollmentExpiresAt,
			EnrollmentJWT:         v0.EnrollmentJWT,
			EnrollmentToken:       v0.EnrollmentToken,
			Fingerprint:           v0.Fingerprint,
			IsTunnelerEnabled:     v0.IsTunnelerEnabled,
			IsVerified:            v0.IsVerified,
			RoleAttributes:        v0.RoleAttributes,
			UnverifiedCertPem:     v0.UnverifiedCertPem,
			UnverifiedFingerprint: v0.UnverifiedFingerprint,
			VersionInfo:           v0.VersionInfo,
		},
	}, nil
}

// ─── EdgeRouterPolicy V0 ───────────────────────────────────────

type EdgeRouterPolicyStateV0 struct {
	EdgeRouterPolicyArgs
	BaseStateEntityV0
	EdgeRouterRoles        rest_model.Roles `pulumi:"edgeRouterRoles"`
	EdgeRouterRolesDisplay NamedRoles       `pulumi:"edgeRouterRolesDisplay"`
	IdentityRoles          rest_model.Roles `pulumi:"identityRoles"`
	IdentityRolesDisplay   NamedRoles       `pulumi:"identityRolesDisplay"`
	Semantic               rest_model.Semantic `pulumi:"semantic"`
}

func migrateEdgeRouterPolicyV0(_ context.Context, v0 EdgeRouterPolicyStateV0) (infer.MigrationResult[EdgeRouterPolicyState], error) {
	return infer.MigrationResult[EdgeRouterPolicyState]{
		Result: &EdgeRouterPolicyState{
			EdgeRouterPolicyArgs:   v0.EdgeRouterPolicyArgs,
			BaseStateEntity:        v0.BaseStateEntityV0.toV1(),
			EdgeRouterRoles:        v0.EdgeRouterRoles,
			EdgeRouterRolesDisplay: v0.EdgeRouterRolesDisplay,
			IdentityRoles:          v0.IdentityRoles,
			IdentityRolesDisplay:   v0.IdentityRolesDisplay,
			Semantic:               v0.Semantic,
		},
	}, nil
}

// ─── ServiceEdgeRouterPolicy V0 ────────────────────────────────

type ServiceEdgeRouterPolicyStateV0 struct {
	ServiceEdgeRouterPolicyArgs
	BaseStateEntityV0
	EdgeRouterRoles        rest_model.Roles `pulumi:"edgeRouterRoles"`
	EdgeRouterRolesDisplay NamedRoles       `pulumi:"edgeRouterRolesDisplay"`
	ServiceRoles           rest_model.Roles `pulumi:"serviceRoles"`
	ServiceRolesDisplay    NamedRoles       `pulumi:"serviceRolesDisplay"`
	Semantic               rest_model.Semantic `pulumi:"semantic"`
}

func migrateServiceEdgeRouterPolicyV0(_ context.Context, v0 ServiceEdgeRouterPolicyStateV0) (infer.MigrationResult[ServiceEdgeRouterPolicyState], error) {
	return infer.MigrationResult[ServiceEdgeRouterPolicyState]{
		Result: &ServiceEdgeRouterPolicyState{
			ServiceEdgeRouterPolicyArgs: v0.ServiceEdgeRouterPolicyArgs,
			BaseStateEntity:             v0.BaseStateEntityV0.toV1(),
			EdgeRouterRoles:             v0.EdgeRouterRoles,
			EdgeRouterRolesDisplay:      v0.EdgeRouterRolesDisplay,
			ServiceRoles:                v0.ServiceRoles,
			ServiceRolesDisplay:         v0.ServiceRolesDisplay,
			Semantic:                    v0.Semantic,
		},
	}, nil
}

// ─── ConfigObj V0 ──────────────────────────────────────────────

type ConfigStateV0 struct {
	ConfigArgs
	BaseStateEntityV0
	ConfigType   EntityRefV0 `pulumi:"configType"`
	ConfigTypeID string      `pulumi:"configTypeId"`
	Data         interface{} `pulumi:"data"`
}

func migrateConfigV0(_ context.Context, v0 ConfigStateV0) (infer.MigrationResult[ConfigState], error) {
	return infer.MigrationResult[ConfigState]{
		Result: &ConfigState{
			ConfigArgs:      v0.ConfigArgs,
			BaseStateEntity: v0.BaseStateEntityV0.toV1(),
			ConfigType:      v0.ConfigType.toV1(),
			ConfigTypeID:    v0.ConfigTypeID,
			Data:            v0.Data,
		},
	}, nil
}

// ─── Identity V0 ───────────────────────────────────────────────

type IdentityStateV0 struct {
	IdentityArgs
	BaseStateEntityV0
	AppData                    Tags                             `pulumi:"appData,optional"`
	AuthPolicy                 EntityRefV0                      `pulumi:"authPolicy"`
	AuthPolicyID               string                           `pulumi:"authPolicyId"`
	Authenticators             rest_model.IdentityAuthenticators `pulumi:"authenticators"`
	DefaultHostingCost         rest_model.TerminatorCost        `pulumi:"defaultHostingCost"`
	DefaultHostingPrecedence   rest_model.TerminatorPrecedence  `pulumi:"defaultHostingPrecedence,optional"`
	Disabled                   bool                             `pulumi:"disabled"`
	DisabledAt                 *string                          `pulumi:"disabledAt,optional"`
	DisabledUntil              *string                          `pulumi:"disabledUntil,optional"`
	Enrollment                 IdentityEnrollmentsV0            `pulumi:"enrollment" provider:"secret,output"`
	EnvInfo                    EnvInfo                          `pulumi:"envInfo"`
	ExternalID                 *string                          `pulumi:"externalId,optional"`
	HasAPISession              bool                             `pulumi:"hasApiSession"`
	HasEdgeRouterConnection    bool                             `pulumi:"hasEdgeRouterConnection"`
	IsAdmin                    bool                             `pulumi:"isAdmin"`
	IsDefaultAdmin             bool                             `pulumi:"isDefaultAdmin"`
	IsMfaEnabled               bool                             `pulumi:"isMfaEnabled"`
	Name                       string                           `pulumi:"name"`
	RoleAttributes             []string                         `pulumi:"roleAttributes"`
	SdkInfo                    SdkInfo                          `pulumi:"sdkInfo"`
	ServiceHostingCosts        rest_model.TerminatorCostMap     `pulumi:"serviceHostingCosts"`
	ServiceHostingPrecedences  rest_model.TerminatorPrecedenceMap `pulumi:"serviceHostingPrecedences"`
	Type                       EntityRefV0                      `pulumi:"type"`
	TypeID                     string                           `pulumi:"typeId"`
}

func migrateIdentityV0(_ context.Context, v0 IdentityStateV0) (infer.MigrationResult[IdentityState], error) {
	return infer.MigrationResult[IdentityState]{
		Result: &IdentityState{
			IdentityArgs:               v0.IdentityArgs,
			BaseStateEntity:            v0.BaseStateEntityV0.toV1(),
			AppData:                    v0.AppData,
			AuthPolicy:                 v0.AuthPolicy.toV1(),
			AuthPolicyID:               v0.AuthPolicyID,
			Authenticators:             v0.Authenticators,
			DefaultHostingCost:         v0.DefaultHostingCost,
			DefaultHostingPrecedence:   v0.DefaultHostingPrecedence,
			Disabled:                   v0.Disabled,
			DisabledAt:                 v0.DisabledAt,
			DisabledUntil:              v0.DisabledUntil,
			Enrollment:                 v0.Enrollment.toV1(),
			EnvInfo:                    v0.EnvInfo,
			ExternalID:                 v0.ExternalID,
			HasAPISession:              v0.HasAPISession,
			HasEdgeRouterConnection:    v0.HasEdgeRouterConnection,
			IsAdmin:                    v0.IsAdmin,
			IsDefaultAdmin:             v0.IsDefaultAdmin,
			IsMfaEnabled:               v0.IsMfaEnabled,
			Name:                       v0.Name,
			RoleAttributes:             v0.RoleAttributes,
			SdkInfo:                    v0.SdkInfo,
			ServiceHostingCosts:        v0.ServiceHostingCosts,
			ServiceHostingPrecedences:  v0.ServiceHostingPrecedences,
			Type:                       v0.Type.toV1(),
			TypeID:                     v0.TypeID,
		},
	}, nil
}
