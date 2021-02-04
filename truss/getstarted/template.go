package getstarted

const starterProto = `
syntax = "proto3";

package {{.PackageName}};

import "google/api/annotations.proto";

service {{.ServiceName}} {
  rpc Status(StatusRequest) returns (StatusResponse) {
    option (google.api.http) = {
      get: "/status"
    };
  }
}

enum ServiceStatus {
  FAIL = 0;
  OK = 1;
}

message StatusRequest {
  bool full = 1;
}

message StatusResponse {
  ServiceStatus status = 1;
}
`
