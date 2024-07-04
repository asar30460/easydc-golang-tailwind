import { BrowserRouter, Route, Routes } from "react-router-dom";
import { LoginOrRegister, Platform } from "./components/";
import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LoginOrRegister />} />
        <Route path="/server/*" element={<Platform />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
