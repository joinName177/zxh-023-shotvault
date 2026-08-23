package main

import (
	"context"
	"fmt"
	"github.com/joinName177/zxh-023-shotvault/internal/vault"
	"os"
)

func main() {
	path := "shotvault.json"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	repo, err := vault.OpenFileRepository(path)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cat := vault.NewCatalog(repo)
	if _, err = cat.Create(ctx, "demo", "FrameLoom"); err != nil && err != vault.ErrWorkspaceNotFound {
		fmt.Println(err)
	}
	fmt.Println("shot archive ready", path)
}
