// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/knockout/api/graphql/model"
)

// region    ************************** generated!.gotpl **************************

type SubscriptionResolver interface {
	Message(ctx context.Context) (<-chan *model.Message, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Message_id(ctx context.Context, field graphql.CollectedField, obj *model.Message) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Message_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Message_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Message",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Message_topic(ctx context.Context, field graphql.CollectedField, obj *model.Message) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Message_topic(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Topic, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Message_topic(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Message",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Message_body(ctx context.Context, field graphql.CollectedField, obj *model.Message) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Message_body(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Body, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Message_body(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Message",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Message_createdAt(ctx context.Context, field graphql.CollectedField, obj *model.Message) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Message_createdAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CreatedAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Message_createdAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Message",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Message_sentAt(ctx context.Context, field graphql.CollectedField, obj *model.Message) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Message_sentAt(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SentAt, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Message_sentAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Message",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Subscription_message(ctx context.Context, field graphql.CollectedField) (ret func(ctx context.Context) graphql.Marshaler) {
	fc, err := ec.fieldContext_Subscription_message(ctx, field)
	if err != nil {
		return nil
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Subscription().Message(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	if resTmp == nil {
		return nil
	}
	return func(ctx context.Context) graphql.Marshaler {
		select {
		case res, ok := <-resTmp.(<-chan *model.Message):
			if !ok {
				return nil
			}
			return graphql.WriterFunc(func(w io.Writer) {
				w.Write([]byte{'{'})
				graphql.MarshalString(field.Alias).MarshalGQL(w)
				w.Write([]byte{':'})
				ec.marshalOMessage2ᚖgithubᚗcomᚋwoocoosᚋknockoutᚋapiᚋgraphqlᚋmodelᚐMessage(ctx, field.Selections, res).MarshalGQL(w)
				w.Write([]byte{'}'})
			})
		case <-ctx.Done():
			return nil
		}
	}
}

func (ec *executionContext) fieldContext_Subscription_message(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Subscription",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Message_id(ctx, field)
			case "topic":
				return ec.fieldContext_Message_topic(ctx, field)
			case "body":
				return ec.fieldContext_Message_body(ctx, field)
			case "createdAt":
				return ec.fieldContext_Message_createdAt(ctx, field)
			case "sentAt":
				return ec.fieldContext_Message_sentAt(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Message", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var messageImplementors = []string{"Message"}

func (ec *executionContext) _Message(ctx context.Context, sel ast.SelectionSet, obj *model.Message) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, messageImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Message")
		case "id":

			out.Values[i] = ec._Message_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "topic":

			out.Values[i] = ec._Message_topic(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "body":

			out.Values[i] = ec._Message_body(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "createdAt":

			out.Values[i] = ec._Message_createdAt(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "sentAt":

			out.Values[i] = ec._Message_sentAt(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var subscriptionImplementors = []string{"Subscription"}

func (ec *executionContext) _Subscription(ctx context.Context, sel ast.SelectionSet) func(ctx context.Context) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, subscriptionImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Subscription",
	})
	if len(fields) != 1 {
		ec.Errorf(ctx, "must subscribe to exactly one stream")
		return nil
	}

	switch fields[0].Name {
	case "message":
		return ec._Subscription_message(ctx, fields[0])
	default:
		panic("unknown field " + strconv.Quote(fields[0].Name))
	}
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalOMessage2ᚖgithubᚗcomᚋwoocoosᚋknockoutᚋapiᚋgraphqlᚋmodelᚐMessage(ctx context.Context, sel ast.SelectionSet, v *model.Message) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Message(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************