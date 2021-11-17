// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package time

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// This resource can be imported using the base UTC RFC3339 value and rotation years, months, days, hours, and minutes, separated by commas (`,`), e.g. for 30 days console
//
// ```sh
//  $ pulumi import time:index/timeRotating:TimeRotating example 2020-02-12T06:36:13Z,0,0,30,0,0
// ```
//
//  Otherwise, to import with the rotation RFC3339 value, the base UTC RFC3339 value and rotation UTC RFC3339 value, separated by commas (`,`), e.g. console
//
// ```sh
//  $ pulumi import time:index/timeRotating:TimeRotating example 2020-02-12T06:36:13Z,2020-02-13T06:36:13Z
// ```
//
//  The `triggers` argument cannot be imported.
type TimeRotating struct {
	pulumi.CustomResourceState

	// Number day of timestamp.
	Day pulumi.IntOutput `pulumi:"day"`
	// Number hour of timestamp.
	Hour pulumi.IntOutput `pulumi:"hour"`
	// Number minute of timestamp.
	Minute pulumi.IntOutput `pulumi:"minute"`
	// Number month of timestamp.
	Month pulumi.IntOutput `pulumi:"month"`
	// Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
	Rfc3339 pulumi.StringOutput `pulumi:"rfc3339"`
	// Number of days to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationDays pulumi.IntPtrOutput `pulumi:"rotationDays"`
	// Number of hours to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationHours pulumi.IntPtrOutput `pulumi:"rotationHours"`
	// Number of minutes to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMinutes pulumi.IntPtrOutput `pulumi:"rotationMinutes"`
	// Number of months to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMonths pulumi.IntPtrOutput `pulumi:"rotationMonths"`
	// Configure the rotation timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationRfc3339 pulumi.StringOutput `pulumi:"rotationRfc3339"`
	// Number of years to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationYears pulumi.IntPtrOutput `pulumi:"rotationYears"`
	// Number second of timestamp.
	Second pulumi.IntOutput `pulumi:"second"`
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. These conditions recreate the resource in addition to other rotation arguments. See the main provider documentation for more information.
	Triggers pulumi.StringMapOutput `pulumi:"triggers"`
	// Number of seconds since epoch time, e.g. `1581489373`.
	Unix pulumi.IntOutput `pulumi:"unix"`
	// Number year of timestamp.
	Year pulumi.IntOutput `pulumi:"year"`
}

// NewTimeRotating registers a new resource with the given unique name, arguments, and options.
func NewTimeRotating(ctx *pulumi.Context,
	name string, args *TimeRotatingArgs, opts ...pulumi.ResourceOption) (*TimeRotating, error) {
	if args == nil {
		args = &TimeRotatingArgs{}
	}

	var resource TimeRotating
	err := ctx.RegisterResource("time:index/timeRotating:TimeRotating", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeRotating gets an existing TimeRotating resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeRotating(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeRotatingState, opts ...pulumi.ResourceOption) (*TimeRotating, error) {
	var resource TimeRotating
	err := ctx.ReadResource("time:index/timeRotating:TimeRotating", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeRotating resources.
type timeRotatingState struct {
	// Number day of timestamp.
	Day *int `pulumi:"day"`
	// Number hour of timestamp.
	Hour *int `pulumi:"hour"`
	// Number minute of timestamp.
	Minute *int `pulumi:"minute"`
	// Number month of timestamp.
	Month *int `pulumi:"month"`
	// Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
	Rfc3339 *string `pulumi:"rfc3339"`
	// Number of days to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationDays *int `pulumi:"rotationDays"`
	// Number of hours to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationHours *int `pulumi:"rotationHours"`
	// Number of minutes to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMinutes *int `pulumi:"rotationMinutes"`
	// Number of months to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMonths *int `pulumi:"rotationMonths"`
	// Configure the rotation timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationRfc3339 *string `pulumi:"rotationRfc3339"`
	// Number of years to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationYears *int `pulumi:"rotationYears"`
	// Number second of timestamp.
	Second *int `pulumi:"second"`
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. These conditions recreate the resource in addition to other rotation arguments. See the main provider documentation for more information.
	Triggers map[string]string `pulumi:"triggers"`
	// Number of seconds since epoch time, e.g. `1581489373`.
	Unix *int `pulumi:"unix"`
	// Number year of timestamp.
	Year *int `pulumi:"year"`
}

type TimeRotatingState struct {
	// Number day of timestamp.
	Day pulumi.IntPtrInput
	// Number hour of timestamp.
	Hour pulumi.IntPtrInput
	// Number minute of timestamp.
	Minute pulumi.IntPtrInput
	// Number month of timestamp.
	Month pulumi.IntPtrInput
	// Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
	Rfc3339 pulumi.StringPtrInput
	// Number of days to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationDays pulumi.IntPtrInput
	// Number of hours to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationHours pulumi.IntPtrInput
	// Number of minutes to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMinutes pulumi.IntPtrInput
	// Number of months to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMonths pulumi.IntPtrInput
	// Configure the rotation timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationRfc3339 pulumi.StringPtrInput
	// Number of years to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationYears pulumi.IntPtrInput
	// Number second of timestamp.
	Second pulumi.IntPtrInput
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. These conditions recreate the resource in addition to other rotation arguments. See the main provider documentation for more information.
	Triggers pulumi.StringMapInput
	// Number of seconds since epoch time, e.g. `1581489373`.
	Unix pulumi.IntPtrInput
	// Number year of timestamp.
	Year pulumi.IntPtrInput
}

func (TimeRotatingState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeRotatingState)(nil)).Elem()
}

type timeRotatingArgs struct {
	// Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
	Rfc3339 *string `pulumi:"rfc3339"`
	// Number of days to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationDays *int `pulumi:"rotationDays"`
	// Number of hours to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationHours *int `pulumi:"rotationHours"`
	// Number of minutes to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMinutes *int `pulumi:"rotationMinutes"`
	// Number of months to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMonths *int `pulumi:"rotationMonths"`
	// Configure the rotation timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationRfc3339 *string `pulumi:"rotationRfc3339"`
	// Number of years to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationYears *int `pulumi:"rotationYears"`
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. These conditions recreate the resource in addition to other rotation arguments. See the main provider documentation for more information.
	Triggers map[string]string `pulumi:"triggers"`
}

// The set of arguments for constructing a TimeRotating resource.
type TimeRotatingArgs struct {
	// Configure the base timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
	Rfc3339 pulumi.StringPtrInput
	// Number of days to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationDays pulumi.IntPtrInput
	// Number of hours to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationHours pulumi.IntPtrInput
	// Number of minutes to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMinutes pulumi.IntPtrInput
	// Number of months to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationMonths pulumi.IntPtrInput
	// Configure the rotation timestamp with an UTC [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`). When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationRfc3339 pulumi.StringPtrInput
	// Number of years to add to the base timestamp to configure the rotation timestamp. When the current time has passed the rotation timestamp, the resource will trigger recreation. Conflicts with other `rotation_` arguments.
	RotationYears pulumi.IntPtrInput
	// Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. These conditions recreate the resource in addition to other rotation arguments. See the main provider documentation for more information.
	Triggers pulumi.StringMapInput
}

func (TimeRotatingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeRotatingArgs)(nil)).Elem()
}

type TimeRotatingInput interface {
	pulumi.Input

	ToTimeRotatingOutput() TimeRotatingOutput
	ToTimeRotatingOutputWithContext(ctx context.Context) TimeRotatingOutput
}

func (*TimeRotating) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeRotating)(nil))
}

func (i *TimeRotating) ToTimeRotatingOutput() TimeRotatingOutput {
	return i.ToTimeRotatingOutputWithContext(context.Background())
}

func (i *TimeRotating) ToTimeRotatingOutputWithContext(ctx context.Context) TimeRotatingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeRotatingOutput)
}

func (i *TimeRotating) ToTimeRotatingPtrOutput() TimeRotatingPtrOutput {
	return i.ToTimeRotatingPtrOutputWithContext(context.Background())
}

func (i *TimeRotating) ToTimeRotatingPtrOutputWithContext(ctx context.Context) TimeRotatingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeRotatingPtrOutput)
}

type TimeRotatingPtrInput interface {
	pulumi.Input

	ToTimeRotatingPtrOutput() TimeRotatingPtrOutput
	ToTimeRotatingPtrOutputWithContext(ctx context.Context) TimeRotatingPtrOutput
}

type timeRotatingPtrType TimeRotatingArgs

func (*timeRotatingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeRotating)(nil))
}

func (i *timeRotatingPtrType) ToTimeRotatingPtrOutput() TimeRotatingPtrOutput {
	return i.ToTimeRotatingPtrOutputWithContext(context.Background())
}

func (i *timeRotatingPtrType) ToTimeRotatingPtrOutputWithContext(ctx context.Context) TimeRotatingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeRotatingPtrOutput)
}

// TimeRotatingArrayInput is an input type that accepts TimeRotatingArray and TimeRotatingArrayOutput values.
// You can construct a concrete instance of `TimeRotatingArrayInput` via:
//
//          TimeRotatingArray{ TimeRotatingArgs{...} }
type TimeRotatingArrayInput interface {
	pulumi.Input

	ToTimeRotatingArrayOutput() TimeRotatingArrayOutput
	ToTimeRotatingArrayOutputWithContext(context.Context) TimeRotatingArrayOutput
}

type TimeRotatingArray []TimeRotatingInput

func (TimeRotatingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TimeRotating)(nil)).Elem()
}

func (i TimeRotatingArray) ToTimeRotatingArrayOutput() TimeRotatingArrayOutput {
	return i.ToTimeRotatingArrayOutputWithContext(context.Background())
}

func (i TimeRotatingArray) ToTimeRotatingArrayOutputWithContext(ctx context.Context) TimeRotatingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeRotatingArrayOutput)
}

// TimeRotatingMapInput is an input type that accepts TimeRotatingMap and TimeRotatingMapOutput values.
// You can construct a concrete instance of `TimeRotatingMapInput` via:
//
//          TimeRotatingMap{ "key": TimeRotatingArgs{...} }
type TimeRotatingMapInput interface {
	pulumi.Input

	ToTimeRotatingMapOutput() TimeRotatingMapOutput
	ToTimeRotatingMapOutputWithContext(context.Context) TimeRotatingMapOutput
}

type TimeRotatingMap map[string]TimeRotatingInput

func (TimeRotatingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TimeRotating)(nil)).Elem()
}

func (i TimeRotatingMap) ToTimeRotatingMapOutput() TimeRotatingMapOutput {
	return i.ToTimeRotatingMapOutputWithContext(context.Background())
}

func (i TimeRotatingMap) ToTimeRotatingMapOutputWithContext(ctx context.Context) TimeRotatingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeRotatingMapOutput)
}

type TimeRotatingOutput struct{ *pulumi.OutputState }

func (TimeRotatingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeRotating)(nil))
}

func (o TimeRotatingOutput) ToTimeRotatingOutput() TimeRotatingOutput {
	return o
}

func (o TimeRotatingOutput) ToTimeRotatingOutputWithContext(ctx context.Context) TimeRotatingOutput {
	return o
}

func (o TimeRotatingOutput) ToTimeRotatingPtrOutput() TimeRotatingPtrOutput {
	return o.ToTimeRotatingPtrOutputWithContext(context.Background())
}

func (o TimeRotatingOutput) ToTimeRotatingPtrOutputWithContext(ctx context.Context) TimeRotatingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeRotating) *TimeRotating {
		return &v
	}).(TimeRotatingPtrOutput)
}

type TimeRotatingPtrOutput struct{ *pulumi.OutputState }

func (TimeRotatingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeRotating)(nil))
}

func (o TimeRotatingPtrOutput) ToTimeRotatingPtrOutput() TimeRotatingPtrOutput {
	return o
}

func (o TimeRotatingPtrOutput) ToTimeRotatingPtrOutputWithContext(ctx context.Context) TimeRotatingPtrOutput {
	return o
}

func (o TimeRotatingPtrOutput) Elem() TimeRotatingOutput {
	return o.ApplyT(func(v *TimeRotating) TimeRotating {
		if v != nil {
			return *v
		}
		var ret TimeRotating
		return ret
	}).(TimeRotatingOutput)
}

type TimeRotatingArrayOutput struct{ *pulumi.OutputState }

func (TimeRotatingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeRotating)(nil))
}

func (o TimeRotatingArrayOutput) ToTimeRotatingArrayOutput() TimeRotatingArrayOutput {
	return o
}

func (o TimeRotatingArrayOutput) ToTimeRotatingArrayOutputWithContext(ctx context.Context) TimeRotatingArrayOutput {
	return o
}

func (o TimeRotatingArrayOutput) Index(i pulumi.IntInput) TimeRotatingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeRotating {
		return vs[0].([]TimeRotating)[vs[1].(int)]
	}).(TimeRotatingOutput)
}

type TimeRotatingMapOutput struct{ *pulumi.OutputState }

func (TimeRotatingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TimeRotating)(nil))
}

func (o TimeRotatingMapOutput) ToTimeRotatingMapOutput() TimeRotatingMapOutput {
	return o
}

func (o TimeRotatingMapOutput) ToTimeRotatingMapOutputWithContext(ctx context.Context) TimeRotatingMapOutput {
	return o
}

func (o TimeRotatingMapOutput) MapIndex(k pulumi.StringInput) TimeRotatingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TimeRotating {
		return vs[0].(map[string]TimeRotating)[vs[1].(string)]
	}).(TimeRotatingOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeRotatingOutput{})
	pulumi.RegisterOutputType(TimeRotatingPtrOutput{})
	pulumi.RegisterOutputType(TimeRotatingArrayOutput{})
	pulumi.RegisterOutputType(TimeRotatingMapOutput{})
}
