import axios from "axios";

const apiUrl = process.env.REACT_APP_API_URL + "/items";

export function GetItems() {
  return fetch(apiUrl).then((res) => {
    return res.json();
  });
}

export async function GetItemByID(id) {
  const res = await fetch(apiUrl + "/" + id);
  return await res.json();
}
