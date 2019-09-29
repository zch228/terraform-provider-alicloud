package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateETLJob invokes the emr.UpdateETLJob API synchronously
// api document: https://help.aliyun.com/api/emr/updateetljob.html
func (client *Client) UpdateETLJob(request *UpdateETLJobRequest) (response *UpdateETLJobResponse, err error) {
	response = CreateUpdateETLJobResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateETLJobWithChan invokes the emr.UpdateETLJob API asynchronously
// api document: https://help.aliyun.com/api/emr/updateetljob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateETLJobWithChan(request *UpdateETLJobRequest) (<-chan *UpdateETLJobResponse, <-chan error) {
	responseChan := make(chan *UpdateETLJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateETLJob(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateETLJobWithCallback invokes the emr.UpdateETLJob API asynchronously
// api document: https://help.aliyun.com/api/emr/updateetljob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateETLJobWithCallback(request *UpdateETLJobRequest, callback func(response *UpdateETLJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateETLJobResponse
		var err error
		defer close(result)
		response, err = client.UpdateETLJob(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateETLJobRequest is the request struct for api UpdateETLJob
type UpdateETLJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	ClusterConfig   string                         `position:"Query" name:"ClusterConfig"`
	TriggerRule     *[]UpdateETLJobTriggerRule     `position:"Query" name:"TriggerRule"  type:"Repeated"`
	AlertConfig     string                         `position:"Query" name:"AlertConfig"`
	Description     string                         `position:"Query" name:"Description"`
	Check           requests.Boolean               `position:"Query" name:"Check"`
	StageConnection *[]UpdateETLJobStageConnection `position:"Query" name:"StageConnection"  type:"Repeated"`
	Stage           *[]UpdateETLJobStage           `position:"Query" name:"Stage"  type:"Repeated"`
	Name            string                         `position:"Query" name:"Name"`
	Id              string                         `position:"Query" name:"Id"`
}

// UpdateETLJobTriggerRule is a repeated param struct in UpdateETLJobRequest
type UpdateETLJobTriggerRule struct {
	CronExpr  string `name:"CronExpr"`
	EndTime   string `name:"EndTime"`
	StartTime string `name:"StartTime"`
	Enabled   string `name:"Enabled"`
}

// UpdateETLJobStageConnection is a repeated param struct in UpdateETLJobRequest
type UpdateETLJobStageConnection struct {
	Port string `name:"Port"`
	From string `name:"From"`
	To   string `name:"To"`
}

// UpdateETLJobStage is a repeated param struct in UpdateETLJobRequest
type UpdateETLJobStage struct {
	StageName   string `name:"StageName"`
	StageConf   string `name:"StageConf"`
	StageType   string `name:"StageType"`
	StagePlugin string `name:"StagePlugin"`
}

// UpdateETLJobResponse is the response struct for api UpdateETLJob
type UpdateETLJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateETLJobRequest creates a request to invoke UpdateETLJob API
func CreateUpdateETLJobRequest() (request *UpdateETLJobRequest) {
	request = &UpdateETLJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "UpdateETLJob", "emr", "openAPI")
	return
}

// CreateUpdateETLJobResponse creates a response to parse from UpdateETLJob response
func CreateUpdateETLJobResponse() (response *UpdateETLJobResponse) {
	response = &UpdateETLJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}