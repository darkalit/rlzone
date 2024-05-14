import axios from "axios";

const apiUrl = process.env.REACT_APP_API_URL + "/users";

export function Register(epicID, email, password) {
  return axios
    .post(apiUrl + "/register", {
      email: email,
      epic_id: epicID,
      password: password,
    })
    .then((res) => {
      if (res.data.User) {
        localStorage.setItem("User", JSON.stringify(res.data.User));
        localStorage.setItem("AccessToken", res.data.AccessToken);
      }
      return res.data;
    });
}

export function GetStorageUser() {
  return JSON.parse(localStorage.getItem("User"));
}

export function GetStorageAccessToken() {
  return localStorage.getItem("AccessToken");
}

export function Login(email, password) {
  // return axios
  //   .post(
  //     apiUrl + "/login",
  //     {
  //       email: email,
  //       password: password,
  //     },
  //     { withCredentials: false }
  //   )
  //   .then((res) => {
  //     if (res.data.User) {
  //       localStorage.setItem("User", JSON.stringify(res.data.User));
  //       localStorage.setItem("AccessToken", res.data.AccessToken);
  //     }
  //     return res.data;
  //   });
  return fetch(apiUrl + "/login", {
    method: "POST",
    body: JSON.stringify({
      email: email,
      password: password,
    }),
    headers: {
      "Content-Type": "application/json",
    },
    mode: "cors",
  })
    .then((res) => {
      return res.json();
    })
    .then((data) => {
      if (data.User) {
        localStorage.setItem("User", JSON.stringify(data.User));
        localStorage.setItem("AccessToken", data.AccessToken);
      }
      return data;
    });
}

export function Refresh() {
  // return axios.get(apiUrl + "/refresh").then((res) => {
  //   if (res.data.User) {
  //     localStorage.setItem("User", JSON.stringify(res.data.User));
  //     localStorage.setItem("AccessToken", res.data.AccessToken);
  //   }
  //   return res.data;
  // });
}

export function GetUsers() {
  return fetch(apiUrl, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      // Authorization: `Bearer ${GetStorageAccessToken()}`,
    },
    mode: "cors",
  }).then((res) => {
    return res.json();
  });
}
