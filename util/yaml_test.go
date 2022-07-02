package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

type ServiceConfig struct {
	EnvironmentStruct `yaml:"environment"`
}

type EnvironmentStruct struct {
	M1
	M2 `yaml:",inline"`
}

type M1 []string
type M2 map[string]string

func getSC(buf []byte) (*ServiceConfig, error) {
	var v ServiceConfig
	err := yaml.Unmarshal(buf, v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func TestMarshalGeneric(t *testing.T) {
	yamlStr := `
    environment:
      LDAP_READONLY_USER: "{{ common.ldap_readonly_user }}"
      LDAP_ORGANISATION: "{{ common.ldap_organisation }}"
      LDAP_DOMAIN: "{{ common.ldap_domain }}"
      LDAP_ADMIN_PASSWORD: "{{ common.ldap_password }}"
      LDAP_NOFILE: "{{common.ldap_nofile}}"
`
	sc, err := getSC([]byte(yamlStr))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(sc)

}
