import "./App.css";
import React, { useEffect, useState } from "react";
import { Footer, Header } from "./Components";
import { UserControl, Register } from "./Pages";
import { Routes, Route } from "react-router-dom";
import axios from "axios";
import { GetStorageAccessToken, GetStorageUser, Login } from "./Services/Users";
import { GetItems } from "./Services/Items";

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
        <Route exact path="/" element={<UserControl />} />
        <Route exact path="/register" element={<Register />} />
      </Routes>
      <Footer />
    </>
  );
}
