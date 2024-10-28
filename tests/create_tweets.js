import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 50,
    duration: '30s',
};

export default function () {
    let payload = JSON.stringify({
        content: 'This is a test tweet',
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNDAxYTBlYTUtYmYzYS00M2NmLWI5OTItZTYwOWI0ZDc2ZTA4Iiwicm9sZSI6InVzZXIiLCJleHAiOjE3MzAyMjUyMTd9.XGNAozk8BJlAQh14dyqV4uSw_nPZ1KK3665YINp1HbI',
        },
    };

    let response = http.post('http://localhost:8080/v1/tweets', payload, params);

    check(response, {
        'is status 200': (r) => r.status === 200,
        'response time is less than 200ms': (r) => r.timings.duration < 200,
    });

}
