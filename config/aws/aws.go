package aws

import (
	"fmt"
	"log"

	"github.com/NYTimes/gizmo/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/bradfitz/gomemcache/memcache"
)

const (
	// RegionUSEast1 is a helper constant for AWS configs.
	RegionUSEast1 = "us-east-1"
	// RegionUSWest is a helper constant for AWS configs.
	RegionUSWest = "us-west-1"
)

type (
	// Config holds common AWS credentials and keys.
	Config struct {
		SecretKey string `envconfig:"AWS_SECRET_KEY"`
		AccessKey string `envconfig:"AWS_ACCESS_KEY"`

		Region string `envconfig:"AWS_REGION"`
	}

	// S3 holds the info required to work with Amazon S3.
	S3 struct {
		Config
		Bucket string `envconfig:"AWS_S3_BUCKET_NAME"`
	}

	// DynamoDB holds some basic info required to work with Amazon DynamoDB.
	DynamoDB struct {
		Config
		TableName string `envconfig:"AWS_DYNAMODB_TABLE_NAME"`
	}

	// ElastiCache holds the basic info required to work with
	// Amazon ElastiCache.
	ElastiCache struct {
		Config
		ClusterID string `envconfig:"AWS_ELASTICACHE_CLUSTER_ID"`
	}
)

// MustClient will use the cache cluster ID to describe
// the cache cluster and instantiate a memcache.Client
// with the cache nodes returned from AWS.
func (e *ElastiCache) MustClient() *memcache.Client {
	var creds *credentials.Credentials
	if e.AccessKey != "" {
		creds = credentials.NewStaticCredentials(e.AccessKey, e.SecretKey, "")
	} else {
		creds = credentials.NewEnvCredentials()
	}

	ecclient := elasticache.New(session.New(&aws.Config{
		Credentials: creds,
		Region:      &e.Region,
	}))

	resp, err := ecclient.DescribeCacheClusters(&elasticache.DescribeCacheClustersInput{
		CacheClusterId:    &e.ClusterID,
		ShowCacheNodeInfo: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("unable to describe cache cluster: %s", err)
	}

	var nodes []string
	for _, cluster := range resp.CacheClusters {
		for _, cnode := range cluster.CacheNodes {
			addr := fmt.Sprintf("%s:%d", *cnode.Endpoint.Address, *cnode.Endpoint.Port)
			nodes = append(nodes, addr)
		}
	}

	return memcache.New(nodes...)
}

// LoadConfigFromEnv will attempt to load the Config struct
// from environment variables.
func LoadConfigFromEnv() Config {
	var aws Config
	config.LoadEnvConfig(&aws)
	return aws
}

// LoadDynamoDBFromEnv will attempt to load the DynamoDB struct
// from environment variables. If not populated, nil
// is returned.
func LoadDynamoDBFromEnv() DynamoDB {
	var ddb DynamoDB
	config.LoadEnvConfig(&ddb)
	return ddb
}

// LoadS3FromEnv will attempt to load the S3 struct
// from environment variables.
func LoadS3FromEnv() S3 {
	var s3 S3
	config.LoadEnvConfig(&s3)
	return s3
}

// LoadElastiCacheFromEnv will attempt to load the ElasiCache struct
// from environment variables.
func LoadElastiCacheFromEnv() ElastiCache {
	var el ElastiCache
	config.LoadEnvConfig(&el)
	return el
}
