package constants

const (
	MySQLDSN             = "douyin:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=true"
	RedisAddr            = "127.0.0.1:6379"
	RedisPassword        = ""
	EtcdAddress          = "127.0.0.1:2379"
	MinioEndPoint        = "192.168.2.19:18001"
	MinioAccessKeyID     = "minio"
	MinioAccessSecretKey = "12345678"

	SecretKey   = "tiktok"
	IdentityKey = "user_id"

	VideoFeedCount = 30
	UserNameMaxLen = 32
	PassWordMaxLen = 32
	PasswordMinLen = 5
)

const (
	UserTableName     = "users"
	RelationTableName = "relations"
	MessageTableName  = "messages"
	PublishTableName  = "videos"
	FavoriteTableName = "favorites"
	CommentTableName  = "comments"

	MinioVideoBucketName      = "videos"
	MinioImageBucketName      = "images"
	MinioAvatarBucketName     = "avatars"
	MinioBackgroundBucketName = "backgrounds"
)

const (
	UserServiceName     = "user"
	RelationServiceName = "relation"
	PublishServiceName  = "publish"
	MessageServiceName  = "message"
	FavoriteServiceName = "favorite"
	CommentServiceName  = "comment"
	FeedServiceName     = "feed"
	ApiServiceName      = "api"
)

const (
	TestAva        = "test.jpg"
	TestBackground = "test.jpg"
)

const (
	RabbitMqURI = "amqp://%s:%s@%s:%d/"
	MQUser      = "douyin"
	MQPassword  = "123456"
	MQHost      = "127.0.0.1"
	MQPort      = 5672
)
