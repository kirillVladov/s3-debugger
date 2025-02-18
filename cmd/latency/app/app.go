package app

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type App struct {
	config      *Config
	minioClient *minio.Client
}

func NewApp() (app App, err error) {
	app.config, err = initConfig()

	if err != nil {
		return app, fmt.Errorf(err.Error())
	}

	app.minioClient, err = minio.New(app.config.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			app.config.AccessKey,
			app.config.SecretKey,
			"",
		),
		Region: app.config.Region,
		Secure: true,
	})

	if err != nil {
		return app, fmt.Errorf(err.Error())
	}

	return app, nil
}

func (a *App) Run() {
	ctx := context.Background()

	listObjects := a.minioClient.ListObjects(ctx, a.config.BucketName, minio.ListObjectsOptions{
		MaxKeys: a.config.GettingCount,
	})

	for i := 0; i < a.config.GettingCount; i++ {
		object := <-listObjects
		startTime := time.Now()

		_, err := a.minioClient.GetObject(
			ctx,
			a.config.BucketName,
			object.Key,
			minio.GetObjectOptions{},
		)

		if err != nil {
			fmt.Printf("Time: %v, Key: %s, Error: %s \n", time.Since(startTime), object.Key, err.Error())
			continue
		}

		fmt.Printf("Time: %v, Key: %s, Size: %v bytes \n", time.Since(startTime), object.Key, object.Size)
	}

	a.Shutdown(ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	ctx.Done()
}
