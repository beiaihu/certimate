package domain

import (
	"time"

	"github.com/usual2970/certimate/internal/pkg/utils/maps"
)

const CollectionNameWorkflow = "workflow"

type Workflow struct {
	Meta
	Name          string                `json:"name" db:"name"`
	Description   string                `json:"description" db:"description"`
	Trigger       WorkflowTriggerType   `json:"trigger" db:"trigger"`
	TriggerCron   string                `json:"triggerCron" db:"triggerCron"`
	Enabled       bool                  `json:"enabled" db:"enabled"`
	Content       *WorkflowNode         `json:"content" db:"content"`
	Draft         *WorkflowNode         `json:"draft" db:"draft"`
	HasDraft      bool                  `json:"hasDraft" db:"hasDraft"`
	LastRunId     string                `json:"lastRunId" db:"lastRunId"`
	LastRunStatus WorkflowRunStatusType `json:"lastRunStatus" db:"lastRunStatus"`
	LastRunTime   time.Time             `json:"lastRunTime" db:"lastRunTime"`
}

type WorkflowNodeType string

const (
	WorkflowNodeTypeStart               = WorkflowNodeType("start")
	WorkflowNodeTypeEnd                 = WorkflowNodeType("end")
	WorkflowNodeTypeApply               = WorkflowNodeType("apply")
	WorkflowNodeTypeDeploy              = WorkflowNodeType("deploy")
	WorkflowNodeTypeNotify              = WorkflowNodeType("notify")
	WorkflowNodeTypeBranch              = WorkflowNodeType("branch")
	WorkflowNodeTypeCondition           = WorkflowNodeType("condition")
	WorkflowNodeTypeExecuteResultBranch = WorkflowNodeType("execute_result_branch")
	WorkflowNodeTypeExecuteSuccess      = WorkflowNodeType("execute_success")
	WorkflowNodeTypeExecuteFailure      = WorkflowNodeType("execute_failure")
)

type WorkflowTriggerType string

const (
	WorkflowTriggerTypeAuto   = WorkflowTriggerType("auto")
	WorkflowTriggerTypeManual = WorkflowTriggerType("manual")
)

type WorkflowNode struct {
	Id   string           `json:"id"`
	Type WorkflowNodeType `json:"type"`
	Name string           `json:"name"`

	Config  map[string]any   `json:"config"`
	Inputs  []WorkflowNodeIO `json:"inputs"`
	Outputs []WorkflowNodeIO `json:"outputs"`

	Next     *WorkflowNode  `json:"next"`
	Branches []WorkflowNode `json:"branches"`

	Validated bool `json:"validated"`
}

type WorkflowNodeConfigForApply struct {
	Domains               string         `json:"domains"`               // 域名列表，以半角逗号分隔
	ContactEmail          string         `json:"contactEmail"`          // 联系邮箱
	Provider              string         `json:"provider"`              // DNS 提供商
	ProviderAccessId      string         `json:"providerAccessId"`      // DNS 提供商授权记录 ID
	ProviderConfig        map[string]any `json:"providerConfig"`        // DNS 提供商额外配置
	KeyAlgorithm          string         `json:"keyAlgorithm"`          // 密钥算法
	Nameservers           string         `json:"nameservers"`           // DNS 服务器列表，以半角逗号分隔
	DnsPropagationTimeout int32          `json:"dnsPropagationTimeout"` // DNS 传播超时时间（默认取决于提供商）
	DnsTTL                int32          `json:"dnsTTL"`                // DNS TTL（默认取决于提供商）
	DisableFollowCNAME    bool           `json:"disableFollowCNAME"`    // 是否禁用 CNAME 跟随
	DisableARI            bool           `json:"disableARI"`            // 是否禁用 ARI
	SkipBeforeExpiryDays  int32          `json:"skipBeforeExpiryDays"`  // 证书到期前多少天前跳过续期（默认值：30）
}

type WorkflowNodeConfigForDeploy struct {
	Certificate         string         `json:"certificate"`         // 前序节点输出的证书，形如“${NodeId}#certificate”
	Provider            string         `json:"provider"`            // 主机提供商
	ProviderAccessId    string         `json:"providerAccessId"`    // 主机提供商授权记录 ID
	ProviderConfig      map[string]any `json:"providerConfig"`      // 主机提供商额外配置
	SkipOnLastSucceeded bool           `json:"skipOnLastSucceeded"` // 上次部署成功时是否跳过
}

type WorkflowNodeConfigForNotify struct {
	Channel string `json:"channel"` // 通知渠道
	Subject string `json:"subject"` // 通知主题
	Message string `json:"message"` // 通知内容
}

func (n *WorkflowNode) getConfigValueAsString(key string) string {
	return maps.GetValueAsString(n.Config, key)
}

func (n *WorkflowNode) getConfigValueAsBool(key string) bool {
	return maps.GetValueAsBool(n.Config, key)
}

func (n *WorkflowNode) getConfigValueAsInt32(key string) int32 {
	return maps.GetValueAsInt32(n.Config, key)
}

func (n *WorkflowNode) getConfigValueAsMap(key string) map[string]any {
	if val, ok := n.Config[key]; ok {
		if result, ok := val.(map[string]any); ok {
			return result
		}
	}

	return make(map[string]any)
}

func (n *WorkflowNode) GetConfigForApply() WorkflowNodeConfigForApply {
	skipBeforeExpiryDays := n.getConfigValueAsInt32("skipBeforeExpiryDays")
	if skipBeforeExpiryDays == 0 {
		skipBeforeExpiryDays = 30
	}

	return WorkflowNodeConfigForApply{
		Domains:               n.getConfigValueAsString("domains"),
		ContactEmail:          n.getConfigValueAsString("contactEmail"),
		Provider:              n.getConfigValueAsString("provider"),
		ProviderAccessId:      n.getConfigValueAsString("providerAccessId"),
		ProviderConfig:        n.getConfigValueAsMap("providerConfig"),
		KeyAlgorithm:          n.getConfigValueAsString("keyAlgorithm"),
		Nameservers:           n.getConfigValueAsString("nameservers"),
		DnsPropagationTimeout: n.getConfigValueAsInt32("dnsPropagationTimeout"),
		DnsTTL:                n.getConfigValueAsInt32("dnsTTL"),
		DisableFollowCNAME:    n.getConfigValueAsBool("disableFollowCNAME"),
		DisableARI:            n.getConfigValueAsBool("disableARI"),
		SkipBeforeExpiryDays:  skipBeforeExpiryDays,
	}
}

func (n *WorkflowNode) GetConfigForDeploy() WorkflowNodeConfigForDeploy {
	return WorkflowNodeConfigForDeploy{
		Certificate:         n.getConfigValueAsString("certificate"),
		Provider:            n.getConfigValueAsString("provider"),
		ProviderAccessId:    n.getConfigValueAsString("providerAccessId"),
		ProviderConfig:      n.getConfigValueAsMap("providerConfig"),
		SkipOnLastSucceeded: n.getConfigValueAsBool("skipOnLastSucceeded"),
	}
}

func (n *WorkflowNode) GetConfigForNotify() WorkflowNodeConfigForNotify {
	return WorkflowNodeConfigForNotify{
		Channel: n.getConfigValueAsString("channel"),
		Subject: n.getConfigValueAsString("subject"),
		Message: n.getConfigValueAsString("message"),
	}
}

type WorkflowNodeIO struct {
	Label         string                      `json:"label"`
	Name          string                      `json:"name"`
	Type          string                      `json:"type"`
	Required      bool                        `json:"required"`
	Value         any                         `json:"value"`
	ValueSelector WorkflowNodeIOValueSelector `json:"valueSelector"`
}

type WorkflowNodeIOValueSelector struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

const WorkflowNodeIONameCertificate string = "certificate"
