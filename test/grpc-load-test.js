import grpc from 'k6/net/grpc';
import { check } from 'k6';
import { Counter, Rate } from "k6/metrics";

// todo 测试前修改目标服务地址
const grpcAddr = '192.168.3.37:8282';

const client = new grpc.Client();
client.load(['.'], 'greeter.proto');

// 注：所有指标名称会自动添加前缀k6_
// Counter指标名称会自动添加后缀_total
let reqTotal = new Counter("grpc_reqs");
// Rate指标名称会自动添加后缀_rate
let failedRate = new Rate("grpc_req_failed");

export default () => {
    client.connect(grpcAddr, {
        plaintext: true   // insecure
    });

    const data = { name: 'qwertyuiopasdfghjklzxcvbnm'.repeat(2) }; // 52 bytes
    const fullPath = 'helloworld.v1.Greeter/SayHello';
    let tags = {
        full_path: fullPath,
        grpc_type: "unary",
        status: "",
    }

    for (let i = 0; i < 1000; i++) {
        let response = client.invoke(fullPath, data);
        let passed = check(response, {
            'status is OK': (r) => {
                tags.status = r.status.toString();
                if (r && r.status === grpc.StatusOK) {
                    return true
                }
                // console.error("Error: ", r.error.message)
                return false
            },
        });

        reqTotal.add(1, tags);
        failedRate.add(passed ? 0 : 1, tags);
        // console.log(JSON.stringify(response.message));
    }

    client.close();
};
