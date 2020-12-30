package aws

import (
	"encoding/base64"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

func TestPublisher(t *testing.T) {
	snstest := &TestSNSAPI{}
	pub := &publisher{sns: snstest}

	test1Key := "yo!"
	test1 := &TestProto{"hi there!"}
	err := pub.Publish(context.Background(), test1Key, test1)
	if err != nil {
		t.Error("Publish returned an unexpected error: ", err)
	}

	if len(snstest.Published) != 1 {
		t.Error("Publish expected 1 published input, got: ", len(snstest.Published))
		return
	}

	var (
		got     TestProto
		gotBody []byte
	)
	gotBody, err = base64.StdEncoding.DecodeString(*snstest.Published[0].Message)
	if err != nil {
		t.Error("Encountered unexpected error decoding message: ", err)
	}

	err = proto.Unmarshal(gotBody, &got)
	if err != nil {
		t.Error("Encountered unexpected error unmarshalling proto message: ", err)
		return
	}

	if !reflect.DeepEqual(test1, &got) {
		t.Errorf("Publish expected message of \"%#v\", got: %#v", test1, got)
		return
	}

	if *snstest.Published[0].Subject != test1Key {
		t.Errorf("Publish expected subject of \"%s\", got: \"%s\"", test1Key, *snstest.Published[0].Subject)
	}
}

type TestSNSAPI struct {
	// Error will be returned by the API when Publish() is called.
	Error error
	// Published allows users to inspect which values have been published.
	Published []*sns.PublishInput
}

var _ snsiface.SNSAPI = &TestSNSAPI{}

func (t *TestSNSAPI) Publish(i *sns.PublishInput) (*sns.PublishOutput, error) {
	t.Published = append(t.Published, i)
	return &sns.PublishOutput{}, t.Error
}

///////////
// ALL METHODS BELOW HERE ARE EMPTY AND JUST SATISFYING THE SQSAPI interface
///////////

var errNotImpl = errors.New("method not implemented")

func (t *TestSNSAPI) SetSMSAttributesRequest(*sns.SetSMSAttributesInput) (*request.Request, *sns.SetSMSAttributesOutput) {
	return nil, nil
}

func (t *TestSNSAPI) SetSMSAttributes(*sns.SetSMSAttributesInput) (*sns.SetSMSAttributesOutput, error) {
	return nil, nil
}

func (t *TestSNSAPI) OptInPhoneNumberRequest(*sns.OptInPhoneNumberInput) (*request.Request, *sns.OptInPhoneNumberOutput) {
	return nil, nil
}

func (t *TestSNSAPI) OptInPhoneNumber(*sns.OptInPhoneNumberInput) (*sns.OptInPhoneNumberOutput, error) {
	return nil, nil
}

func (t *TestSNSAPI) ListPhoneNumbersOptedOutRequest(*sns.ListPhoneNumbersOptedOutInput) (*request.Request, *sns.ListPhoneNumbersOptedOutOutput) {
	return nil, nil
}

func (t *TestSNSAPI) ListPhoneNumbersOptedOut(*sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	return nil, nil
}

func (t *TestSNSAPI) GetSMSAttributesRequest(*sns.GetSMSAttributesInput) (*request.Request, *sns.GetSMSAttributesOutput) {
	return nil, nil
}

func (t *TestSNSAPI) GetSMSAttributes(*sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error) {
	return nil, nil
}

func (t *TestSNSAPI) CheckIfPhoneNumberIsOptedOutRequest(*sns.CheckIfPhoneNumberIsOptedOutInput) (*request.Request, *sns.CheckIfPhoneNumberIsOptedOutOutput) {
	return nil, nil
}
func (t *TestSNSAPI) CheckIfPhoneNumberIsOptedOut(*sns.CheckIfPhoneNumberIsOptedOutInput) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
	return nil, nil
}
func (t *TestSNSAPI) AddPermissionRequest(*sns.AddPermissionInput) (*request.Request, *sns.AddPermissionOutput) {
	return nil, nil
}
func (t *TestSNSAPI) AddPermission(*sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ConfirmSubscriptionRequest(*sns.ConfirmSubscriptionInput) (*request.Request, *sns.ConfirmSubscriptionOutput) {
	return nil, nil
}
func (t *TestSNSAPI) ConfirmSubscription(*sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) CreatePlatformApplicationRequest(*sns.CreatePlatformApplicationInput) (*request.Request, *sns.CreatePlatformApplicationOutput) {
	return nil, nil
}
func (t *TestSNSAPI) CreatePlatformApplication(*sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) CreatePlatformEndpointRequest(*sns.CreatePlatformEndpointInput) (*request.Request, *sns.CreatePlatformEndpointOutput) {
	return nil, nil
}
func (t *TestSNSAPI) CreatePlatformEndpoint(*sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) CreateTopicRequest(*sns.CreateTopicInput) (*request.Request, *sns.CreateTopicOutput) {
	return nil, nil
}
func (t *TestSNSAPI) CreateTopic(*sns.CreateTopicInput) (*sns.CreateTopicOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) DeleteEndpointRequest(*sns.DeleteEndpointInput) (*request.Request, *sns.DeleteEndpointOutput) {
	return nil, nil
}
func (t *TestSNSAPI) DeleteEndpoint(*sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) DeletePlatformApplicationRequest(*sns.DeletePlatformApplicationInput) (*request.Request, *sns.DeletePlatformApplicationOutput) {
	return nil, nil
}
func (t *TestSNSAPI) DeletePlatformApplication(*sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) DeleteTopicRequest(*sns.DeleteTopicInput) (*request.Request, *sns.DeleteTopicOutput) {
	return nil, nil
}
func (t *TestSNSAPI) DeleteTopic(*sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) GetEndpointAttributesRequest(*sns.GetEndpointAttributesInput) (*request.Request, *sns.GetEndpointAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) GetEndpointAttributes(*sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) GetPlatformApplicationAttributesRequest(*sns.GetPlatformApplicationAttributesInput) (*request.Request, *sns.GetPlatformApplicationAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) GetPlatformApplicationAttributes(*sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) GetSubscriptionAttributesRequest(*sns.GetSubscriptionAttributesInput) (*request.Request, *sns.GetSubscriptionAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) GetSubscriptionAttributes(*sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) GetTopicAttributesRequest(*sns.GetTopicAttributesInput) (*request.Request, *sns.GetTopicAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) GetTopicAttributes(*sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ListEndpointsByPlatformApplicationRequest(*sns.ListEndpointsByPlatformApplicationInput) (*request.Request, *sns.ListEndpointsByPlatformApplicationOutput) {
	return nil, nil
}
func (t *TestSNSAPI) ListEndpointsByPlatformApplication(*sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ListEndpointsByPlatformApplicationPages(*sns.ListEndpointsByPlatformApplicationInput, func(*sns.ListEndpointsByPlatformApplicationOutput, bool) bool) error {
	return nil
}
func (t *TestSNSAPI) ListPlatformApplicationsRequest(*sns.ListPlatformApplicationsInput) (*request.Request, *sns.ListPlatformApplicationsOutput) {
	return nil, nil
}
func (t *TestSNSAPI) ListPlatformApplications(*sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ListPlatformApplicationsPages(*sns.ListPlatformApplicationsInput, func(*sns.ListPlatformApplicationsOutput, bool) bool) error {
	return nil
}
func (t *TestSNSAPI) ListSubscriptionsRequest(*sns.ListSubscriptionsInput) (*request.Request, *sns.ListSubscriptionsOutput) {
	return nil, nil
}
func (t *TestSNSAPI) ListSubscriptions(*sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ListSubscriptionsPages(*sns.ListSubscriptionsInput, func(*sns.ListSubscriptionsOutput, bool) bool) error {
	return errNotImpl
}
func (t *TestSNSAPI) ListSubscriptionsByTopicRequest(*sns.ListSubscriptionsByTopicInput) (*request.Request, *sns.ListSubscriptionsByTopicOutput) {
	return nil, nil
}
func (t *TestSNSAPI) ListSubscriptionsByTopic(*sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ListSubscriptionsByTopicPages(*sns.ListSubscriptionsByTopicInput, func(*sns.ListSubscriptionsByTopicOutput, bool) bool) error {
	return errNotImpl
}
func (t *TestSNSAPI) ListTopicsRequest(*sns.ListTopicsInput) (*request.Request, *sns.ListTopicsOutput) {
	return nil, nil
}
func (t *TestSNSAPI) ListTopics(*sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) ListTopicsPages(*sns.ListTopicsInput, func(*sns.ListTopicsOutput, bool) bool) error {
	return errNotImpl
}
func (t *TestSNSAPI) PublishRequest(*sns.PublishInput) (*request.Request, *sns.PublishOutput) {
	return nil, nil
}

func (t *TestSNSAPI) RemovePermissionRequest(*sns.RemovePermissionInput) (*request.Request, *sns.RemovePermissionOutput) {
	return nil, nil
}
func (t *TestSNSAPI) RemovePermission(*sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) SetEndpointAttributesRequest(*sns.SetEndpointAttributesInput) (*request.Request, *sns.SetEndpointAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) SetEndpointAttributes(*sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) SetPlatformApplicationAttributesRequest(*sns.SetPlatformApplicationAttributesInput) (*request.Request, *sns.SetPlatformApplicationAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) SetPlatformApplicationAttributes(*sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) SetSubscriptionAttributesRequest(*sns.SetSubscriptionAttributesInput) (*request.Request, *sns.SetSubscriptionAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) SetSubscriptionAttributes(*sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) SetTopicAttributesRequest(*sns.SetTopicAttributesInput) (*request.Request, *sns.SetTopicAttributesOutput) {
	return nil, nil
}
func (t *TestSNSAPI) SetTopicAttributes(*sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) SubscribeRequest(*sns.SubscribeInput) (*request.Request, *sns.SubscribeOutput) {
	return nil, nil
}
func (t *TestSNSAPI) Subscribe(*sns.SubscribeInput) (*sns.SubscribeOutput, error) {
	return nil, errNotImpl
}
func (t *TestSNSAPI) UnsubscribeRequest(*sns.UnsubscribeInput) (*request.Request, *sns.UnsubscribeOutput) {
	return nil, nil
}
func (t *TestSNSAPI) Unsubscribe(*sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error) {
	return nil, errNotImpl
}
