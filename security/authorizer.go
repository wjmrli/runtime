// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package security

import (
	"net/http"

	"github.com/wjmrli/runtime"
)

// Authorized provides a default implementation of the Authorizer interface where all
// requests are authorized (successful)
func Authorized() runtime.Authorizer {
	return runtime.AuthorizerFunc(func(_ *http.Request, _ interface{}) error { return nil })
}
