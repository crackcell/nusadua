namespace go nusadua.shepherd.rpc

exception ShepherdException {
    1: i32 status,
    2: string message
}

service ShepherdService {
    void setNodes(1:list<string> nodes) throws (1:ShepherdException ex),
    list<string> getNodesByFeature(1:list<list<i64>> key) throws (1:ShepherdException ex)
}
