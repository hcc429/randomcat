import http from "k6/http";
import { sleep } from "k6";



export default () => {
  http.get("http://host.docker.internal:8080/image/?w=100&h=170");
  // MORE STEPS
  // Here you can have more steps or complex script
  // Step1
  // Step2
  // etc.
};
