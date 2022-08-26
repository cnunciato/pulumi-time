// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.time;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.time.OffsetArgs;
import com.pulumiverse.time.Utilities;
import com.pulumiverse.time.inputs.OffsetState;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.time.Offset;
 * import com.pulumi.time.OffsetArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Offset(&#34;example&#34;, OffsetArgs.builder()        
 *             .offsetDays(7)
 *             .build());
 * 
 *         ctx.export(&#34;oneWeekFromNow&#34;, example.rfc3339());
 *     }
 * }
 * ```
 * ### Triggers Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.time.Offset;
 * import com.pulumi.time.OffsetArgs;
 * import com.pulumi.aws.ec2.Instance;
 * import com.pulumi.aws.ec2.InstanceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var amiUpdate = new Offset(&#34;amiUpdate&#34;, OffsetArgs.builder()        
 *             .triggers(Map.of(&#34;ami_id&#34;, data.aws_ami().example().id()))
 *             .offsetDays(7)
 *             .build());
 * 
 *         var server = new Instance(&#34;server&#34;, InstanceArgs.builder()        
 *             .ami(amiUpdate.triggers().applyValue(triggers -&gt; triggers.amiId()))
 *             .tags(Map.of(&#34;ExpirationTime&#34;, amiUpdate.rfc3339()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using the base UTC RFC3339 timestamp and offset years, months, days, hours, minutes, and seconds, separated by commas (`,`), e.g.
 * 
 * ```sh
 *  $ pulumi import time:index/offset:Offset example 2020-02-12T06:36:13Z,0,0,7,0,0,0
 * ```
 * 
 *  The `triggers` argument cannot be imported.
 * 
 */
@ResourceType(type="time:index/offset:Offset")
public class Offset extends com.pulumi.resources.CustomResource {
    /**
     * Base timestamp in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format (see [RFC3339 time
     * string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., `YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
     * 
     */
    @Export(name="baseRfc3339", type=String.class, parameters={})
    private Output<String> baseRfc3339;

    /**
     * @return Base timestamp in [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format (see [RFC3339 time
     * string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., `YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.
     * 
     */
    public Output<String> baseRfc3339() {
        return this.baseRfc3339;
    }
    /**
     * Number day of offset timestamp.
     * 
     */
    @Export(name="day", type=Integer.class, parameters={})
    private Output<Integer> day;

    /**
     * @return Number day of offset timestamp.
     * 
     */
    public Output<Integer> day() {
        return this.day;
    }
    /**
     * Number hour of offset timestamp.
     * 
     */
    @Export(name="hour", type=Integer.class, parameters={})
    private Output<Integer> hour;

    /**
     * @return Number hour of offset timestamp.
     * 
     */
    public Output<Integer> hour() {
        return this.hour;
    }
    /**
     * Number minute of offset timestamp.
     * 
     */
    @Export(name="minute", type=Integer.class, parameters={})
    private Output<Integer> minute;

    /**
     * @return Number minute of offset timestamp.
     * 
     */
    public Output<Integer> minute() {
        return this.minute;
    }
    /**
     * Number month of offset timestamp.
     * 
     */
    @Export(name="month", type=Integer.class, parameters={})
    private Output<Integer> month;

    /**
     * @return Number month of offset timestamp.
     * 
     */
    public Output<Integer> month() {
        return this.month;
    }
    /**
     * Number of days to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    @Export(name="offsetDays", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> offsetDays;

    /**
     * @return Number of days to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    public Output<Optional<Integer>> offsetDays() {
        return Codegen.optional(this.offsetDays);
    }
    /**
     * Number of hours to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    @Export(name="offsetHours", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> offsetHours;

    /**
     * @return Number of hours to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    public Output<Optional<Integer>> offsetHours() {
        return Codegen.optional(this.offsetHours);
    }
    /**
     * Number of minutes to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    @Export(name="offsetMinutes", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> offsetMinutes;

    /**
     * @return Number of minutes to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    public Output<Optional<Integer>> offsetMinutes() {
        return Codegen.optional(this.offsetMinutes);
    }
    /**
     * Number of months to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    @Export(name="offsetMonths", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> offsetMonths;

    /**
     * @return Number of months to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    public Output<Optional<Integer>> offsetMonths() {
        return Codegen.optional(this.offsetMonths);
    }
    /**
     * Number of seconds to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    @Export(name="offsetSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> offsetSeconds;

    /**
     * @return Number of seconds to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    public Output<Optional<Integer>> offsetSeconds() {
        return Codegen.optional(this.offsetSeconds);
    }
    /**
     * Number of years to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    @Export(name="offsetYears", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> offsetYears;

    /**
     * @return Number of years to offset the base timestamp. At least one of the &#39;offset_&#39; arguments must be configured.
     * 
     */
    public Output<Optional<Integer>> offsetYears() {
        return Codegen.optional(this.offsetYears);
    }
    /**
     * RFC3339 format of the offset timestamp, e.g. `2020-02-12T06:36:13Z`.
     * 
     */
    @Export(name="rfc3339", type=String.class, parameters={})
    private Output<String> rfc3339;

    /**
     * @return RFC3339 format of the offset timestamp, e.g. `2020-02-12T06:36:13Z`.
     * 
     */
    public Output<String> rfc3339() {
        return this.rfc3339;
    }
    /**
     * Number second of offset timestamp.
     * 
     */
    @Export(name="second", type=Integer.class, parameters={})
    private Output<Integer> second;

    /**
     * @return Number second of offset timestamp.
     * 
     */
    public Output<Integer> second() {
        return this.second;
    }
    /**
     * Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. See [the main provider
     * documentation](../index.md) for more information.
     * 
     */
    @Export(name="triggers", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> triggers;

    /**
     * @return Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved. See [the main provider
     * documentation](../index.md) for more information.
     * 
     */
    public Output<Optional<Map<String,String>>> triggers() {
        return Codegen.optional(this.triggers);
    }
    /**
     * Number of seconds since epoch time, e.g. `1581489373`.
     * 
     */
    @Export(name="unix", type=Integer.class, parameters={})
    private Output<Integer> unix;

    /**
     * @return Number of seconds since epoch time, e.g. `1581489373`.
     * 
     */
    public Output<Integer> unix() {
        return this.unix;
    }
    /**
     * Number year of offset timestamp.
     * 
     */
    @Export(name="year", type=Integer.class, parameters={})
    private Output<Integer> year;

    /**
     * @return Number year of offset timestamp.
     * 
     */
    public Output<Integer> year() {
        return this.year;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Offset(String name) {
        this(name, OffsetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Offset(String name, @Nullable OffsetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Offset(String name, @Nullable OffsetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("time:index/offset:Offset", name, args == null ? OffsetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Offset(String name, Output<String> id, @Nullable OffsetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("time:index/offset:Offset", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Offset get(String name, Output<String> id, @Nullable OffsetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Offset(name, id, state, options);
    }
}