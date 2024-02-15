// Copyright 2023 Chronosphere Inc.
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

package swagger

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// TextConsumer creates a new text consumer.
//
// Specifically we attempt to deal with the case here where nginx returns a 500
// that is HTML, so we can't unmarshal into our RuntimeError object. In the case
// HTML is returned, we set the error's message field to the HTML body.
func TextConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data any) error {
		if reader == nil {
			return fmt.Errorf("unexpected response, missing response body")
		}

		b, err := io.ReadAll(reader)
		if err != nil {
			return errors.Wrap(err, "could not read unexpected response body")
		}

		// If we're receiving a APIError, meaning a non 200 code occurred,
		// put our response body into the message.
		if err, ok := data.(*models.APIError); ok {
			err.Message = string(b)
			err.Code = int32(codes.Unknown)

			// attach request ID if available.
			if ider, ok := reader.(RequestIDer); ok {
				err.Message = fmt.Sprintf("%s (request_id=%s)", err.Message, ider.RequestID())
			}
			return nil
		}

		// We don't know what to do with our response body.
		return fmt.Errorf("received an unexpected response type as text, response body: %s", b)
	})
}
