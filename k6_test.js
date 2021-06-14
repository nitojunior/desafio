import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
    vus: 10,
    duration: '10s',
};

export default function() {
    let indexRes = http.get('http://localhost:8080/')

    if (indexRes !== 200) {
        console.log(JSON.stringify(indexRes.status));
    }

    let loginData = { username: 'John', password: 'Doe'};
    let loginRes = http.post('http://localhost:8080/login', JSON.stringify(loginData), { headers: { 'Content-Type': 'application/json' }});

    if (loginRes !== 200) {
        console.log(JSON.stringify(loginRes.status));
    }

    let contentData = loginRes.json();
    let contentRes = http.post('http://localhost:8080/content', JSON.stringify(contentData), { headers: { 'Content-Type': 'application/json' }});

    if (contentRes !== 200) {
        console.log(JSON.stringify(contentRes.status));
    }

}