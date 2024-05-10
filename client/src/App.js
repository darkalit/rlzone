import "./App.css";
import { Footer, Header } from "./Components";
import { UserControl } from "./Pages";
import { Routes, Route } from "react-router-dom";

export default function App() {
  return (
    <>
      <Header role="Guest" />

      <Routes>
        <Route exact path="/" element={<UserControl />} />
      </Routes>

      <Footer />
    </>
  );
}
