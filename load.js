import http from "k6/http";
import { sleep } from "k6";
//docker run --rm -i grafana/k6 run - <load.js


export const options = {
  stages: [
    { duration: '10s', target: 1000 },
    { duration: '30s', target: 1000 },
    // { duration: '10s', target: 2000 },
    // { duration: '30s', target: 2000 },
    // { duration: '10s', target: 3000 },
    // { duration: '30s', target: 3000 },
    // { duration: '10s', target: 4000 },
    // { duration: '30s', target: 4000 },
    // { duration: '10s', target: 5000 },
    // { duration: '30s', target: 5000 },
    // { duration: '10s', target: 6000 },
    // { duration: '30s', target: 6000 },
    // { duration: '10s', target: 7000 },
    // { duration: '30s', target: 7000 },
    // { duration: '10s', target: 8000 },
    // { duration: '30s', target: 8000 },
    // { duration: '10s', target: 9000 },
    // { duration: '30s', target: 9000 },

  ],
};
export const sameImage = false
export default function () {
    //const url = "https://randomcat.io/api/image/?w=100&h=170";
    const url = "https://randomcat.io/api/images?p=1&limit=5&sort=newest";
  // if(sameImage){
  //   const ratio = Math.random()+1;
  //   const h = Math.floor(100*ratio); 
  //   url = `https://randomcat.io/api/image/?w=100&h=${h}`;
  // }
  http.get(url);
  sleep(1);
  //sleep(1);
  // Print the response time for each request
  //console.log(`Response time for w=${w} and h=${h}: ${responseTime} ms`);
}
