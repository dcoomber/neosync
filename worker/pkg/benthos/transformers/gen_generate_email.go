
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_email.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateEmail struct{}

type GenerateEmailOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
	emailType string
}

func NewGenerateEmail() *GenerateEmail {
	return &GenerateEmail{}
}

func (t *GenerateEmail) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateEmail",
		Description: "",
	}, nil
}

func (t *GenerateEmail) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateEmailOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 100000
	}
	transformerOpts.maxLength = maxLength

	emailType, ok := opts["emailType"].(string)
	if !ok {
		emailType = GenerateEmailType_UuidV4.String()
	}
	transformerOpts.emailType = emailType

	var seed int64
	seedArg, ok := opts["seed"].(int64)
	if ok {
		seed = seedArg
	} else {
		var err error
		seed, err = transformer_utils.GenerateCryptoSeed()
		if err != nil {
			return nil, fmt.Errorf("unable to generate seed: %w", err)
		}
	}
	transformerOpts.randomizer = rng.New(seed)

	return transformerOpts, nil
}