import http from 'k6/http';
import { check } from 'k6';

// todo 测试前修改待测试的目标服务IP地址
const baseUrl = 'http://192.168.3.37:8080';

const data = { name: 'qwertyuiopasdfghjklzxcvbnm'.repeat(2) }; // 52 bytes

export default function() {
    const response = http.get(`${baseUrl}/helloworld/${data}`);
    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}
