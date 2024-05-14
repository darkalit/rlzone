import axios from "axios";

const apiUrl = process.env.REACT_APP_API_URL + "/items";

export function GetItems() {
  // return axios.get(apiUrl, {}).then((res) => {
  //   return res.data;
  // });
  return fetch(apiUrl).then((res) => {
    return res.json();
  });
}
