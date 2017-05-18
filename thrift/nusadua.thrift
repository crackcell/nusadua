namespace go nusadua.rpc

exception RpcException {
  1: i32 status,
  2: string message
}


service ParameterServer {
  void create_task(1:string name, 2:i64 max_key) throws (1:RpcException ex),

  // work with parameters
  void multi_push(1:list<list<i64>> keys, 2:list<double> values) throws (1:RpcException ex),
  list<double> multi_pull(1:list<list<i64>> keys) throws (1:RpcException ex),

  void range_push(1:i64 start_key, 2:i64 end_key, 3:list<double> values) throws (1:RpcException ex),
  list<double> range_pull(1:i64 start_key, 2:i64 end_key) throws (1:RpcException ex)
}
