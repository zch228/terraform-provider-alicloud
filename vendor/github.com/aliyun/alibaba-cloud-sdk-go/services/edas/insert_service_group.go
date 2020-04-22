package edas

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

// InsertServiceGroup invokes the edas.InsertServiceGroup API synchronously
// api document: https://help.aliyun.com/api/edas/insertservicegroup.html
func (client *Client) InsertServiceGroup(request *InsertServiceGroupRequest) (response *InsertServiceGroupResponse, err error) {
	response = CreateInsertServiceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// InsertServiceGroupWithChan invokes the edas.InsertServiceGroup API asynchronously
// api document: https://help.aliyun.com/api/edas/insertservicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertServiceGroupWithChan(request *InsertServiceGroupRequest) (<-chan *InsertServiceGroupResponse, <-chan error) {
	responseChan := make(chan *InsertServiceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertServiceGroup(request)
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

// InsertServiceGroupWithCallback invokes the edas.InsertServiceGroup API asynchronously
// api document: https://help.aliyun.com/api/edas/insertservicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertServiceGroupWithCallback(request *InsertServiceGroupRequest, callback func(response *InsertServiceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertServiceGroupResponse
		var err error
		defer close(result)
		response, err = client.InsertServiceGroup(request)
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

// InsertServiceGroupRequest is the request struct for api InsertServiceGroup
type InsertServiceGroupRequest struct {
	*requests.RoaRequest
	GroupName string `position:"Query" name:"GroupName"`
}

// InsertServiceGroupResponse is the response struct for api InsertServiceGroup
type InsertServiceGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInsertServiceGroupRequest creates a request to invoke InsertServiceGroup API
func CreateInsertServiceGroupRequest() (request *InsertServiceGroupRequest) {
	request = &InsertServiceGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "InsertServiceGroup", "/pop/v5/service/serviceGroups", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInsertServiceGroupResponse creates a response to parse from InsertServiceGroup response
func CreateInsertServiceGroupResponse() (response *InsertServiceGroupResponse) {
	response = &InsertServiceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
