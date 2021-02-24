package helpers

import (
	"context"
	"time"
	"fmt"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware wraps the logs for incoming requests
type LoggingMiddleware struct {
	Logger log.Logger
	Next EncryptService
}

// Encrypt logs the encryption requests 
func (mw LoggingMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time){
		mw.Logger.Log(
			"method", "encrypt",
			"key", key,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	} (time.Now())
	
	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}

// Decrypt logs the encyption requests
func (mw LoggingMiddleware) Decrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "decrypt",
			"key", key,
			"message", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	} (time.Now())

	output, err = mw.Next.Decrypt(ctx, key, text)
	return 
}