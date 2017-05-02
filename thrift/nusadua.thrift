namespace go server

exception ParameterServerException {
    1:i32 status,
    2:string message
}

service ParameterServer {
    void push(1:list<i64> keys, 2:list<double> values) throws (1:ParameterServerException exp),
    list<double> pull(1:list<i64> keys) throws (1:ParameterServerException exp)
}
