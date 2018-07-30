if [[ -z "$KAFKA_ZOOKEEPER_CONNECT" ]]; then
    echo "ERROR: missing mandatory config: KAFKA_ZOOKEEPER_CONNECT"
    exit 1
fi
./opt/kafka_2.11-"$KAFKA_VERSION"/bin/kafka-server-start.sh /opt/kafka_2.11-"$KAFKA_VERSION"/config/server.properties --override zookeeper.connect="$KAFKA_ZOOKEEPER_CONNECT"
