namespace go nusadua.router.rpc

exception RouterException {
    1: i32 status,
    2: string message
}

service RouterService {
    void setNodes(1:list<string> nodes) throws (1:RouterException ex),
    list<string> getNodesByFeature(1:list<i64> key) throws (1:RouterException ex)
}
