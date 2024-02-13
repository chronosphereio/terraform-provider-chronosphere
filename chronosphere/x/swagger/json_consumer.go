package swagger

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// JSONConsumer is a fork of https://github.com/go-openapi/runtime/blob/master/json.go
func JSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data any) error {
		dec := json.NewDecoder(reader)
		dec.UseNumber() // preserve number formats
		if err := dec.Decode(data); err != nil {
			return err
		}

		// If we're receiving a RuntimeError, meaning a non 200 code occurred, we include
		// our request ID in the error message.
		if err, ok := data.(*models.APIError); ok {
			if ider, ok := reader.(RequestIDer); ok {
				err.Message = fmt.Sprintf("%s (request_id=%s)", err.Message, ider.RequestID())
			}
		}
		return nil
	})
}
