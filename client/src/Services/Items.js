import axios from "axios";

const apiUrl = process.env.REACT_APP_API_URL + "/items";

export function GetItems() {
  return fetch(apiUrl).then((res) => {
    return res.json();
  });
}

export function GetItemByID(id) {
  return fetch(apiUrl + "/" + id).then((res) => {
    return res.json();
  });
}
