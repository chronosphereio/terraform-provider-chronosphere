// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configv1

//go:generate swagger generate client --spec=./swagger.json --target=. --additional-initialism=SLO
//go:generate mockgen -source=./client/collection/collection_client.go -destination=./mocks/collection_client_mock.go -package=mocks ClientService --additional-initialism=SLO
