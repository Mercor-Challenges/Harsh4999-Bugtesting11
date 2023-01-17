
package io.fabric8.tekton.triggers.v1beta1;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import io.fabric8.kubernetes.api.model.Container;
import io.fabric8.kubernetes.api.model.ContainerPort;
import io.fabric8.kubernetes.api.model.EnvVar;
import io.fabric8.kubernetes.api.model.IntOrString;
import io.fabric8.kubernetes.api.model.KubernetesResource;
import io.fabric8.kubernetes.api.model.LocalObjectReference;
import io.fabric8.kubernetes.api.model.ObjectMeta;
import io.fabric8.kubernetes.api.model.ObjectReference;
import io.fabric8.kubernetes.api.model.PersistentVolumeClaim;
import io.fabric8.kubernetes.api.model.PodTemplateSpec;
import io.fabric8.kubernetes.api.model.ResourceRequirements;
import io.fabric8.kubernetes.api.model.Volume;
import io.fabric8.kubernetes.api.model.VolumeMount;
import io.sundr.builder.annotations.Buildable;
import io.sundr.builder.annotations.BuildableReference;
import lombok.EqualsAndHashCode;
import lombok.Setter;
import lombok.ToString;
import lombok.experimental.Accessors;

@JsonDeserialize(using = com.fasterxml.jackson.databind.JsonDeserializer.None.class)
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder({
    "apiVersion",
    "kind",
    "metadata",
    "cloudEventURI",
    "labelSelector",
    "namespaceSelector",
    "resources",
    "serviceAccountName",
    "triggerGroups",
    "triggers"
})
@ToString
@EqualsAndHashCode
@Setter
@Accessors(prefix = {
    "_",
    ""
})
@Buildable(editableEnabled = false, validationEnabled = false, generateBuilderPackage = false, lazyCollectionInitEnabled = false, builderPackage = "io.fabric8.kubernetes.api.builder", refs = {
    @BuildableReference(ObjectMeta.class),
    @BuildableReference(io.fabric8.kubernetes.api.model.LabelSelector.class),
    @BuildableReference(Container.class),
    @BuildableReference(PodTemplateSpec.class),
    @BuildableReference(ResourceRequirements.class),
    @BuildableReference(IntOrString.class),
    @BuildableReference(ObjectReference.class),
    @BuildableReference(LocalObjectReference.class),
    @BuildableReference(PersistentVolumeClaim.class),
    @BuildableReference(EnvVar.class),
    @BuildableReference(ContainerPort.class),
    @BuildableReference(Volume.class),
    @BuildableReference(VolumeMount.class)
})
public class EventListenerSpec implements KubernetesResource
{

    @JsonProperty("cloudEventURI")
    private String cloudEventURI;
    @JsonProperty("labelSelector")
    private io.fabric8.kubernetes.api.model.LabelSelector labelSelector;
    @JsonProperty("namespaceSelector")
    private NamespaceSelector namespaceSelector;
    @JsonProperty("resources")
    private Resources resources;
    @JsonProperty("serviceAccountName")
    private String serviceAccountName;
    @JsonProperty("triggerGroups")
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private List<EventListenerTriggerGroup> triggerGroups = new ArrayList<EventListenerTriggerGroup>();
    @JsonProperty("triggers")
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private List<EventListenerTrigger> triggers = new ArrayList<EventListenerTrigger>();
    @JsonIgnore
    private Map<String, Object> additionalProperties = new HashMap<String, Object>();

    /**
     * No args constructor for use in serialization
     * 
     */
    public EventListenerSpec() {
    }

    /**
     * 
     * @param triggerGroups
     * @param serviceAccountName
     * @param cloudEventURI
     * @param labelSelector
     * @param namespaceSelector
     * @param resources
     * @param triggers
     */
    public EventListenerSpec(String cloudEventURI, io.fabric8.kubernetes.api.model.LabelSelector labelSelector, NamespaceSelector namespaceSelector, Resources resources, String serviceAccountName, List<EventListenerTriggerGroup> triggerGroups, List<EventListenerTrigger> triggers) {
        super();
        this.cloudEventURI = cloudEventURI;
        this.labelSelector = labelSelector;
        this.namespaceSelector = namespaceSelector;
        this.resources = resources;
        this.serviceAccountName = serviceAccountName;
        this.triggerGroups = triggerGroups;
        this.triggers = triggers;
    }

    @JsonProperty("cloudEventURI")
    public String getCloudEventURI() {
        return cloudEventURI;
    }

    @JsonProperty("cloudEventURI")
    public void setCloudEventURI(String cloudEventURI) {
        this.cloudEventURI = cloudEventURI;
    }

    @JsonProperty("labelSelector")
    public io.fabric8.kubernetes.api.model.LabelSelector getLabelSelector() {
        return labelSelector;
    }

    @JsonProperty("labelSelector")
    public void setLabelSelector(io.fabric8.kubernetes.api.model.LabelSelector labelSelector) {
        this.labelSelector = labelSelector;
    }

    @JsonProperty("namespaceSelector")
    public NamespaceSelector getNamespaceSelector() {
        return namespaceSelector;
    }

    @JsonProperty("namespaceSelector")
    public void setNamespaceSelector(NamespaceSelector namespaceSelector) {
        this.namespaceSelector = namespaceSelector;
    }

    @JsonProperty("resources")
    public Resources getResources() {
        return resources;
    }

    @JsonProperty("resources")
    public void setResources(Resources resources) {
        this.resources = resources;
    }

    @JsonProperty("serviceAccountName")
    public String getServiceAccountName() {
        return serviceAccountName;
    }

    @JsonProperty("serviceAccountName")
    public void setServiceAccountName(String serviceAccountName) {
        this.serviceAccountName = serviceAccountName;
    }

    @JsonProperty("triggerGroups")
    public List<EventListenerTriggerGroup> getTriggerGroups() {
        return triggerGroups;
    }

    @JsonProperty("triggerGroups")
    public void setTriggerGroups(List<EventListenerTriggerGroup> triggerGroups) {
        this.triggerGroups = triggerGroups;
    }

    @JsonProperty("triggers")
    public List<EventListenerTrigger> getTriggers() {
        return triggers;
    }

    @JsonProperty("triggers")
    public void setTriggers(List<EventListenerTrigger> triggers) {
        this.triggers = triggers;
    }

    @JsonAnyGetter
    public Map<String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    @JsonAnySetter
    public void setAdditionalProperty(String name, Object value) {
        this.additionalProperties.put(name, value);
    }

}
