import { check, group } from "k6";
import http from "k6/http";

export const options = {
  vus: 1,
  //vus: 5, //Ramp up test
  //duration:'10s',
  thresholds: {
    http_req_failed: ["rate<0.01"], // http errors should be less than 1%
    http_req_duration: ["p(99)<3000"], // 99% of requests should be below 3s
  },
};

const img1binFile = open("/scripts/assets/lemon.jpg", "b");
const img2binFile = open("/scripts/assets/demo.jpg", "b");
const img3binFile = open("/scripts/assets/blackmarble.jpg", "b");
const img4binFile = open("/scripts/assets/bluemarble.jpeg", "b");
const img5binFile = open("/scripts/assets/bluemarble.png", "b");


export default function () {
  const url = `http://${__ENV.CONTROL_PLANE}/api/v1/challenge/`;


  group("123KB", function () {
    const data = {
      image: http.file(img1binFile, "lemon.jpg"),
    };
  
    const res = http.post(url, data);
  
    check(res, {
      "is status 201:created - 123KB": (r) => r.status === 201,
    });
   
  });

  group("496KB", function () {
    const data = {
      image: http.file(img2binFile, "demo.jpg"),
    };
  
    const res = http.post(url, data);
  
    check(res, {
      "is status 201:created - 496KB": (r) => r.status === 201,
    });
   
  });

  group("3.5MB", function () {

    const data = {
      image: http.file(img3binFile, "blackmarble.jpg"),
    };
  
    const res = http.post(url, data);
  
    check(res, {
      "is status 201:created - 3.5MB": (r) => r.status === 201,
    });
   
  });

  group("17.1MB", function () {
    const data = {
      image: http.file(img4binFile, "bluemarble.jpeg"),
    };
  
    const res = http.post(url, data);
  
    check(res, {
      "is status 201:created - 17.1MB": (r) => r.status === 201,
    });
   
  });
  
  group("23.4MB", function () {
    const data = {
      image: http.file(img5binFile, "bluemarble.png"),
    };
  
    const res = http.post(url, data);
  
    check(res, {
      "is status 201:created - 23.4MB": (r) => r.status === 201,
    });
  });

  group("495KB", function () {
    const data = {
      image: http.file(img2binFile, "demo.jpg"),
    };
  
    const res = http.post(url, data);
  
    check(res, {
      "is status 201:created - 496KB": (r) => r.status === 201,
    });
  });
  
}
