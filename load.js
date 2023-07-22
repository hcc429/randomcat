import http from "k6/http";
import { sleep } from "k6";
import { check } from 'k6';
import { Rate, Trend } from "k6/metrics";
//docker run --rm -i grafana/k6 run - <load.js

// Define custom metrics
//const failedRequests = new Rate("failed_requests");
//const requestsPerSecond = new Trend("requests_per_second");

export const options = {
  stages: [
    { duration: '10s', target: 200 },
    { duration: '100s', target: 200 },
    { duration: '10s', target: 100}
  ],
};
// export const options = {
//   vus: 2000,
//   duration: '10s',
//   rps: 120
// };
export const sameImage = false
export default function () {
  const url = "https://randomcat.io/api/image?w=100&h=170";
  const res = http.get(url);
  check(res, { 'status was 200': (r) => r.status == 200 });
  //sleep(1);
}
