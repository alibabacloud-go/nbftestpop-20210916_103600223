// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddRequest struct {
	X *int32 `json:"x,omitempty" xml:"x,omitempty"`
	Y *int32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s AddRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRequest) GoString() string {
	return s.String()
}

func (s *AddRequest) SetX(v int32) *AddRequest {
	s.X = &v
	return s
}

func (s *AddRequest) SetY(v int32) *AddRequest {
	s.Y = &v
	return s
}

type AddResponseBody struct {
	Sum *int32 `json:"sum,omitempty" xml:"sum,omitempty"`
}

func (s AddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResponseBody) GoString() string {
	return s.String()
}

func (s *AddResponseBody) SetSum(v int32) *AddResponseBody {
	s.Sum = &v
	return s
}

type AddResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddResponseBody   `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResponse) GoString() string {
	return s.String()
}

func (s *AddResponse) SetHeaders(v map[string]*string) *AddResponse {
	s.Headers = v
	return s
}

func (s *AddResponse) SetBody(v *AddResponseBody) *AddResponse {
	s.Body = v
	return s
}

type ResolveOpenApiRequest struct {
	Name    *string                       `json:"name,omitempty" xml:"name,omitempty"`
	Age     *int32                        `json:"age,omitempty" xml:"age,omitempty"`
	Teacher *ResolveOpenApiRequestTeacher `json:"teacher,omitempty" xml:"teacher,omitempty" type:"Struct"`
	Friends []map[string]interface{}      `json:"friends,omitempty" xml:"friends,omitempty" type:"Repeated"`
}

func (s ResolveOpenApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ResolveOpenApiRequest) GoString() string {
	return s.String()
}

func (s *ResolveOpenApiRequest) SetName(v string) *ResolveOpenApiRequest {
	s.Name = &v
	return s
}

func (s *ResolveOpenApiRequest) SetAge(v int32) *ResolveOpenApiRequest {
	s.Age = &v
	return s
}

func (s *ResolveOpenApiRequest) SetTeacher(v *ResolveOpenApiRequestTeacher) *ResolveOpenApiRequest {
	s.Teacher = v
	return s
}

func (s *ResolveOpenApiRequest) SetFriends(v []map[string]interface{}) *ResolveOpenApiRequest {
	s.Friends = v
	return s
}

type ResolveOpenApiRequestTeacher struct {
	TeacherAge  *int32  `json:"teacherAge,omitempty" xml:"teacherAge,omitempty"`
	TeacherName *string `json:"teacherName,omitempty" xml:"teacherName,omitempty"`
}

func (s ResolveOpenApiRequestTeacher) String() string {
	return tea.Prettify(s)
}

func (s ResolveOpenApiRequestTeacher) GoString() string {
	return s.String()
}

func (s *ResolveOpenApiRequestTeacher) SetTeacherAge(v int32) *ResolveOpenApiRequestTeacher {
	s.TeacherAge = &v
	return s
}

func (s *ResolveOpenApiRequestTeacher) SetTeacherName(v string) *ResolveOpenApiRequestTeacher {
	s.TeacherName = &v
	return s
}

type ResolveOpenApiResponseBody struct {
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
}

func (s ResolveOpenApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResolveOpenApiResponseBody) GoString() string {
	return s.String()
}

func (s *ResolveOpenApiResponseBody) SetSuccess(v bool) *ResolveOpenApiResponseBody {
	s.Success = &v
	return s
}

func (s *ResolveOpenApiResponseBody) SetErrorCode(v string) *ResolveOpenApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResolveOpenApiResponseBody) SetErrorMsg(v string) *ResolveOpenApiResponseBody {
	s.ErrorMsg = &v
	return s
}

type ResolveOpenApiResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResolveOpenApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResolveOpenApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ResolveOpenApiResponse) GoString() string {
	return s.String()
}

func (s *ResolveOpenApiResponse) SetHeaders(v map[string]*string) *ResolveOpenApiResponse {
	s.Headers = v
	return s
}

func (s *ResolveOpenApiResponse) SetBody(v *ResolveOpenApiResponseBody) *ResolveOpenApiResponse {
	s.Body = v
	return s
}

type AssetPublishTestOpenApiRequest struct {
	Name    *string                                `json:"name,omitempty" xml:"name,omitempty"`
	Teacher *AssetPublishTestOpenApiRequestTeacher `json:"teacher,omitempty" xml:"teacher,omitempty" type:"Struct"`
}

func (s AssetPublishTestOpenApiRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetPublishTestOpenApiRequest) GoString() string {
	return s.String()
}

func (s *AssetPublishTestOpenApiRequest) SetName(v string) *AssetPublishTestOpenApiRequest {
	s.Name = &v
	return s
}

func (s *AssetPublishTestOpenApiRequest) SetTeacher(v *AssetPublishTestOpenApiRequestTeacher) *AssetPublishTestOpenApiRequest {
	s.Teacher = v
	return s
}

type AssetPublishTestOpenApiRequestTeacher struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Age  *int32  `json:"age,omitempty" xml:"age,omitempty"`
}

func (s AssetPublishTestOpenApiRequestTeacher) String() string {
	return tea.Prettify(s)
}

func (s AssetPublishTestOpenApiRequestTeacher) GoString() string {
	return s.String()
}

func (s *AssetPublishTestOpenApiRequestTeacher) SetName(v string) *AssetPublishTestOpenApiRequestTeacher {
	s.Name = &v
	return s
}

func (s *AssetPublishTestOpenApiRequestTeacher) SetAge(v int32) *AssetPublishTestOpenApiRequestTeacher {
	s.Age = &v
	return s
}

type AssetPublishTestOpenApiResponseBody struct {
	Teacher   *AssetPublishTestOpenApiResponseBodyTeacher `json:"teacher,omitempty" xml:"teacher,omitempty" type:"Struct"`
	Success   *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	ErrorCode *string                                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
}

func (s AssetPublishTestOpenApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssetPublishTestOpenApiResponseBody) GoString() string {
	return s.String()
}

func (s *AssetPublishTestOpenApiResponseBody) SetTeacher(v *AssetPublishTestOpenApiResponseBodyTeacher) *AssetPublishTestOpenApiResponseBody {
	s.Teacher = v
	return s
}

func (s *AssetPublishTestOpenApiResponseBody) SetSuccess(v bool) *AssetPublishTestOpenApiResponseBody {
	s.Success = &v
	return s
}

func (s *AssetPublishTestOpenApiResponseBody) SetErrorCode(v string) *AssetPublishTestOpenApiResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AssetPublishTestOpenApiResponseBody) SetErrorMsg(v string) *AssetPublishTestOpenApiResponseBody {
	s.ErrorMsg = &v
	return s
}

type AssetPublishTestOpenApiResponseBodyTeacher struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Age  *string `json:"age,omitempty" xml:"age,omitempty"`
}

func (s AssetPublishTestOpenApiResponseBodyTeacher) String() string {
	return tea.Prettify(s)
}

func (s AssetPublishTestOpenApiResponseBodyTeacher) GoString() string {
	return s.String()
}

func (s *AssetPublishTestOpenApiResponseBodyTeacher) SetName(v string) *AssetPublishTestOpenApiResponseBodyTeacher {
	s.Name = &v
	return s
}

func (s *AssetPublishTestOpenApiResponseBodyTeacher) SetAge(v string) *AssetPublishTestOpenApiResponseBodyTeacher {
	s.Age = &v
	return s
}

type AssetPublishTestOpenApiResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssetPublishTestOpenApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssetPublishTestOpenApiResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetPublishTestOpenApiResponse) GoString() string {
	return s.String()
}

func (s *AssetPublishTestOpenApiResponse) SetHeaders(v map[string]*string) *AssetPublishTestOpenApiResponse {
	s.Headers = v
	return s
}

func (s *AssetPublishTestOpenApiResponse) SetBody(v *AssetPublishTestOpenApiResponseBody) *AssetPublishTestOpenApiResponse {
	s.Body = v
	return s
}

type AddTestRequest struct {
	X *int32 `json:"x,omitempty" xml:"x,omitempty"`
	Y *int32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s AddTestRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTestRequest) GoString() string {
	return s.String()
}

func (s *AddTestRequest) SetX(v int32) *AddTestRequest {
	s.X = &v
	return s
}

func (s *AddTestRequest) SetY(v int32) *AddTestRequest {
	s.Y = &v
	return s
}

type AddTestResponseBody struct {
	Sum *int32 `json:"sum,omitempty" xml:"sum,omitempty"`
}

func (s AddTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTestResponseBody) GoString() string {
	return s.String()
}

func (s *AddTestResponseBody) SetSum(v int32) *AddTestResponseBody {
	s.Sum = &v
	return s
}

type AddTestResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTestResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTestResponse) GoString() string {
	return s.String()
}

func (s *AddTestResponse) SetHeaders(v map[string]*string) *AddTestResponse {
	s.Headers = v
	return s
}

func (s *AddTestResponse) SetBody(v *AddTestResponseBody) *AddTestResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nbftestpop"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Add(request *AddRequest) (_result *AddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddResponse{}
	_body, _err := client.AddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWithOptions(request *AddRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.X)) {
		query["x"] = request.X
	}

	if !tea.BoolValue(util.IsUnset(request.Y)) {
		query["y"] = request.Y
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddResponse{}
	_body, _err := client.DoROARequest(tea.String("Add"), tea.String("2021-09-16_10-36-00-223"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/kxRoaProduct/9_0_9/add"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResolveOpenApi(request *ResolveOpenApiRequest) (_result *ResolveOpenApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResolveOpenApiResponse{}
	_body, _err := client.ResolveOpenApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResolveOpenApiWithOptions(request *ResolveOpenApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResolveOpenApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Age)) {
		query["age"] = request.Age
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Teacher))) {
		query["teacher"] = request.Teacher
	}

	if !tea.BoolValue(util.IsUnset(request.Friends)) {
		query["friends"] = request.Friends
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ResolveOpenApiResponse{}
	_body, _err := client.DoROARequest(tea.String("ResolveOpenApi"), tea.String("2021-09-16_10-36-00-223"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/kxRoaProduct/9_0_9/resolveOpenApi"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssetPublishTestOpenApi(request *AssetPublishTestOpenApiRequest) (_result *AssetPublishTestOpenApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AssetPublishTestOpenApiResponse{}
	_body, _err := client.AssetPublishTestOpenApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssetPublishTestOpenApiWithOptions(request *AssetPublishTestOpenApiRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AssetPublishTestOpenApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Teacher))) {
		query["teacher"] = request.Teacher
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AssetPublishTestOpenApiResponse{}
	_body, _err := client.DoROARequest(tea.String("AssetPublishTestOpenApi"), tea.String("2021-09-16_10-36-00-223"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/kxRoaProduct/9_0_9/assetPublishTestOpenApi"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTest(request *AddTestRequest) (_result *AddTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddTestResponse{}
	_body, _err := client.AddTestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTestWithOptions(request *AddTestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.X)) {
		query["x"] = request.X
	}

	if !tea.BoolValue(util.IsUnset(request.Y)) {
		query["y"] = request.Y
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddTestResponse{}
	_body, _err := client.DoROARequest(tea.String("AddTest"), tea.String("2021-09-16_10-36-00-223"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/kxRoaProduct/9_0_9/addTest"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
