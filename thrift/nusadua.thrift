namespace go nusadua.rpc

exception RpcException {
    1: i32 status,
    2: string message
}

service Router {
    void set_nodes(1:list<string> nodes) throws (1:RpcException ex),
    list<string> get_nodes_by_feature(1:list<i64> key) throws (1:RpcException ex)
}

service Server {
    void multi_push(1:list<list<i64>> keys, 2:list<double> values) throws (1:RpcException ex),
    list<double> multi_pull(1:list<list<i64>> keys) throws (1:RpcException ex),

    void range_push(1:i64 start_key, 2:i64 end_key, 3:list<double> values) throws (1:RpcException ex),
    list<double> range_pull(1:i64 start_key, 2:i64 end_key) throws (1:RpcException ex)
}
