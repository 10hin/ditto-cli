# Ditto CLI

## About

### About this project

CLI command to operating with Eclipse [Ditto][ditto-repo]™.

[ditto-repo]: https://github.com/eclipse/ditto/

I'll implement commands not only for operating Things in Ditto, but also operating Ditto itself with [DevOps Commands][devops-commands].

[devops-commands]: https://www.eclipse.org/ditto/installation-operating.html#devops-commands

And also, I want to implement clients with variable protocols, including HTTP, Websocket, MQTT, and so on.

### About Eclipse Ditto™

[Eclipse Ditto][ditto-web]™ is a technology in the IoT implementing a software pattern called “digital twins”.

[ditto-web]: https://eclipse.org/ditto/

## Progress and Plans

### Supported command/protocol Matrix

Legend:
- <code>&#x2713;</code>: Implemented.
- <code>&#x2717;</code> or empty cell: Not implemented yet.

#### Thing commands

| Command&nbsp;\\&nbsp;Protocol                 | CLI Command                                         | HTTP     | Websocket | AMQP 0.9.1 | AMQP 1.0 | MQTT 3.1.1 | MQTT 5 | HTTP 1.1 | Kafka 2.x |
|:----------------------------------------------|-----------------------------------------------------|----------|-----------|------------|----------|------------|--------|----------|-----------|
| `GET /things?ids={thingId1},{thingId2}[,...]` | `get thing {thingId1} {thingId2} [...]`             | &#x2713; |           |            |          |            |        |          |           |
| `POST /things`                                | `create thing -f {filePath}` (no `{thingId}` param) |          |           |            |          |            |        |          |           |
| `GET /things/{thingId}`                       | `get thing {thingId}`                               | &#x2713; |           |            |          |            |        |          |           |
| `PUT /things/{thingId}` (`If-Non-Match: *`)   | `create thing {thingId} -f {filePath}`              |          |           |            |          |            |        |          |           |
| `PUT /things/{thingId}`                       | `replace thing {thingId} -f {filePath}`             |          |           |            |          |            |        |          |           |
| `PATCH /things/{thingId}`                     | `patch thing {thingId} -f {filePath}`               |          |           |            |          |            |        |          |           |
| `DELETE /things/{thingId}`                    | `delete thing {thingId}`                            |          |           |            |          |            |        |          |           |
| `GET /search/things`                          | `search thing`                                      |          |           |            |          |            |        |          |           |
| `GET /search/things/count`                    | `count thing`                                       |          |           |            |          |            |        |          |           |

#### Policy commands

| Command&nbsp;\\&nbsp;Protocol | CLI Command                 | HTTP | Websocket | AMQP 0.9.1 | AMQP 1.0 | MQTT 3.1.1 | MQTT 5 | HTTP 1.1 | Kafka 2.x |
|:------------------------------|-----------------------------|------|-----------|------------|----------|------------|--------|----------|-----------|
| `GET /policies/{poicyId}`     | `get policy {polityId}`     |      |           |            |          |            |        |          |           |
| `PUT /policies/{poicyId}`     | `replace policy {polityId}` |      |           |            |          |            |        |          |           |
| `DELETE /policies/{poicyId}`  | `delete policy {polityId}`  |      |           |            |          |            |        |          |           |
| `GET /whoami`                 | `whoami`                    |      |           |            |          |            |        |          |           |

#### Message commands

| Command&nbsp;\\&nbsp;Protocol                                                  | CLI Command                                            | HTTP | Websocket | AMQP 0.9.1 | AMQP 1.0 | MQTT 3.1.1 | MQTT 5 | HTTP 1.1 | Kafka 2.x |
|:-------------------------------------------------------------------------------|--------------------------------------------------------|------|-----------|------------|----------|------------|--------|----------|-----------|
| `POST /things/{thingId}/inbox/claim`                                           | `send claim {thingId}`                                 |      |           |            |          |            |        |          |           |
| `POST /things/{thingId}/inbox/messages/{messageSubject}`                       | `send message {thingId} {messageSubject}`              |      |           |            |          |            |        |          |           |
| `POST /things/{thingId}/outbox/messages/{messageSubject}`                      | `reply message {thingId} {messageSubject}`             |      |           |            |          |            |        |          |           |
| `POST /things/{thingId}/features/{featureId}/inbox/messages/{messageSubject}`  | `send message {thingId}/{featureId} {messageSubject}`  |      |           |            |          |            |        |          |           |
| `POST /things/{thingId}/features/{featureId}/outbox/messages/{messageSubject}` | `reply message {thingId}/{featureId} {messageSubject}` |      |           |            |          |            |        |          |           |

#### CloudEvent commands

I have no idea.

#### DevOps commands

| Command&nbsp;\\&nbsp;Protocol                                                                    | CLI Command                                             | HTTP | Websocket | AMQP 0.9.1 | AMQP 1.0 | MQTT 3.1.1 | MQTT 5 | HTTP 1.1 | Kafka 2.x |
|:-------------------------------------------------------------------------------------------------|---------------------------------------------------------|------|-----------|------------|----------|------------|--------|----------|-----------|
| `GET /devops/logging/{service}`                                                                  | `admin get logging {service}`                           |      |           |            |          |            |        |          |           |
| `PUT /devops/logging`                                                                            | `admin replace logging {service} -f {filePath}`         |      |           |            |          |            |        |          |           |
| `GET /devops/logging`                                                                            | `admin get logging all`                                 |      |           |            |          |            |        |          |           |
| `PUT /devops/logging`                                                                            | `admin replace logging all -f {filePath}`               |      |           |            |          |            |        |          |           |
| `GET /devops/config/{service}`                                                                   | `admin get config`                                      |      |           |            |          |            |        |          |           |
| `GET /devops/config`                                                                             | `admin get config all`                                  |      |           |            |          |            |        |          |           |
| `PUT /devops/config`                                                                             | `admin replace config all -f {filePath}`                |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/policies`; `type: policies.commands:createPolicy`                        | `admin create policy {policyId} -f {filePath}`          |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/policies`; `type: policies.commands:retrievePolicy`                      | `admin get policy {policyId}`                           |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/things`; `type: things.commands:createThing`                             | `admin create thing {thingId} -f {filePath}`            |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/things`; `type: things.commands:retrieveThing`                           | `admin get thing {thingId}`                             |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/{service}`; `type: status.commands:retrieveHealth`;                      | `admin cleaner {service} health`                        |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/{service}`; `type: common.commands:retrieveConfig`                       | `admin cleaner {service} get config`                    |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/{service}`; `type: common.commands:modifyConfig`                         | `admin cleaner {service} patch config -f {filePath}`    |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/{service}/{instanceIndex}`; `type: common.commands:shutdown`             | `admin cleaner {service} shutdown`                      |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/{service}/{instanceIndex}`; `type: cleanup.commands:cleanupPersistence`  | `admin cleaner {service} cleanup {entityId}`            |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/thing-search`; `type: status.commands:retrieveHealth`;                   | `admin cleaner thingsSearch health`                     |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/thing-search`; `type: common.commands:retrieveConfig`                    | `admin cleaner thingsSearch get config`                 |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/thing-search`; `type: common.commands:modifyConfig`                      | `admin cleaner thingsSearch patch config -f {filePath}` |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/thing-search/{instanceIndex}`; `type: common.commands:shutdown`          | `admin cleaner thingsSearch shutdown`                   |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/thing-search/{instanceIndex}`; `type: thing-search.commands:updateThing` | `admin cleaner thingsSearch updateThing {thingId}`      |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:createConnection`            | `admin create connection -f {filePath}`                 |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:modifyConnection`            | `admin replace connection -f {filePath}`                |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:retrieveConnection`          | `admin get connection {connectionId}`                   |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:openConnection`              | `admin open connection {connectionId}`                  |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:closeConnection`             | `admin close connection {connectionId}`                 |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:deleteConnection`            | `admin delete connection {connectionId}`                |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:testConnection`              | `admin test connection -f {filePath}`                   |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:retrieveAllConnectionIds`    | `admin get connection allIds`                           |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:retrieveConnectionStatus`    | `admin get connection status {connectionId}`            |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:regrieveConnectionStatus`    | `admin get connection metrics {connectionId}`           |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:resetConnectionMetrics`      | `admin reset connection metrics {connectionId}`         |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:enableConnectionLogs`        | `admin enable connection logs {connectionId}`           |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:retrieveConnectionLogs`      | `admin get connection logs {connectionId}`              |      |           |            |          |            |        |          |           |
| `POST /devops/piggyback/connectivity`; `type: connectivity.commands:resetConnectionLogs`         | `admin reset connection logs {connectionId}`            |      |           |            |          |            |        |          |           |

#### Authentication

- HTTP
  - Basic Authentication
  - OpenID Connect
