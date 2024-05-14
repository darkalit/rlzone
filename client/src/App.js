import "./App.css";
import React, { useEffect, useState } from "react";
import { Footer, Header } from "./Components";
import { UserControl, Register, ViewItem } from "./Pages";
import { Routes, Route } from "react-router-dom";
import axios from "axios";
import { GetStorageAccessToken, GetStorageUser, Login } from "./Services/Users";

axios.interceptors.request.use(
  (req) => {
    req.headers.Authorization = `Bearer ${GetStorageAccessToken()}`;
    return req;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default function App() {
  const [user, setUser] = useState({});
  useEffect(() => {
    setUser(GetStorageUser());
  }, [setUser]);

  useEffect(() => {
    // Login("admin@gmail.com", "admin");
    (async function () {
      console.log(await Login("admin@gmail.com", "admin"));
    })();
  }, []);

  return (
    <>
      <Header role={user?.Role} />

      <Routes>
        <Route exact path="/users" element={<UserControl />} />
        <Route exact path="/register" element={<Register />} />
        <Route exact path="/items/:id" element={<ViewItem />} />
      </Routes>
      <Footer />
    </>
  );
}
