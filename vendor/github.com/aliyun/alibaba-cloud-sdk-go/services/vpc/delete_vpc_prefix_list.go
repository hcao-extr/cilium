package vpc

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

// DeleteVpcPrefixList invokes the vpc.DeleteVpcPrefixList API synchronously
func (client *Client) DeleteVpcPrefixList(request *DeleteVpcPrefixListRequest) (response *DeleteVpcPrefixListResponse, err error) {
	response = CreateDeleteVpcPrefixListResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVpcPrefixListWithChan invokes the vpc.DeleteVpcPrefixList API asynchronously
func (client *Client) DeleteVpcPrefixListWithChan(request *DeleteVpcPrefixListRequest) (<-chan *DeleteVpcPrefixListResponse, <-chan error) {
	responseChan := make(chan *DeleteVpcPrefixListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVpcPrefixList(request)
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

// DeleteVpcPrefixListWithCallback invokes the vpc.DeleteVpcPrefixList API asynchronously
func (client *Client) DeleteVpcPrefixListWithCallback(request *DeleteVpcPrefixListRequest, callback func(response *DeleteVpcPrefixListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVpcPrefixListResponse
		var err error
		defer close(result)
		response, err = client.DeleteVpcPrefixList(request)
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

// DeleteVpcPrefixListRequest is the request struct for api DeleteVpcPrefixList
type DeleteVpcPrefixListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PrefixListId         string           `position:"Query" name:"PrefixListId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteVpcPrefixListResponse is the response struct for api DeleteVpcPrefixList
type DeleteVpcPrefixListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVpcPrefixListRequest creates a request to invoke DeleteVpcPrefixList API
func CreateDeleteVpcPrefixListRequest() (request *DeleteVpcPrefixListRequest) {
	request = &DeleteVpcPrefixListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpcPrefixList", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVpcPrefixListResponse creates a response to parse from DeleteVpcPrefixList response
func CreateDeleteVpcPrefixListResponse() (response *DeleteVpcPrefixListResponse) {
	response = &DeleteVpcPrefixListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}