package in._10h.ditto;

import com.neovisionaries.ws.client.WebSocket;
import org.eclipse.ditto.base.model.json.JsonSchemaVersion;
import org.eclipse.ditto.client.DisconnectedDittoClient;
import org.eclipse.ditto.client.DittoClient;
import org.eclipse.ditto.client.DittoClients;
import org.eclipse.ditto.client.configuration.BasicAuthenticationConfiguration;
import org.eclipse.ditto.client.configuration.WebSocketMessagingConfiguration;
import org.eclipse.ditto.client.messaging.AuthenticationProvider;
import org.eclipse.ditto.client.messaging.AuthenticationProviders;
import org.eclipse.ditto.client.messaging.MessagingProvider;
import org.eclipse.ditto.client.messaging.MessagingProviders;
import org.eclipse.ditto.things.model.Attributes;
import org.eclipse.ditto.things.model.Feature;
import org.eclipse.ditto.things.model.FeatureBuilder;
import org.eclipse.ditto.things.model.FeatureProperties;
import org.eclipse.ditto.things.model.Thing;
import org.eclipse.ditto.things.model.ThingId;

import java.util.List;

/**
 * Hello world!
 *
 */
public class App {

    public static void main( String[] args ) {

        final DittoClient client = App.initializeClient();
        final ThingId targetId = ThingId.of("com.example.ditto:thingtest");
        final Thing thing = client.twin().create(targetId)
                .toCompletableFuture()
                .join();
        final Attributes attribute = Attributes.newBuilder()
                .set("attrhoge", "valfuga")
                .build();
        final Feature feature = Feature.newBuilder()
                .properties(
                        FeatureProperties.newBuilder()
                                .set("hoge", "fuga")
                                .build()
                )
                .withId("featuretest")
                .build();
        thing.setAttributes(attribute);
        thing.setFeature(feature);
        client.twin().put(thing);

        final List<Thing> foundThings = client.twin().retrieve(targetId)
                .toCompletableFuture()
                .join();
        for (final Thing foundThing : foundThings) {
            System.out.println("foundThing: " + foundThing);
        }

    }

    private static DittoClient initializeClient() {
        final AuthenticationProvider<WebSocket> authenticationProvider =
                AuthenticationProviders.basic(BasicAuthenticationConfiguration.newBuilder()
                        .username("demo1")
                        .password("demo")
                        .build());
        final MessagingProvider messagingProvider = MessagingProviders.webSocket(WebSocketMessagingConfiguration.newBuilder()
                .endpoint("wss://ditto.eclipseprojects.io")
                .jsonSchemaVersion(JsonSchemaVersion.V_2)
                .build(), authenticationProvider);

        final DisconnectedDittoClient disconnectedDittoClient = DittoClients.newInstance(messagingProvider);

        return disconnectedDittoClient.connect()
                .toCompletableFuture()
                .join();

    }

}
