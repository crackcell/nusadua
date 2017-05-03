namespace go rpc
namespace java com.crackcell.nusadua.rpc

exception ParameterServiceException {
    1:i32 status,
    2:string message
}

service ParameterService {
    void push(1:list<i64> keys, 2:list<double> values) throws (1:ParameterServiceException ex),
    list<double> pull(1:list<i64> keys) throws (1:ParameterServiceException ex)
}
