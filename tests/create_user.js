import http from 'k6/http';
import { check, sleep } from 'k6';
import { uuid } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export let options = {
    vus: 50,
    duration: '30s',
};

export default function () {
    let payload = JSON.stringify({
        username: `user_${uuid()}`,
        email: `user_${uuid()}@example.com`,
        password: 'securepassword',
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    let response = http.post('http://localhost:8000/v1/users', payload, params);

    check(response, {
        'is status 200': (r) => r.status === 200,
        'response time is less than 200ms': (r) => r.timings.duration < 200,
    });

}
