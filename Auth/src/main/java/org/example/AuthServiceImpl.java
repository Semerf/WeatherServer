package org.example;

import com.auth.grpc.Auth;
import com.auth.grpc.AuthServiceGrpc;
import io.grpc.stub.StreamObserver;

public class AuthServiceImpl extends AuthServiceGrpc.AuthServiceImplBase {
    @Override
    public void auth(Auth.AuthRequest request,
                     StreamObserver<Auth.AuthResponse> responseObserver) {
        System.out.println(request);
        Auth.AuthResponse response = Auth.AuthResponse.newBuilder().setAccess(true).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

}
