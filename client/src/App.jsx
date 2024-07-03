import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Login, Platform } from "./components/";
import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/server/*" element={<Platform />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
