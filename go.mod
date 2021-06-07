module github.com/b-social/terraform-provider-kafkaconnect

require (
	cloud.google.com/go v0.41.0
	contrib.go.opencensus.io/exporter/ocagent v0.4.12
	github.com/Azure/azure-sdk-for-go v31.0.0+incompatible
	github.com/Azure/go-autorest v12.2.0+incompatible
	github.com/Azure/go-ntlmssp v0.0.0-20180810175552-4a21cbd618b4
	github.com/ChrisTrenkamp/goxpath v0.0.0-20160627023518-2ad3b31cf4a2
	github.com/Unknwon/com v0.0.0-20181010210213-41959bdd855f
	github.com/agext/levenshtein v1.2.2
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190704060827-a16769e82bb4
	github.com/aliyun/aliyun-oss-go-sdk v2.0.0+incompatible
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apparentlymart/go-cidr v1.0.0
	github.com/apparentlymart/go-textseg v1.0.0
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878
	github.com/armon/go-radix v1.0.0
	github.com/aws/aws-sdk-go v1.20.15
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d
	github.com/bgentry/speakeasy v0.1.0
	github.com/blang/semver v0.0.0-20190414102917-ba2c2ddd8906
	github.com/census-instrumentation/opencensus-proto v0.2.0
	github.com/chzyer/readline v0.0.0-20160726135117-62c6fe619375
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/coreos/go-semver v0.3.0
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimchansky/utfbom v1.1.0
	github.com/dylanmei/iso8601 v0.1.0
	github.com/fatih/color v1.7.0
	github.com/go-kafka/connect v0.9.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	github.com/google/go-cmp v0.3.0
	github.com/google/go-querystring v1.0.0
	github.com/googleapis/gax-go v1.0.3
	github.com/gophercloud/gophercloud v0.2.0
	github.com/gophercloud/utils v0.0.0-20190527093828-25f1b77b8c03
	github.com/grpc-ecosystem/grpc-gateway v1.9.3
	github.com/hashicorp/aws-sdk-go-base v0.3.0
	github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/errwrap v1.0.0
	github.com/hashicorp/go-azure-helpers v0.4.1
	github.com/hashicorp/go-checkpoint v0.5.0
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-getter v1.3.0
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-immutable-radix v1.1.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/go-retryablehttp v0.5.4
	github.com/hashicorp/go-rootcerts v1.0.1
	github.com/hashicorp/go-safetemp v1.0.0
	github.com/hashicorp/go-slug v0.3.1
	github.com/hashicorp/go-tfe v0.3.19
	github.com/hashicorp/go-uuid v1.0.1
	github.com/hashicorp/go-version v1.2.0
	github.com/hashicorp/golang-lru v0.5.1
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/hcl2 v0.0.0-20190702185634-5b39d9ff3a9a
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/serf v0.8.3
	github.com/hashicorp/terraform v0.12.3
	github.com/hashicorp/terraform-config-inspect v0.0.0-20190628153518-9c24e68f3f10
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/joyent/triton-go v1.7.0
	github.com/json-iterator/go v1.1.6
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/lib/pq v1.1.1
	github.com/lusis/go-artifactory v0.0.0-20180304164534-a47f63f234b2
	github.com/masterzen/simplexml v0.0.0-20190410153822-31eea3082786
	github.com/masterzen/winrm v0.0.0-20190308153735-1d17eaf15943
	github.com/mattn/go-colorable v0.0.9
	github.com/mattn/go-isatty v0.0.8
	github.com/mattn/go-shellwords v1.0.5
	github.com/mitchellh/cli v1.0.0
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
	github.com/mitchellh/copystructure v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-linereader v0.0.0-20190213213312-1b945b3263eb
	github.com/mitchellh/go-testing-interface v1.0.0
	github.com/mitchellh/go-wordwrap v1.0.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/mitchellh/panicwrap v0.0.0-20190228164358-f67bf3f3d291
	github.com/mitchellh/prefixedio v0.0.0-20190213213902-5733675afd51
	github.com/mitchellh/reflectwalk v1.0.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/oklog/run v1.0.0
	github.com/packer-community/winrmcp v0.0.0-20180921211025-c76d91c1e7db
	github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9
	github.com/posener/complete v1.2.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/afero v1.2.2
	github.com/svanharmelen/jsonapi v0.0.0-20170708005851-46d3ced04344
	github.com/terraform-providers/terraform-provider-openstack v1.19.0
	github.com/ulikunitz/xz v0.5.6
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/xanzy/ssh-agent v0.2.1
	github.com/xlab/treeprint v0.0.0-20181112141820-a009c3971eca
	github.com/zclconf/go-cty v1.0.0
	github.com/zclconf/go-cty-yaml v0.1.0
	go.opencensus.io v0.19.3
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190626221950-04f50cda93cb
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/api v0.7.0
	google.golang.org/appengine v1.6.1
	google.golang.org/genproto v0.0.0-20190701230453-710ae3a149df
	google.golang.org/grpc v1.22.0
	gopkg.in/ini.v1 v1.42.0
	gopkg.in/yaml.v2 v2.2.2
)
