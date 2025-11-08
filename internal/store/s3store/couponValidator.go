package s3store

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type CouponValidator struct {
	client *s3.Client
	bucket string
	keys   []string
}

func NewCouponValidator(ctx context.Context, bucket string, keys []string) (*CouponValidator, error) {
	if len(keys) == 0 {
		return nil, errors.New("keys slice cannot be empty")
	}
	region := "ap-southeast-2"
	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}
	opts = append(opts, config.WithCredentialsProvider(aws.AnonymousCredentials{}))

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	return &CouponValidator{client: client, bucket: bucket, keys: keys}, nil
}

type checkResult struct {
	found bool
	err   error
}

func (v *CouponValidator) ValidatePromo(ctx context.Context, code string) error {
	code = strings.TrimSpace(strings.ToUpper(code))
	if len(code) < 8 || len(code) > 10 {
		return errors.New("invalid promo: code length must be between 8 and 10")
	}

	var wg sync.WaitGroup
	results := make(chan checkResult, len(v.keys))

	for _, key := range v.keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			found, err := v.checkCodeInS3File(ctx, k, code)
			results <- checkResult{found: found, err: err}
		}(key)
	}

	wg.Wait()
	close(results)

	count := 0
	var firstS3Error error

	for res := range results {
		if res.err != nil {

			if firstS3Error == nil {
				firstS3Error = fmt.Errorf("file %s: %w", res.err.Error(), res.err) // Provide more context
			}

			continue
		}
		if res.found {
			count++
		}
	}

	if count >= 2 {
		return nil
	}

	if firstS3Error != nil {
		return fmt.Errorf("invalid promo: not found in enough coupon files (encountered S3 error: %w)", firstS3Error)
	}

	return errors.New("invalid promo: not found in enough coupon files")
}

func (v *CouponValidator) checkCodeInS3File(ctx context.Context, key, code string) (bool, error) {
	query := fmt.Sprintf("SELECT s._1 FROM S3Object s WHERE s._1 = '%s'", code)

	input := &s3.SelectObjectContentInput{
		Bucket:         aws.String(v.bucket),
		Key:            aws.String(key),
		ExpressionType: s3types.ExpressionTypeSql,
		Expression:     aws.String(query),
		InputSerialization: &s3types.InputSerialization{
			CompressionType: s3types.CompressionTypeGzip,
			CSV: &s3types.CSVInput{
				FileHeaderInfo:  s3types.FileHeaderInfoNone,
				RecordDelimiter: aws.String("\n"),
			},
		},
		OutputSerialization: &s3types.OutputSerialization{
			CSV: &s3types.CSVOutput{},
		},
	}

	resp, err := v.client.SelectObjectContent(ctx, input)
	if err != nil {
		return false, fmt.Errorf("s3 select failed for key %s: %w", key, err)
	}

	stream := resp.GetStream()
	defer stream.Close()

	var buf bytes.Buffer

	for event := range stream.Events() {
		switch e := event.(type) {
		case *s3types.SelectObjectContentEventStreamMemberRecords:
			buf.Write(e.Value.Payload)
		case *s3types.SelectObjectContentEventStreamMemberEnd:
		case *s3types.SelectObjectContentEventStreamMemberStats:
		case *s3types.SelectObjectContentEventStreamMemberProgress:
		default:
			return false, fmt.Errorf("received unexpected S3 event type: %T", event)
		}
	}

	if err := stream.Err(); err != nil {
		if !errors.Is(err, io.EOF) {
			return false, fmt.Errorf("stream error for key %s: %w", key, err)
		}
	}

	return buf.Len() > 0, nil
}
