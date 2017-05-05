namespace go server.rpc

exception ParameterServerException {
    1: i32 status,
    2: string message
}

service ParameterServer {
    void push(1:list<i64> key, 2:list<double> value) throws (1: ParameterServerException ex),
    list<double> pull(1:list<i64> key) throws (1: ParameterServerException ex)
}
