python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./auth/auth.proto

python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./predict/predict.proto

python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./ml_model/ml_model.proto