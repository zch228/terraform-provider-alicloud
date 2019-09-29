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

// CreateParagraph invokes the emr.CreateParagraph API synchronously
// api document: https://help.aliyun.com/api/emr/createparagraph.html
func (client *Client) CreateParagraph(request *CreateParagraphRequest) (response *CreateParagraphResponse, err error) {
	response = CreateCreateParagraphResponse()
	err = client.DoAction(request, response)
	return
}

// CreateParagraphWithChan invokes the emr.CreateParagraph API asynchronously
// api document: https://help.aliyun.com/api/emr/createparagraph.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateParagraphWithChan(request *CreateParagraphRequest) (<-chan *CreateParagraphResponse, <-chan error) {
	responseChan := make(chan *CreateParagraphResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateParagraph(request)
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

// CreateParagraphWithCallback invokes the emr.CreateParagraph API asynchronously
// api document: https://help.aliyun.com/api/emr/createparagraph.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateParagraphWithCallback(request *CreateParagraphRequest, callback func(response *CreateParagraphResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateParagraphResponse
		var err error
		defer close(result)
		response, err = client.CreateParagraph(request)
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

// CreateParagraphRequest is the request struct for api CreateParagraph
type CreateParagraphRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NoteId          string           `position:"Query" name:"NoteId"`
	Text            string           `position:"Query" name:"Text"`
}

// CreateParagraphResponse is the response struct for api CreateParagraph
type CreateParagraphResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateCreateParagraphRequest creates a request to invoke CreateParagraph API
func CreateCreateParagraphRequest() (request *CreateParagraphRequest) {
	request = &CreateParagraphRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CreateParagraph", "emr", "openAPI")
	return
}

// CreateCreateParagraphResponse creates a response to parse from CreateParagraph response
func CreateCreateParagraphResponse() (response *CreateParagraphResponse) {
	response = &CreateParagraphResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}