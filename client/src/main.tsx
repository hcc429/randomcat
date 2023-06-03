import ReactDOM from "react-dom/client";
import Home from "./pages/Home.tsx";
import "./styles/index.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { IconContext } from "react-icons";
import Navbar from "./components/Navbar.tsx";
import Gallery from "./pages/Gallery.tsx";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <IconContext.Provider value={{ className: "text-primary", size: "30" }}>
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home key={1} />} />
        <Route path="/gallery" element={<Gallery key={2} />} />
      </Routes>
    </BrowserRouter>
  </IconContext.Provider>
);
