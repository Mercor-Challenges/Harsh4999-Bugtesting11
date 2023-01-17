
package io.fabric8.openshift.api.model.hive.gcp.v1;

import java.util.HashMap;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import io.fabric8.kubernetes.api.model.Container;
import io.fabric8.kubernetes.api.model.IntOrString;
import io.fabric8.kubernetes.api.model.KubernetesResource;
import io.fabric8.kubernetes.api.model.LabelSelector;
import io.fabric8.kubernetes.api.model.LocalObjectReference;
import io.fabric8.kubernetes.api.model.ObjectMeta;
import io.fabric8.kubernetes.api.model.ObjectReference;
import io.fabric8.kubernetes.api.model.PersistentVolumeClaim;
import io.fabric8.kubernetes.api.model.PodTemplateSpec;
import io.fabric8.kubernetes.api.model.ResourceRequirements;
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
    "kmsKey",
    "kmsKeyServiceAccount"
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
    @BuildableReference(LabelSelector.class),
    @BuildableReference(Container.class),
    @BuildableReference(PodTemplateSpec.class),
    @BuildableReference(ResourceRequirements.class),
    @BuildableReference(IntOrString.class),
    @BuildableReference(ObjectReference.class),
    @BuildableReference(LocalObjectReference.class),
    @BuildableReference(PersistentVolumeClaim.class)
})
public class EncryptionKeyReference implements KubernetesResource
{

    @JsonProperty("kmsKey")
    private KMSKeyReference kmsKey;
    @JsonProperty("kmsKeyServiceAccount")
    private String kmsKeyServiceAccount;
    @JsonIgnore
    private Map<String, Object> additionalProperties = new HashMap<String, Object>();

    /**
     * No args constructor for use in serialization
     * 
     */
    public EncryptionKeyReference() {
    }

    /**
     * 
     * @param kmsKeyServiceAccount
     * @param kmsKey
     */
    public EncryptionKeyReference(KMSKeyReference kmsKey, String kmsKeyServiceAccount) {
        super();
        this.kmsKey = kmsKey;
        this.kmsKeyServiceAccount = kmsKeyServiceAccount;
    }

    @JsonProperty("kmsKey")
    public KMSKeyReference getKmsKey() {
        return kmsKey;
    }

    @JsonProperty("kmsKey")
    public void setKmsKey(KMSKeyReference kmsKey) {
        this.kmsKey = kmsKey;
    }

    @JsonProperty("kmsKeyServiceAccount")
    public String getKmsKeyServiceAccount() {
        return kmsKeyServiceAccount;
    }

    @JsonProperty("kmsKeyServiceAccount")
    public void setKmsKeyServiceAccount(String kmsKeyServiceAccount) {
        this.kmsKeyServiceAccount = kmsKeyServiceAccount;
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
