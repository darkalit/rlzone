import "./App.css";
import { Footer, Header } from "./Components";
import { UserControl, Registration } from "./Pages";
import { Routes, Route  } from "react-router-dom";

export default function App() {
  
  return (
    <>
      <Header role="Auth" />

      <Routes>
        <Route exact path="/" element={<UserControl />} />
        <Route exact path="/registration" element={<Registration />} />
      </Routes>
      <Footer />
    </>
  );
}
