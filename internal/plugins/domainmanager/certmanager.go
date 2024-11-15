package domainmanager

import (
	"errors"
	"fmt"
	"strings"

	"github.com/duc-cnzj/mars/v5/internal/application"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/util/hash"
)

var (
	name            = "cert-manager_domain_manager"
	maxDomainLength = 64
)

var _ application.DomainManager = (*certManager)(nil)

func init() {
	dr := &certManager{}
	application.RegisterPlugin(dr.Name(), dr)
}

// certManager 因为 lets encrypt 对 subdomain 长度要求为 64，所以需要处理。
type certManager struct {
	nsPrefix       string
	clusterIssuer  string
	wildcardDomain string
	domainSuffix   string

	logger mlog.Logger
}

func (d *certManager) Name() string {
	return name
}

func (d *certManager) Initialize(app application.App, args map[string]any) error {
	d.logger = app.Logger()

	if p, ok := args["ns_prefix"]; ok {
		d.nsPrefix = p.(string)
	}

	if issuer, ok := args["cluster_issuer"]; ok {
		d.clusterIssuer = issuer.(string)
	}

	if wd, ok := args["wildcard_domain"]; ok {
		d.wildcardDomain = wd.(string)
		d.domainSuffix = strings.TrimLeft(d.wildcardDomain, "*.")
	}

	if d.clusterIssuer == "" || d.wildcardDomain == "" {
		return errors.New("cluster_issuer, wildcard_domain required")
	}

	d.logger.Info("[Plugin]: " + d.Name() + " plugin Initialize...")
	return nil
}

func (d *certManager) Destroy() error {
	d.logger.Info("[Plugin]: " + d.Name() + " plugin Destroy...")
	return nil
}

func (d *certManager) GetCertSecretName(projectName string, index int) string {
	return fmt.Sprintf("mars-tls-%s", hash.Hash(fmt.Sprintf("%s-%d", projectName, index)))
}

func (d *certManager) GetClusterIssuer() string {
	return d.clusterIssuer
}

func (d *certManager) GetDomainByIndex(projectName, namespace string, index, preOccupiedLen int) string {
	return Subdomain{
		maxLen:       maxDomainLength - preOccupiedLen,
		projectName:  projectName,
		namespace:    namespace,
		index:        index,
		nsPrefix:     d.nsPrefix,
		domainSuffix: d.domainSuffix,
	}.SubStr()
}

func (d *certManager) GetDomain(projectName, namespace string, preOccupiedLen int) string {
	return Subdomain{
		maxLen:       maxDomainLength - preOccupiedLen,
		projectName:  projectName,
		namespace:    namespace,
		index:        -1,
		nsPrefix:     d.nsPrefix,
		domainSuffix: d.domainSuffix,
	}.SubStr()
}

func (d *certManager) GetCerts() (name, key, crt string) {
	return "", "", ""
}

type Subdomain struct {
	maxLen       int
	projectName  string
	namespace    string
	index        int
	nsPrefix     string
	domainSuffix string
}

func (s Subdomain) SubStr() string {
	if s.maxLen == 0 {
		return s.CompleteSubdomain()
	}

	if len(s.CompleteSubdomain()) <= s.maxLen {
		return s.CompleteSubdomain()
	}

	if len(s.MediumSubdomain()) <= s.maxLen {
		return s.MediumSubdomain()
	}

	return s.SimpleSubdomain()
}

func (s Subdomain) HasIndex() bool {
	return s.index != -1
}

// CompleteSubdomain 获取完整的名称 mars-devops-test-default.test.com
func (s Subdomain) CompleteSubdomain() string {
	if s.HasIndex() {
		return fmt.Sprintf("%s-%s-%d.%s", s.projectName, s.namespace, s.index, s.domainSuffix)
	}

	return fmt.Sprintf("%s-%s.%s", s.projectName, s.namespace, s.domainSuffix)
}

// MediumSubdomain 中等版本, 去掉了 ns "devops-" 前缀
func (s Subdomain) MediumSubdomain() string {
	nname := strings.TrimLeft(s.namespace, s.nsPrefix)
	if s.HasIndex() {
		return fmt.Sprintf("%s-%s-%d.%s", s.projectName, nname, s.index, s.domainSuffix)
	}

	return fmt.Sprintf("%s-%s.%s", s.projectName, nname, s.domainSuffix)
}

// SimpleSubdomain 简单版本
func (s Subdomain) SimpleSubdomain() string {
	leftLen := s.maxLen - len(s.domainSuffix) - 1
	if leftLen <= 0 {
		panic(fmt.Errorf("substr error: max len: %d, left len: %d, domainSuffix: %s, project: %s, ns: %s, index: %d", s.maxLen, leftLen, s.domainSuffix, s.projectName, s.namespace, s.index))
	}
	var str = fmt.Sprintf("%s-%s", s.projectName, s.namespace)
	if s.HasIndex() {
		str = fmt.Sprintf("%s-%s-%d", s.projectName, s.namespace, s.index)
	}
	ss := substr(hash.Hash(str), leftLen)

	return fmt.Sprintf("%s.%s", ss, s.domainSuffix)
}

func substr(s string, length int) string {
	if len(s) < length {
		return s
	}

	return s[0:length]
}
