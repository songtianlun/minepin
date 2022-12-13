package constvar

type UserType uint32

const (
	DefaultCfgEnvPrefix = "MINEPIN"
	DefaultCfgPath      = "./"
	DefaultCfgName      = "config"
	DefaultCfgType      = "yaml"
	DefaultCfgFile      = DefaultCfgPath + DefaultCfgName + "." + DefaultCfgType

	UserVisitor    UserType = 0
	UserRegistered UserType = 1
	UserVIP        UserType = 2
	UserAdmin      UserType = 10

	DefaultGroupName = "Default"

	CRSBd09  = "BD09"
	CRSWgs84 = "WGS84"
	CRSGcj02 = "GCJ02"

	PingsMapCluster = "cluster"
	PingsMapRoute   = "route"
)

var (
	PinsMapTypes = map[string]string{
		"cluster": "点聚合图",
		"route":   "路线图 (BETA)",
	}
)
