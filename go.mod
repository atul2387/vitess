module vitess.io/vitess

go 1.16

require (
	cloud.google.com/go/storage v1.10.0
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20210330150358-dbd898e17899
	github.com/Azure/azure-pipeline-go v0.2.2
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest v0.10.0 // indirect
	github.com/DataDog/datadog-go v2.2.0+incompatible
	github.com/GeertJohan/go.rice v1.0.0
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/aquarapid/vaultlib v0.5.1
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/aws/aws-sdk-go v1.34.2
	github.com/buger/jsonparser v0.0.0-20200322175846-f7e751efca13
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/corpix/uarand v0.1.1 // indirect
	github.com/dave/jennifer v1.4.1
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab
	github.com/go-sql-driver/mysql v1.5.1-0.20210202043019-fe2230a8b20c
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.1
	github.com/google/go-cmp v0.5.5
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510
	github.com/google/uuid v1.1.2
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/consul/api v1.5.0
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/serf v0.9.2 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/icrowley/fake v0.0.0-20180203215853-4178557ae428
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/klauspost/compress v1.4.1 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/klauspost/pgzip v1.2.4
	github.com/krishicks/yaml-patch v0.0.10
	github.com/magiconair/properties v1.8.1
	github.com/martini-contrib/auth v0.0.0-20150219114609-fa62c19b7ae8
	github.com/martini-contrib/gzip v0.0.0-20151124214156-6c035326b43f
	github.com/martini-contrib/render v0.0.0-20150707142108-ec18f8345a11
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/minio/minio-go v0.0.0-20190131015406-c8a261de75c1
	github.com/mitchellh/go-testing-interface v1.14.0 // indirect
	github.com/mitchellh/mapstructure v1.2.3 // indirect
	github.com/montanaflynn/stats v0.6.3
	github.com/olekukonko/tablewriter v0.0.5-0.20200416053754-163badb3bac6
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20180928155321-4b5a12d3ff02
	github.com/opentracing/opentracing-go v1.1.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pborman/uuid v1.2.0
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pires/go-proxyproto v0.0.0-20191211124218-517ecdf5bb2b
	github.com/pkg/errors v0.9.1
	github.com/planetscale/pargzip v0.0.0-20201116224723-90c7fc03ea8a
	github.com/planetscale/tengo v0.9.6-ps.v4
	github.com/planetscale/vtprotobuf v0.0.0-20210521163914-5a02622d1e2a
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.29.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/samuel/go-zookeeper v0.0.0-20200724154423-2164a8ac840e
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sjmudd/stopwatch v0.0.0-20170613150411-f380bf8a9be1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spyzhov/ajson v0.4.2
	github.com/stretchr/testify v1.7.0
	github.com/tchap/go-patricia v0.0.0-20160729071656-dd168db6051b
	github.com/tebeka/selenium v0.9.9
	github.com/tinylib/msgp v1.1.1 // indirect
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	github.com/z-division/go-zookeeper v0.0.0-20190128072838-6d7457066b9b
	go.etcd.io/etcd/api/v3 v3.5.0
	go.etcd.io/etcd/client/pkg/v3 v3.5.0
	go.etcd.io/etcd/client/v3 v3.5.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools v0.1.4
	google.golang.org/api v0.45.0
	google.golang.org/genproto v0.0.0-20210701191553-46259e63a0a9 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/grpc/examples v0.0.0-20210430044426-28078834f35b
	google.golang.org/protobuf v1.27.1
	gopkg.in/DataDog/dd-trace-go.v1 v1.17.0
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/ldap.v2 v2.5.0
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
	honnef.co/go/tools v0.0.1-2020.1.4
	k8s.io/apiextensions-apiserver v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v0.17.3
	k8s.io/code-generator v0.17.3
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/skeema/tengo => github.com/planetscale/tengo v0.9.6-ps.v1

replace google.golang.org/grpc => google.golang.org/grpc v1.37.0
