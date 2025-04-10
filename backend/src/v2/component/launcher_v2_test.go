// Copyright 2023 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package component

import (
	"context"
	"github.com/kubeflow/pipelines/backend/src/v2/cacheutils"
	"k8s.io/client-go/kubernetes"
	"testing"

	"github.com/kubeflow/pipelines/api/v2alpha1/go/pipelinespec"
	"github.com/kubeflow/pipelines/backend/src/v2/metadata"
	"github.com/kubeflow/pipelines/backend/src/v2/objectstore"
	"github.com/stretchr/testify/assert"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/memblob"
	"google.golang.org/protobuf/types/known/structpb"
	"k8s.io/client-go/kubernetes/fake"
)

var addNumbersComponent = &pipelinespec.ComponentSpec{
	Implementation: &pipelinespec.ComponentSpec_ExecutorLabel{ExecutorLabel: "add"},
	InputDefinitions: &pipelinespec.ComponentInputsSpec{
		Parameters: map[string]*pipelinespec.ComponentInputsSpec_ParameterSpec{
			"a": {ParameterType: pipelinespec.ParameterType_NUMBER_INTEGER, DefaultValue: structpb.NewNumberValue(5)},
			"b": {ParameterType: pipelinespec.ParameterType_NUMBER_INTEGER},
		},
	},
	OutputDefinitions: &pipelinespec.ComponentOutputsSpec{
		Parameters: map[string]*pipelinespec.ComponentOutputsSpec_ParameterSpec{
			"Output": {ParameterType: pipelinespec.ParameterType_NUMBER_INTEGER},
		},
	},
}

// Tests happy and unhappy paths for constructing a new LauncherV2
func Test_NewLauncherV2(t *testing.T) {
	cmdArgs := []string{"sh", "-c", "echo \"hello world\""}
	opts := &LauncherV2Options{
		Namespace:         "my-namespace",
		PodName:           "my-pod",
		PodUID:            "abcd",
		MLMDServerAddress: "example.com",
		MLMDServerPort:    "1234",
	}
	deps := LauncherV2Dependencies{
		k8sClientProvider: func() (kubernetes.Interface, error) {
			return &fake.Clientset{}, nil
		},
		metadataClientProvider: func(address string, port string) (*metadata.Client, error) {
			return &metadata.Client{}, nil
		},
		cacheClientProvider: func() (*cacheutils.Client, error) {
			return &cacheutils.Client{}, nil
		},
	}
	type args struct {
		ctx               context.Context
		executionID       int64
		executorInputJSON string
		componentSpecJSON string
	}
	tests := []struct {
		name    string
		args    *args
		wantErr bool
	}{
		{
			name: "happy path",
			args: &args{
				ctx:               context.TODO(),
				executionID:       1,
				executorInputJSON: "{}",
				componentSpecJSON: "{}",
			},
			wantErr: false,
		},
		{
			name: "missing executionID",
			args: &args{
				ctx:               context.TODO(),
				executionID:       0,
				executorInputJSON: "{}",
				componentSpecJSON: "{}",
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := test.args
			_, err := NewLauncherV2(args.ctx, args.executionID, args.executorInputJSON, args.componentSpecJSON, cmdArgs, opts, &deps)
			if (err != nil) != test.wantErr {
				t.Errorf("NewLauncherV2() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

// Tests that launcher correctly executes the user component and successfully writes output parameters to file.
func Test_executeV2_Parameters(t *testing.T) {
	tests := []struct {
		name          string
		executorInput *pipelinespec.ExecutorInput
		executorArgs  []string
		wantErr       bool
	}{
		{
			"happy pass",
			&pipelinespec.ExecutorInput{
				Inputs: &pipelinespec.ExecutorInput_Inputs{
					ParameterValues: map[string]*structpb.Value{"a": structpb.NewNumberValue(1), "b": structpb.NewNumberValue(2)},
				},
			},
			[]string{"-c", "test {{$.inputs.parameters['a']}} -eq 1 || exit 1\ntest {{$.inputs.parameters['b']}} -eq 2 || exit 1"},
			false,
		},
		{
			"use default value",
			&pipelinespec.ExecutorInput{
				Inputs: &pipelinespec.ExecutorInput_Inputs{
					ParameterValues: map[string]*structpb.Value{"b": structpb.NewNumberValue(2)},
				},
			},
			[]string{"-c", "test {{$.inputs.parameters['a']}} -eq 5 || exit 1\ntest {{$.inputs.parameters['b']}} -eq 2 || exit 1"},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fakeKubernetesClientset := &fake.Clientset{}
			fakeMetadataClient := metadata.NewFakeClient()
			bucket, err := blob.OpenBucket(context.Background(), "mem://test-bucket")
			assert.Nil(t, err)
			bucketConfig, err := objectstore.ParseBucketConfig("mem://test-bucket/pipeline-root/", nil)
			assert.Nil(t, err)
			_, _, err = executeV2(context.Background(), test.executorInput, addNumbersComponent, "sh", test.executorArgs, bucket, bucketConfig, fakeMetadataClient, "namespace", fakeKubernetesClientset)

			if test.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)

			}
		})
	}
}
