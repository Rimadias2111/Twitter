import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 50,
    duration: '30s',
};

export default function () {
    let payload = JSON.stringify({
        user_id: '123e4567-e89b-12d3-a456-426614174000',
        content: 'This is a test tweet',
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer YOUR_ACCESS_TOKEN',
        },
    };

    let response = http.post('http://localhost:8000/v1/tweets', payload, params);

    check(response, {
        'is status 200': (r) => r.status === 200,
        'response time is less than 200ms': (r) => r.timings.duration < 200,
    });

}
