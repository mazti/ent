package main

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/entc/integration/enthook"
	_ "github.com/facebookincubator/ent/entc/integration/enthook/hooks"
	_ "github.com/facebookincubator/ent/entc/integration/enthook/runtime"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	client, err := enthook.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}
	client.Card.WithHooks(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			fmt.Printf("Op: %s\tType: %s\tConcreteType: %T\n", m.Op(), m.Type(), m)
			if ns, ok := m.(interface{ SetName(string) }); ok {
				ns.SetName("hook name")
			}
			return next.Mutate(ctx, m)
		})
	})
	u := client.Card.Create().SetNumber("A").SaveX(ctx)
	u.Update().SetName("Boring2").SaveX(ctx)
	client.Card.Update().SetName("foo").SaveX(ctx)
	client.Card.DeleteOneID(u.ID).ExecX(ctx)
	client.Card.Delete().ExecX(ctx)
}
