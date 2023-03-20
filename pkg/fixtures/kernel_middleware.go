package fixtures

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
)

func KernelMiddlewareLoader(factory FixtureBuilderFactory) func(context context.Context, config cfg.GosoConf, logger log.GosoLogger) error {
	return func(ctx context.Context, config cfg.GosoConf, logger log.GosoLogger) error {
		loader := NewFixtureLoader(ctx, config, logger)
		builder := factory(ctx)

		if err := loader.Load(ctx, builder.Fixtures()); err != nil {
			return fmt.Errorf("can not load fixtureSets: %w", err)
		}

		return nil
	}
}
