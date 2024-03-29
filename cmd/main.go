// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	ghttp "github.com/nathonNot/weaver-ecs/component/http"
	"github.com/nathonNot/weaver-ecs/component/search"
	"net/http"
)

func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

// app is the main component of our application.
type app struct {
	weaver.Implements[weaver.Main]
	searcher weaver.Ref[search.Searcher]
	hello    weaver.Listener
}

// run implements the application main.
func run(ctx context.Context, a *app) error {
	emojis, err := a.searcher.Get().Search(ctx, "pig")
	if err != nil {
		return err
	}
	fmt.Printf("hello listener available on %v\n", a.hello)
	fmt.Println(emojis)
	return http.Serve(a.hello, ghttp.GetHandler())
}
