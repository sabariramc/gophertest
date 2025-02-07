module gopertest

go 1.23

replace (
	gitlab.com/engineering/products/api_security/go-common/api => gitlab.com/engineering/products/api_security/go-common.git/api v0.0.0-20241030070104-54ede13f9886
	gitlab.com/engineering/products/api_security/go-common/cfgpublisher => gitlab.com/engineering/products/api_security/go-common.git/cfgpublisher v0.0.0-20240229174346-1b111c8e03df
	gitlab.com/engineering/products/api_security/go-common/configpublisher => gitlab.com/engineering/products/api_security/go-common.git/configpublisher v1.0.0
	gitlab.com/engineering/products/api_security/go-common/credmanager => gitlab.com/engineering/products/api_security/go-common.git/credmanager v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/datastore => gitlab.com/engineering/products/api_security/go-common.git/datastore v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/datastore/redisclient => gitlab.com/engineering/products/api_security/go-common.git/datastore/redisclient v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/deployment => gitlab.com/engineering/products/api_security/go-common.git/deployment v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/endpoints/definitions => gitlab.com/engineering/products/api_security/go-common.git/endpoints/definitions v0.0.0-20241126063812-bdc18ebfe243
	gitlab.com/engineering/products/api_security/go-common/endpoints/regex => gitlab.com/engineering/products/api_security/go-common.git/endpoints/regex v0.0.0-20241007051307-0993f64382cf
	gitlab.com/engineering/products/api_security/go-common/environment => gitlab.com/engineering/products/api_security/go-common.git/environment v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/eventdefinitions => gitlab.com/engineering/products/api_security/go-common.git/eventdefinitions v0.0.0-20240923065756-8911abca004b
	gitlab.com/engineering/products/api_security/go-common/featureactivation => gitlab.com/engineering/products/api_security/go-common.git/featureactivation v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/imlock => gitlab.com/engineering/products/api_security/go-common.git/imlock v0.0.0-20240325172312-53f7cd465d4b
	gitlab.com/engineering/products/api_security/go-common/log => gitlab.com/engineering/products/api_security/go-common.git/log v1.0.1
	gitlab.com/engineering/products/api_security/go-common/logger => gitlab.com/engineering/products/api_security/go-common.git/logger v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/metrics => gitlab.com/engineering/products/api_security/go-common.git/metrics v0.0.0-20250203084918-580e6c823ad0
	gitlab.com/engineering/products/api_security/go-common/metricscollector => gitlab.com/engineering/products/api_security/go-common.git/metricscollector v0.0.0-20241010170342-e42e8cae2b35
	gitlab.com/engineering/products/api_security/go-common/metricsv2 => gitlab.com/engineering/products/api_security/go-common.git/metricsv2 v0.0.0-20240919145913-ba84d8848f17
	gitlab.com/engineering/products/api_security/go-common/policy/definitions => gitlab.com/engineering/products/api_security/go-common.git/policy/definitions v0.0.0-20241216135115-bb0a3ec08323
	gitlab.com/engineering/products/api_security/go-common/retryablehttp => gitlab.com/engineering/products/api_security/go-common.git/retryablehttp v1.0.1
	gitlab.com/engineering/products/api_security/go-common/serviceauthentication => gitlab.com/engineering/products/api_security/go-common.git/serviceauthentication v0.0.0-20241104052244-8649b4b14e33
	gitlab.com/engineering/products/api_security/go-common/streams => gitlab.com/engineering/products/api_security/go-common.git/streams v0.0.0-20241010170342-e42e8cae2b35
	gitlab.com/engineering/products/api_security/go-common/tester => gitlab.com/engineering/products/api_security/go-common.git/tester v0.0.0-20240916042410-0723add98983
	gitlab.com/engineering/products/api_security/go-common/vault => gitlab.com/engineering/products/api_security/go-common.git/vault v0.0.0-20240819102856-8d2a6de04c60
	gitlab.com/engineering/products/api_security/go-common/worker => gitlab.com/engineering/products/api_security/go-common.git/worker v1.1.1
)

require (
	github.com/google/uuid v1.6.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/redis/go-redis/v9 v9.7.0
	gitlab.com/engineering/products/api_security/go-common/api v0.0.0-00010101000000-000000000000
	gitlab.com/engineering/products/api_security/go-common/environment v0.0.0-00010101000000-000000000000
	gitlab.com/engineering/products/api_security/go-common/log v0.0.0-00010101000000-000000000000
	gitlab.com/engineering/products/api_security/go-common/logger v0.0.0-00010101000000-000000000000
	gitlab.com/engineering/products/api_security/go-common/metrics v0.0.0-00010101000000-000000000000
	gitlab.com/engineering/products/api_security/go-common/tester v0.0.0-00010101000000-000000000000
	gotest.tools v2.2.0+incompatible
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/MicahParks/keyfunc v1.9.0 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/containerd/continuity v0.0.0-20190827140505-75bee3e2ccb6 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/docker/cli v20.10.11+incompatible // indirect
	github.com/docker/docker v20.10.7+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.0.2 // indirect
	github.com/ory/dockertest/v3 v3.8.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.54.0 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/romana/rlog v0.0.0-20220412051723-c08f605858a9 // indirect
	github.com/segmentio/kafka-go v0.4.31 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	gitlab.com/engineering/products/api_security/go-common/credmanager v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/engineering/products/api_security/go-common/deployment v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/engineering/products/api_security/go-common/serviceauthentication v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/engineering/products/api_security/go-common/vault v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
