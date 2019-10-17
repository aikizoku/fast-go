package main

import (
	"context"

	"github.com/aikizoku/rabbitgo/command/common"
	"github.com/aikizoku/rabbitgo/command/seed/content"
)

func main() {
	ctx := context.Background()

	// env.jsonの読み込み
	env := common.LoadEnvFile(common.Production)

	// Inject
	fCli := common.NewFirestoreClient(env)

	u := &content.Sample{
		FCli: fCli,
	}

	// 実行
	u.Generate(ctx)
}