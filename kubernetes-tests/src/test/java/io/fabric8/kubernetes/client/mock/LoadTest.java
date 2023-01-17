/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.fabric8.kubernetes.client.mock;

import io.fabric8.kubernetes.api.model.HasMetadata;
import io.fabric8.kubernetes.api.model.networking.v1.NetworkPolicy;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.server.mock.KubernetesServer;
import org.junit.Rule;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.migrationsupport.rules.EnableRuleMigrationSupport;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@EnableRuleMigrationSupport
class LoadTest {

  @Rule
  public KubernetesServer server = new KubernetesServer();

  @Test
  void testResourceGetFromLoadWhenMultipleDocumentsWithDelimiter() throws Exception {
    // given
    KubernetesClient client = server.getClient();

    // when
    List<HasMetadata> result = client.load(getClass().getResourceAsStream("/multiple-document-template.yml")).get();

    // then
    assertNotNull(result);
    assertEquals(6, result.size());
    HasMetadata deploymentResource = result.get(1);
    assertEquals("apps/v1", deploymentResource.getApiVersion());
    assertEquals("Deployment", deploymentResource.getKind());
    assertEquals("redis-master", deploymentResource.getMetadata().getName());
  }

  @Test
  void testNetworkPolicyLoad() {
    KubernetesClient client = server.getClient();
    List<HasMetadata> itemList = client.load(getClass().getResourceAsStream("/test-networkpolicy.yml")).get();

    assertEquals(1, itemList.size());
    NetworkPolicy ingress = (NetworkPolicy) itemList.get(0);
    assertEquals("test-network-policy", ingress.getMetadata().getName());
    assertEquals(1, ingress.getSpec().getIngress().size());
    assertEquals(1, ingress.getSpec().getEgress().size());
  }
}
