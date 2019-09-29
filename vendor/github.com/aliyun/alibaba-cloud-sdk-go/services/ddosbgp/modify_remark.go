package ddosbgp

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

// ModifyRemark invokes the ddosbgp.ModifyRemark API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/modifyremark.html
func (client *Client) ModifyRemark(request *ModifyRemarkRequest) (response *ModifyRemarkResponse, err error) {
	response = CreateModifyRemarkResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRemarkWithChan invokes the ddosbgp.ModifyRemark API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/modifyremark.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRemarkWithChan(request *ModifyRemarkRequest) (<-chan *ModifyRemarkResponse, <-chan error) {
	responseChan := make(chan *ModifyRemarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRemark(request)
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

// ModifyRemarkWithCallback invokes the ddosbgp.ModifyRemark API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/modifyremark.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRemarkWithCallback(request *ModifyRemarkRequest, callback func(response *ModifyRemarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRemarkResponse
		var err error
		defer close(result)
		response, err = client.ModifyRemark(request)
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

// ModifyRemarkRequest is the request struct for api ModifyRemark
type ModifyRemarkRequest struct {
	*requests.RpcRequest
	Remark           string `position:"Query" name:"Remark"`
	ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	SourceIp         string `position:"Query" name:"SourceIp"`
	Lang             string `position:"Query" name:"Lang"`
	ResourceRegionId string `position:"Query" name:"ResourceRegionId"`
}

// ModifyRemarkResponse is the response struct for api ModifyRemark
type ModifyRemarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRemarkRequest creates a request to invoke ModifyRemark API
func CreateModifyRemarkRequest() (request *ModifyRemarkRequest) {
	request = &ModifyRemarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "ModifyRemark", "ddosbgp", "openAPI")
	return
}

// CreateModifyRemarkResponse creates a response to parse from ModifyRemark response
func CreateModifyRemarkResponse() (response *ModifyRemarkResponse) {
	response = &ModifyRemarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}